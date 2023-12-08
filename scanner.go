package escpos

import (
	"errors"
	"time"

	"github.com/bangzek/clock"
)

var (
	ctime = clock.New()
)

type Scanner struct {
	Controller *Controller
	Config     *ScannerConfig

	data   scannerData
	ticker *clock.Ticker
}

type scannerData struct {
	Cache     []Event
	Print     PrintCmd
	Connected bool
	InErr     bool
	Drawer    bool
	FeedBtn   bool
	NearEnd   bool
	State     State
}

var DisconnectErr = errors.New("Printer disconnected")

func (s *Scanner) Scan(cmdCh <-chan []Cmd) <-chan []Event {
	if s.Config == nil {
		s.Config = DefaultScannerConfig()
	}
	s.data.State = OfflineState
	if s.ticker == nil {
		s.ticker = ctime.NewTicker(s.Config.Interval)
	}

	evtCh := make(chan []Event, 1)
	go s.run(cmdCh, evtCh)
	return evtCh
}

func (s *Scanner) run(cmdCh <-chan []Cmd, evtCh chan []Event) {
	defer logPanic()
	defer close(evtCh)

	for s.step(cmdCh) {
		s.flush(evtCh)
	}
	s.data.setConnected(ctime.Now(), false)
	s.flush(evtCh)
}

func (s *Scanner) flush(evtCh chan []Event) {
	if len(s.data.Cache) > 0 {
		evtCh <- append([]Event(nil), s.data.Cache...)
		s.data.Cache = s.data.Cache[:0]
	}
}

func (s *Scanner) step(cmdCh <-chan []Cmd) bool {
	if s.data.Connected {
		return s.connStep(cmdCh)
	} else {
		return s.disStep(cmdCh)
	}
}

func (s *Scanner) connStep(cmdCh <-chan []Cmd) bool {
	select {
	case list, ok := <-cmdCh:
		if !ok {
			return false
		}
		for _, cmd := range list {
			if s.data.Connected {
				dis := false
				if pc, ok := cmd.(PrintCmd); ok {
					dis = pc.NeedDisablePulseLevel()
					s.data.startPrint(ctime.Now(), pc)
				}
				if dis {
					if err := s.Controller.DisablePulseLevel(); err != nil {
						now := ctime.Now()
						s.connErr(true, now, err, "disable pulse level")
						s.data.purge(now, cmd)
						continue
					}
				}
				debugLog("EXEC %s", cmd)
				res := cmd.exec(s.Controller)
				now := ctime.Now()
				s.data.cmdRes(now, cmd, res)
				if err := res.Error(); err != nil {
					s.connErr(false, now, err, "exec "+cmd.String())
				}
				if s.data.Connected && dis {
					if err := s.Controller.EnablePulseLevel(); err != nil {
						s.connErr(false, now, err, "enable pulse level")
					}
				}
			} else {
				s.data.purge(ctime.Now(), cmd)
			}
		}
	case <-s.ticker.C:
		if ps, err := s.Controller.PrinterStatus(); err != nil {
			s.connErr(true, ctime.Now(), err, "getting printer status")
			return true
		} else {
			now := ctime.Now()
			s.data.finishPrint(now)
			s.data.setDrawer(now, ps.IsDrawerPin3())
			s.data.setFeedBtn(now, ps.IsFeedButton())
			state := OnlineState
			if ps.IsOffline() {
				if os, err := s.Controller.OfflineStatus(); err != nil {
					s.connErr(false, ctime.Now(), err, "getting offline status")
					return true
				} else {
					now = ctime.Now()
					if os.IsCoverOpen() {
						state = CoverOpenState
					} else if os.IsPaperEnd() {
						state = NoPaperState
					} else if os.IsError() {
						if es, err := s.Controller.ErrorStatus(); err != nil {
							s.connErr(false, ctime.Now(), err,
								"getting error status")
							return true
						} else {
							now = ctime.Now()
							if es.IsAutocutter() {
								state = AutocutterErrState
							} else if es.IsUnrecoverable() {
								state = UnrecoverableErrState
							} else if es.IsAutoRecoverable() {
								state = AutoRecoverErrState
							} else if es.IsRecoverable() {
								state = RecoverableErrState
							}
						}
					} else if os.IsFedByButton() {
						state = FedByButtonState
					} else {
						state = OfflineState
					}
				}
			} else if ps.IsRecovery() {
				state = WaitRecoveryState
			}
			s.data.setState(now, state)

			if rs, err := s.Controller.RollStatus(); err != nil {
				s.connErr(false, ctime.Now(), err, "getting roll status")
				return true
			} else {
				s.data.setNearEnd(ctime.Now(), rs.IsNearEnd())
			}
		}
	}
	return true
}

func (s *Scanner) disStep(cmdCh <-chan []Cmd) bool {
	if s.data.InErr {
		select {
		case list, ok := <-cmdCh:
			if !ok {
				return false
			}
			now := ctime.Now()
			for _, cmd := range list {
				s.data.purge(now, cmd)
			}
			return true
		case <-s.ticker.C:
			// just wait
		}
	} else {
		<-s.ticker.C
	}

	if err := s.Controller.Reset(); err != nil {
		s.disErr(err, "first reset")
		return true
	}

	if ps, err := s.Controller.PrinterStatus(); err != nil {
		s.disErr(err, "getting printer status")
		return true
	} else {
		now := ctime.Now()
		state := OnlineState
		if ps.IsOffline() {
			if os, err := s.Controller.OfflineStatus(); err != nil {
				s.disErr(err, "getting offline status")
				return true
			} else {
				now = ctime.Now()
				if os.IsCoverOpen() {
					state = CoverOpenState
				} else if os.IsPaperEnd() {
					state = NoPaperState
				} else if os.IsError() {
					if es, err := s.Controller.ErrorStatus(); err != nil {
						s.disErr(err, "getting error status")
						return true
					} else {
						now = ctime.Now()
						if es.IsAutocutter() {
							state = AutocutterErrState
						} else if es.IsUnrecoverable() {
							state = UnrecoverableErrState
						} else if es.IsAutoRecoverable() {
							state = AutoRecoverErrState
						} else if es.IsRecoverable() {
							state = RecoverableErrState
						}
					}
				} else if os.IsFedByButton() {
					state = FedByButtonState
				} else {
					state = OfflineState
				}
			}
		} else if ps.IsRecovery() {
			state = WaitRecoveryState
		}

		if rs, err := s.Controller.RollStatus(); err != nil {
			s.disErr(err, "getting roll status")
			return true
		} else {
			s.data.setConnected(now, true)
			s.data.finishPrint(now)
			s.data.setDrawer(now, ps.IsDrawerPin3())
			s.data.setFeedBtn(now, ps.IsFeedButton())
			s.data.setState(now, state)
			s.data.setNearEnd(ctime.Now(), rs.IsNearEnd())
			if s.data.InErr {
				s.data.InErr = false
				s.ticker.Reset(s.Config.Interval)
			}
		}
	}
	return true
}

func (s *Scanner) connErr(close bool, now time.Time, err error, when string) {
	debugLog("ERR %s when %s", err, when)
	if close && !s.Controller.IsClosed() {
		s.Controller.Close()
	}
	if s.Controller.IsClosed() {
		s.data.setConnected(now, false)
	}
}

func (s *Scanner) disErr(err error, when string) {
	debugLog("ERR %s when %s", err, when)
	if !s.Controller.IsClosed() {
		s.Controller.Close()
	}
	if !s.data.InErr {
		s.data.InErr = true
		s.ticker.Reset(s.Config.ErrDelay)
	}
}

func (d *scannerData) setConnected(now time.Time, x bool) {
	if d.Connected != x {
		d.Cache = append(d.Cache, ConnectionEvent{now, x})
		d.Connected = x
	}
}

func (d *scannerData) startPrint(now time.Time, cmd PrintCmd) {
	d.Cache = append(d.Cache, StartPrintEvent{now, cmd})
	d.Print = cmd
}

func (d *scannerData) finishPrint(now time.Time) {
	if d.Print != nil {
		d.Cache = append(d.Cache, FinishPrintEvent{now, d.Print})
		d.Print = nil
	}
}

func (d *scannerData) purge(now time.Time, cmd Cmd) {
	debugLog("PURGE %s", cmd)
	d.Cache = append(d.Cache, CmdEvent{now, cmd, CmdRes{DisconnectErr}})
}

func (d *scannerData) cmdRes(now time.Time, cmd Cmd, res Res) {
	d.Cache = append(d.Cache, CmdEvent{now, cmd, res})
}

func (d *scannerData) setDrawer(now time.Time, x bool) {
	if d.Drawer != x {
		d.Cache = append(d.Cache, DrawerEvent{now, x})
		d.Drawer = x
	}
}

func (d *scannerData) setFeedBtn(now time.Time, x bool) {
	if d.FeedBtn != x {
		d.Cache = append(d.Cache, FeedButtonEvent{now, x})
		d.FeedBtn = x
	}
}

func (d *scannerData) setState(now time.Time, x State) {
	if d.State != x {
		d.Cache = append(d.Cache, StateEvent{now, x})
		d.State = x
	}
}

func (d *scannerData) setNearEnd(now time.Time, x bool) {
	if d.NearEnd != x {
		d.Cache = append(d.Cache, PaperNearEndEvent{now, x})
		d.NearEnd = x
	}
}
