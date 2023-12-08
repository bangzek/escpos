package escpos

import (
	"io"
	"time"
)

const (
	buff_LEN = 129
	CHUNK    = 40
)

type Controller struct {
	Dev    Dev
	Config *ControllerConfig

	dev  io.ReadWriteCloser
	buff []byte
}

func (c *Controller) Close() {
	if c.dev != nil {
		log("Closing connection")
		c.dev.Close()
		c.dev = nil
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
	return nil
}

// Disable both real-time pulse command (DLE DC4 1 or 2).
func (c *Controller) DisablePulseLevel() error {
	return c.Send([]byte{GS, '(', 'D', 5, 0, 20, 1, 1, 2, 1})
}

// Enable both real-time pulse command (DLE DC4 1 or 2).
func (c *Controller) EnablePulseLevel() error {
	return c.Send([]byte{GS, '(', 'D', 5, 0, 20, 1, 0, 2, 0})
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
					debugLog("%d/%d ~RX: % X", i+1, p, c.buff[0])
					cont = false
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
	n, err := c.dev.Read(c.buff[:l])
	if err != nil {
		return n, err
	}

	if n == 0 {
		timeout := ctime.Now().Add(c.Config.Timeout)
		for n == 0 && ctime.Now().Before(timeout) {
			if n, err = c.dev.Read(c.buff[:l]); err != nil {
				return n, err
			}
		}
	}

	if n > 0 {
		debugLog("RX: % X", c.buff[:n])
		if c.Dev.UseXonXoff() {
			for tn := trimXonXoff(c.buff[:n]); tn < n; {
				n = tn
				if nn, err := c.dev.Read(c.buff[n:l]); err != nil {
					return n + nn, err
				} else if nn > 0 {
					debugLog("RX:%d> % X", n, c.buff[:n+nn])
					tn += trimXonXoff(c.buff[n : n+nn])
					n += nn
				}
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
