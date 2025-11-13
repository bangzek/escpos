package escpos

import (
	"io"
	"regexp"
	"time"

	. "github.com/hanindo/util/v2"
)

const (
	buff_LEN = 129
	CHUNK    = 40
)

var (
	ctime = NewClock()
)

type Controller struct {
	Dev    Dev
	Config *ControllerConfig

	dev    io.ReadWriteCloser
	buff   []byte
	isASB  bool
	asbReq ASBReq
	tASB   time.Time
	asb    []ASB
}

func (c *Controller) Close() {
	if c.dev != nil {
		log("Closing connection")
		c.dev.Close()
		c.dev = nil
		c.isASB = false
	}
}

func (c *Controller) IsClosed() bool {
	return c.dev == nil
}

func (c *Controller) Send(b []byte) error {
	if err := c.init(); err != nil {
		c.Close()
		return err
	}
	err := c.write(b)
	if err != nil {
		c.Close()
	} else {
		time.Sleep(c.Config.Wait)
	}
	return err
}

func (c *Controller) Reset() error {
	if err := c.Send([]byte{ESC, '@'}); err != nil {
		return err
	}
	if c.isASB {
		return c.StartASB(c.asbReq)
	}
	return nil
}

const (
	// DISABLE: GS ( D 5 0 20 1 1 2 1
	//  ENABLE: GS ( D 5 0 20 1 0 2 0

	DISABLE_PULSE_LEVEL = "\x1D(D\x05\x00\x14\x01\x01\x02\x01"
	ENABLE_PULSE_LEVEL  = "\x1D(D\x05\x00\x14\x01\x00\x02\x00"
)

// DLE DC4 0 or 1
var pulseLvlRe = regexp.MustCompile("\x10\x14[\x00\x01]")

func NeedDisablePulseLevel(b []byte) bool {
	return pulseLvlRe.Match(b)
}

// Disable both real-time pulse command (DLE DC4 1 or 2).
func (c *Controller) DisablePulseLevel() error {
	return c.Send([]byte(DISABLE_PULSE_LEVEL))
}

// Enable both real-time pulse command (DLE DC4 1 or 2).
func (c *Controller) EnablePulseLevel() error {
	return c.Send([]byte(ENABLE_PULSE_LEVEL))
}

// Press Feed Button in real-time.
// Note: support depend on printer type.
func (c *Controller) PressFeedButton() error {
	return c.Send([]byte{DLE, ENQ, 0})
}

// Recover from recoverable error and restart printing from the line
// where the error occured in real-time.
// Note: support depend on printer type.
func (c *Controller) Recover() error {
	return c.Send([]byte{DLE, ENQ, 1})
}

// Recover from recoverable error after clearing the receive and print buffers
// in real-time.
func (c *Controller) ClearAndRecover() error {
	return c.Send([]byte{DLE, ENQ, 2})
}

// Pulse the pin 2 of drawer connector for (tOn * 2ms) ON and (tOff * 2ms) OFF.
// Most printer force tOff >= tOn and some has minimal tOn and/or tOff.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin2(tOn, tOff byte) error {
	return c.Send([]byte{ESC, 'p', 0, tOn, tOff})
}

// Pulse the pin 5 of drawer connector for (tOn * 2ms) ON and (tOff * 2ms) OFF.
// Most printer force tOff >= tOn and some has minimal tOn and/or tOff.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin5(tOn, tOff byte) error {
	return c.Send([]byte{ESC, 'p', 1, tOn, tOff})
}

// Pulse the pin 2 of drawer connector for 100ms ON and 100ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin2Level1() error {
	return c.Send([]byte{DLE, DC4, 1, 0, 1})
}

// Pulse the pin 2 of drawer connector for 200ms ON and 200ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin2Level2() error {
	return c.Send([]byte{DLE, DC4, 1, 0, 2})
}

// Pulse the pin 2 of drawer connector for 300ms ON and 300ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin2Level3() error {
	return c.Send([]byte{DLE, DC4, 1, 0, 3})
}

// Pulse the pin 2 of drawer connector for 400ms ON and 400ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin2Level4() error {
	return c.Send([]byte{DLE, DC4, 1, 0, 4})
}

// Pulse the pin 2 of drawer connector for 500ms ON and 500ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin2Level5() error {
	return c.Send([]byte{DLE, DC4, 1, 0, 5})
}

// Pulse the pin 2 of drawer connector for 600ms ON and 600ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin2Level6() error {
	return c.Send([]byte{DLE, DC4, 1, 0, 6})
}

// Pulse the pin 2 of drawer connector for 700ms ON and 700ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin2Level7() error {
	return c.Send([]byte{DLE, DC4, 1, 0, 7})
}

// Pulse the pin 2 of drawer connector for 800ms ON and 800ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin2Level8() error {
	return c.Send([]byte{DLE, DC4, 1, 0, 8})
}

// Pulse the pin 5 of drawer connector for 100ms ON and 100ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin5Level1() error {
	return c.Send([]byte{DLE, DC4, 1, 1, 1})
}

// Pulse the pin 5 of drawer connector for 200ms ON and 200ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin5Level2() error {
	return c.Send([]byte{DLE, DC4, 1, 1, 2})
}

// Pulse the pin 5 of drawer connector for 300ms ON and 300ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin5Level3() error {
	return c.Send([]byte{DLE, DC4, 1, 1, 3})
}

// Pulse the pin 5 of drawer connector for 400ms ON and 400ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin5Level4() error {
	return c.Send([]byte{DLE, DC4, 1, 1, 4})
}

// Pulse the pin 5 of drawer connector for 500ms ON and 500ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin5Level5() error {
	return c.Send([]byte{DLE, DC4, 1, 1, 5})
}

// Pulse the pin 5 of drawer connector for 600ms ON and 600ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin5Level6() error {
	return c.Send([]byte{DLE, DC4, 1, 1, 6})
}

// Pulse the pin 5 of drawer connector for 700ms ON and 700ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin5Level7() error {
	return c.Send([]byte{DLE, DC4, 1, 1, 7})
}

// Pulse the pin 5 of drawer connector for 800ms ON and 800ms OFF in real-time.
// From Epson docs it's not possible to pulse both pin at the same time.
func (c *Controller) PulseDrawerPin5Level8() error {
	return c.Send([]byte{DLE, DC4, 1, 1, 8})
}

// Enable Automatic Status Back.
func (c *Controller) StartASB(req ASBReq) error {
	if err := c.Send([]byte{GS, 'a', byte(req)}); err != nil {
		return err
	}

	if !c.isASB {
		c.asb = nil
		if n, err := c.read(0); err != nil {
			c.Close()
			return err
		} else {
			if s := AS(c.buff[:n]); !s.IsValid() {
				// disable it
				c.Send([]byte{GS, 'a', 0})
				return io.ErrUnexpectedEOF
			}
		}
		c.tASB = ctime.Now()
		c.isASB = true
	}
	c.asbReq = req

	return nil
}

// Force Automatic Status Back.
func (c *Controller) RestartASB() error {
	c.isASB = false
	return c.StartASB(c.asbReq)
}

// Disable Automatic Status Back.
func (c *Controller) StopASB() error {
	if c.isASB {
		// Read again to clear the buffer
		if _, err := c.read(0); err != nil {
			c.Close()
			return err
		}

		if err := c.Send([]byte{GS, 'a', 0}); err != nil {
			return err
		}
		c.isASB = false
	}
	return nil
}

// Is Automatic Status Back Enabled.
func (c *Controller) IsASB() bool {
	return c.isASB
}

// Get current ASB.
func (c *Controller) GetASBs() ([]ASB, error) {
	if !c.isASB {
		if err := c.StartASB(ASB_ALL); err != nil {
			return nil, err
		}
	} else if ctime.Now().Sub(c.tASB) > c.Config.ASBCache {
		if _, err := c.read(0); err != nil {
			c.Close()
			return nil, err
		}
	}

	// trim c.asb
	l := c.asb
	c.asb = nil
	return l, nil
}

// Get real-time printer status.
func (c *Controller) PrinterStatus() (*PrinterStatus, error) {
	if err := c.Send([]byte{DLE, EOT, 1}); err != nil {
		return nil, err
	}

	if n, err := c.read(1); err != nil {
		c.Close()
		return nil, err
	} else {
		for i := 0; i < n; i++ {
			if s := PrinterStatus(c.buff[i]); s.IsValid() {
				return &s, nil
			}
		}
		return nil, io.ErrUnexpectedEOF
	}
}

// Get real-time offline status.
func (c *Controller) OfflineStatus() (*OfflineStatus, error) {
	if err := c.Send([]byte{DLE, EOT, 2}); err != nil {
		return nil, err
	}

	if n, err := c.read(1); err != nil {
		c.Close()
		return nil, err
	} else {
		for i := 0; i < n; i++ {
			if s := OfflineStatus(c.buff[i]); s.IsValid() {
				return &s, nil
			}
		}
		return nil, io.ErrUnexpectedEOF
	}
}

// Get real-time error status.
func (c *Controller) ErrorStatus() (*ErrorStatus, error) {
	if err := c.Send([]byte{DLE, EOT, 3}); err != nil {
		return nil, err
	}

	if n, err := c.read(1); err != nil {
		c.Close()
		return nil, err
	} else {
		for i := 0; i < n; i++ {
			if s := ErrorStatus(c.buff[i]); s.IsValid() {
				return &s, nil
			}
		}
		return nil, io.ErrUnexpectedEOF
	}
}

// Get real-time paper roll status.
func (c *Controller) RollStatus() (*RollStatus, error) {
	if err := c.Send([]byte{DLE, EOT, 4}); err != nil {
		return nil, err
	}

	if n, err := c.read(1); err != nil {
		c.Close()
		return nil, err
	} else {
		for i := 0; i < n; i++ {
			if s := RollStatus(c.buff[i]); s.IsValid() {
				return &s, nil
			}
		}
		return nil, io.ErrUnexpectedEOF
	}
}

func (c *Controller) PaperStatus() (*PaperStatus, error) {
	if err := c.Send([]byte{GS, 'r', 1}); err != nil {
		return nil, err
	}

	if n, err := c.read(1); err != nil {
		c.Close()
		return nil, err
	} else {
		for i := 0; i < n; i++ {
			if s := PaperStatus(c.buff[i]); s.IsValid() {
				return &s, nil
			}
		}
		return nil, io.ErrUnexpectedEOF
	}
}

func (c *Controller) DrawerStatus() (*DrawerStatus, error) {
	if err := c.Send([]byte{GS, 'r', 2}); err != nil {
		return nil, err
	}

	if n, err := c.read(1); err != nil {
		c.Close()
		return nil, err
	} else {
		for i := 0; i < n; i++ {
			if s := DrawerStatus(c.buff[i]); s.IsValid() {
				return &s, nil
			}
		}
		return nil, io.ErrUnexpectedEOF
	}
}

func (c *Controller) ModelID() (*ModelID, error) {
	if err := c.Send([]byte{GS, 'I', 1}); err != nil {
		return nil, err
	}

	if n, err := c.read(1); err != nil {
		c.Close()
		return nil, err
	} else {
		for i := 0; i < n; i++ {
			if id := ModelID(c.buff[i]); id.IsValid() {
				return &id, nil
			}
		}
		return nil, io.ErrUnexpectedEOF
	}
}

func (c *Controller) TypeID() (*TypeID, error) {
	if err := c.Send([]byte{GS, 'I', 2}); err != nil {
		return nil, err
	}

	if n, err := c.read(1); err != nil {
		c.Close()
		return nil, err
	} else {
		for i := 0; i < n; i++ {
			if id := TypeID(c.buff[i]); id.IsValid() {
				return &id, nil
			}
		}
		return nil, io.ErrUnexpectedEOF
	}
}

func (c *Controller) Firmware() (Info, error) {
	return c.getInfo('A')
}

func (c *Controller) Maker() (Info, error) {
	return c.getInfo('B')
}

func (c *Controller) Model() (Info, error) {
	return c.getInfo('C')
}

func (c *Controller) SerialNo() (Info, error) {
	return c.getInfo('D')
}

//----------------------------------------------------------------------

func (c *Controller) init() error {
	if c.dev == nil {
		var err error
		c.dev, err = c.Dev.Open()
		if err != nil {
			c.dev = nil
			return err
		}
	}
	if len(c.buff) == 0 {
		c.buff = make([]byte, buff_LEN)
	}
	if c.Config == nil {
		c.Config = DefaultControllerConfig()
	}
	return nil
}

func (c *Controller) write(b []byte) error {
	debugLog("TX: %q", string(b))
	if c.Dev.UseXonXoff() && len(b) > CHUNK {
		p := len(b) / CHUNK
		if len(b)%CHUNK > 0 {
			p++
		}
		for i := 0; i < p; i++ {
			j := i * CHUNK
			k := j + CHUNK
			if k > len(b) {
				k = len(b)
			}
			n, err := c.dev.Write(b[j:k])
			if err == nil && n < len(b[j:k]) {
				return io.ErrShortWrite
			}

			if n, err := c.dev.Read(c.buff[:1]); err != nil {
				return err
			} else if n == 0 {
				continue
			}

			cont := true
			var b []byte
			for cont {
				switch c.buff[0] {
				case XON:
					if n, err := c.dev.Read(c.buff[:1]); err != nil {
						return err
					} else if n == 0 {
						cont = false
					}
				case XOFF:
					if _, err := c.dev.Read(c.buff[:1]); err != nil {
						return err
					}
				default:
					if c.isASB &&
						(c.buff[0]&AS_MASK_1 == AS_FIXED_1 || len(b) > 0) {
						if len(b) == 0 {
							b = append(b, c.buff[0])
							debugLog("%d/%d RX: % X", i+1, p, b)
						} else {
							b = append(b, c.buff[0])
							debugLog("%d/%d RX:%d> % X", i+1, p, len(b)-1, b)
							if len(b) == 4 {
								if s := AS(b); s.IsValid() {
									c.tASB = ctime.Now()
									c.asb = append(c.asb, ASB{c.tASB, s})
								}
								b = nil
								cont = false
							}
						}
						if b != nil {
							if _, err := c.dev.Read(c.buff[:1]); err != nil {
								return err
							}
						}
					} else {
						debugLog("%d/%d ~RX: % X", i+1, p, c.buff[0])
					}
				}
			}
		}
		return nil
	} else {
		n, err := c.dev.Write(b)
		if err == nil && n < len(b) {
			return io.ErrShortWrite
		}
		return err
	}
}

func (c *Controller) read(l int) (int, error) {
	rl := l
	if l == 0 || (l == 1 && c.isASB) {
		rl += 4
		//rl += 128
	}
	n, err := c.dev.Read(c.buff[:rl])
	if err != nil {
		return n, err
	}

	if n == 0 && (l > 0 || !c.isASB) {
		timeout := ctime.Now().Add(c.Config.Timeout)
		for n == 0 && ctime.Now().Before(timeout) {
			if n, err = c.dev.Read(c.buff[:rl]); err != nil {
				return n, err
			}
		}
	}

	if n > 0 || l > 0 || !c.isASB {
		debugLog("RX: % X", c.buff[:n])
	}
	if n > 0 {
		if c.Dev.UseXonXoff() {
			t := ctime.Now()
			for {
				for tn := trimXonXoff(c.buff[:n]); tn < n; {
					n = tn
					if nn, err := c.dev.Read(c.buff[n:rl]); err != nil {
						return n + nn, err
					} else if nn > 0 {
						debugLog("RX:%d> % X", n, c.buff[:n+nn])
						tn += trimXonXoff(c.buff[n : n+nn])
						n += nn
					}
				}
				if (l == 0 && n > 0 && n < 4) ||
					(l == 1 && c.isASB &&
						((n > 1 && n < 5) ||
							(n == 1 && c.buff[0]&AS_MASK_1 == AS_FIXED_1))) {

					if nn, err := c.dev.Read(c.buff[n:rl]); err != nil {
						return n + nn, err
					} else if nn > 0 {
						debugLog("RX:%d> % X", n, c.buff[:n+nn])
						n += nn
						t = ctime.Now()
					} else if ctime.Now().Sub(t) >= c.Config.Timeout {
						break
					}
				} else {
					break
				}
			}
		}
	}

	if n >= 4 && (l == 0 || (l == 1 && c.isASB)) {
		for i := 0; i <= n-4; i++ {
			if s := AS(c.buff[i : i+4]); s.IsValid() {
				c.tASB = ctime.Now()
				c.asb = append(c.asb, ASB{c.tASB, s})
				if n > i+4 {
					copy(c.buff[i:i+4], c.buff[i+4:])
				}
				if c.isASB {
					n -= 4
				}
				break
			}
		}
	}
	return n, nil
}

func (c *Controller) getInfo(cmd byte) (Info, error) {
	if err := c.Send([]byte{GS, 'I', cmd}); err != nil {
		return nil, err
	}

	if n, err := c.read(80); err != nil {
		c.Close()
		return nil, err
	} else {
		if info := MakeInfo(c.buff[:n]); info.IsValid() {
			return info, nil
		}
		return nil, io.ErrUnexpectedEOF
	}
}

func trimXonXoff(b []byte) int {
	j := -1
	for i := 0; i < len(b); i++ {
		if b[i] != XOFF && b[i] != XON {
			j++
			if j < i {
				b[j] = b[i]
			}
		}
	}
	return j + 1
}
