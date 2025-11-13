package escpos

import (
	"errors"
	"io"

	. "github.com/hanindo/util/v2"
)

type Scanner struct {
	Controller *Controller
	Config     *ScannerConfig

	data  scannerData
	cache []Event
	inErr bool
	prn   []byte
	timer *Timer
	wait  bool
}

type scannerData struct {
	Connected bool
	State     State
	Drawer    bool
	FeedBtn   bool
	NearEnd   bool
}

func (s *Scanner) Scan(cmdCh <-chan []Cmd) <-chan []Event {
	if s.Config == nil {
		s.Config = DefaultScannerConfig()
	}
	s.data.State = OFFLINE
	if s.timer == nil {
		s.timer = ctime.NewTimer(0)
		<-s.timer.C
		s.wait = false
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
	if s.data.Connected {
		s.cache = append(s.cache, ConnectionEvent{ctime.Now(), false})
		s.data.Connected = false
	}
	s.flush(evtCh)
}

func (s *Scanner) flush(evtCh chan []Event) {
	if len(s.cache) > 0 {
		list := make([]Event, len(s.cache))
		copy(list, s.cache)
		evtCh <- list
		s.cache = s.cache[:0]
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
				rasb := false
				if pc, ok := cmd.(PrintCmd); ok {
					b := []byte(pc)
					dis = NeedDisablePulseLevel(b)
					s.add(StartPrintEvent{ctime.Now(), b})
					s.prn = b
					rasb = true
				}
				if dis {
					if err := s.Controller.DisablePulseLevel(); err != nil {
						debugLog("ERR %s when disable pulse level", err)
						if !s.Controller.IsClosed() {
							s.Controller.Close()
						}
						s.data.Connected = false
						s.add(ConnectionEvent{ctime.Now(), false})
						debugLog("PURGE %s", cmd)
						s.add(CmdEvent{
							ctime.Now(), cmd, CmdRes{DisconnectErr},
						})
						s.timer.Reset(s.Config.ErrDelay)
						s.wait = true
						continue
					}
				}
				debugLog("EXEC %s", cmd)
				res := cmd.exec(s.Controller)
				now := ctime.Now()
				s.add(CmdEvent{now, cmd, res})
				if err := res.Error(); err != nil {
					debugLog("ERR %s when exec %s", err, cmd)
					if s.Controller.IsClosed() {
						s.data.Connected = false
						s.add(ConnectionEvent{now, false})
						s.timer.Reset(s.Config.ErrDelay)
						s.wait = true
					}
				}
				if s.data.Connected && dis {
					if err := s.Controller.EnablePulseLevel(); err != nil {
						debugLog("ERR %s when enable pulse level", err)
						if s.Controller.IsClosed() {
							s.data.Connected = false
							s.add(ConnectionEvent{ctime.Now(), false})
							s.timer.Reset(s.Config.ErrDelay)
							s.wait = true
						}
					}
				}
				if s.data.Connected && rasb {
					if err := s.Controller.StartASB(ASB_ALL); err != nil {
						debugLog("ERR %s when start ASB", err)
						if s.Controller.IsClosed() {
							s.data.Connected = false
							s.add(ConnectionEvent{ctime.Now(), false})
							s.timer.Reset(s.Config.ErrDelay)
							s.wait = true
						}
					}
				}
			} else {
				debugLog("PURGE %s", cmd)
				s.add(CmdEvent{ctime.Now(), cmd, CmdRes{DisconnectErr}})
			}
		}
	case <-s.timer.C:
		if err := s.Controller.RestartASB(); err != nil {
			debugLog("ERR %s when restart ASB", err)
			if err == io.ErrUnexpectedEOF {
				s.Controller.Close()
			}
			if s.Controller.IsClosed() {
				s.data.Connected = false
				s.add(ConnectionEvent{ctime.Now(), false})
				s.timer.Reset(s.Config.ErrDelay)
				s.wait = true
			}
			return true
		}
		s.timer.Reset(s.Config.Ping)
		s.wait = true
	default:
		if list, err := s.Controller.GetASBs(); err != nil {
			debugLog("ERR %s when get ASB", err)
			if s.Controller.IsClosed() {
				s.data.Connected = false
				s.add(ConnectionEvent{ctime.Now(), false})
				s.timer.Reset(s.Config.ErrDelay)
				s.wait = true
			}
		} else {
			s.processASBs(list)
		}
	}
	return true
}

var DisconnectErr = errors.New("Printer disconnected")

func (s *Scanner) disStep(cmdCh <-chan []Cmd) bool {
	if s.inErr {
		select {
		case list, ok := <-cmdCh:
			if !ok {
				return false
			}
			now := ctime.Now()
			res := CmdRes{DisconnectErr}
			for _, cmd := range list {
				debugLog("PURGE %s", cmd)
				s.add(CmdEvent{now, cmd, res})
			}
			return true
		case <-s.timer.C:
			// just wait
			s.wait = false
		}
	}
	if err := s.Controller.Reset(); err != nil {
		debugLog("ERR %s when first reset", err)
		if !s.Controller.IsClosed() {
			s.Controller.Close()
		}
		s.inErr = true
		s.timer.Reset(s.Config.ErrDelay)
		s.wait = true
		return true
	}

	now := ctime.Now()
	if !s.Controller.IsASB() {
		if err := s.Controller.StartASB(ASB_ALL); err != nil {
			debugLog("ERR %s when start ASB", err)
			if !s.Controller.IsClosed() {
				s.Controller.Close()
			}
			s.inErr = true
			s.timer.Reset(s.Config.ErrDelay)
			s.wait = true
			return true
		}
	}
	list, err := s.Controller.GetASBs()
	if err != nil {
		debugLog("ERR %s when get ASB", err)
		if !s.Controller.IsClosed() {
			s.Controller.Close()
		}
		s.inErr = true
		s.timer.Reset(s.Config.ErrDelay)
		s.wait = true
		return true
	}
	s.data.Connected = true
	s.add(ConnectionEvent{now, true})
	s.processASBs(list)
	if s.Config.Ping > 0 {
		s.timer.Reset(s.Config.Ping)
		s.wait = true
	}
	return true
}

func (s *Scanner) processASBs(list []ASB) {
	restart := false
	for _, asb := range list {
		if s.prn != nil {
			s.add(FinishPrintEvent{ctime.Now(), s.prn})
			s.prn = nil
		}

		if state := asb.AS.State(); s.data.State != state {
			s.add(StateEvent{asb.Time, state})
			s.data.State = state
			if state == UNRECOVERABLE_ERR {
				restart = true
			}
		}
		if drawer := asb.AS.IsDrawerPin3(); s.data.Drawer != drawer {
			s.add(DrawerEvent{asb.Time, drawer})
			s.data.Drawer = drawer
		}
		if feedbtn := asb.AS.IsFeedButton(); s.data.FeedBtn != feedbtn {
			s.add(FeedButtonEvent{asb.Time, feedbtn})
			s.data.FeedBtn = feedbtn
		}
		if nearend := asb.AS.IsPaperNearEnd(); s.data.NearEnd != nearend {
			s.add(PaperNearEndEvent{asb.Time, nearend})
			s.data.NearEnd = nearend
		}
	}
	if restart {
		if err := s.Controller.RestartASB(); err != nil {
			debugLog("ERR %s when restart ASB", err)
			if err == io.ErrUnexpectedEOF {
				s.Controller.Close()
			}
			if s.Controller.IsClosed() {
				s.data.Connected = false
				s.add(ConnectionEvent{ctime.Now(), false})
				s.timer.Reset(s.Config.ErrDelay)
				s.wait = true
			}
		}
	}
}

func (s *Scanner) add(evt Event) {
	s.cache = append(s.cache, evt)
}
