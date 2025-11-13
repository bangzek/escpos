package escpos

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/albenik/go-serial/v2"
)

type Dev interface {
	UseXonXoff() bool
	Open() (io.ReadWriteCloser, error)
}

//----------------------------------------------------------------------

type SerialDev struct {
	Device string
	Config *SerialConfig
}

func (d *SerialDev) UseXonXoff() bool {
	d.init()
	return d.Config.UseXonXoff
}

func (d *SerialDev) Open() (io.ReadWriteCloser, error) {
	d.init()
	log("Opening %s", d.Device)
	return serial.Open(d.Device,
		serial.WithBaudrate(d.Config.Baudrate),
		serial.WithDataBits(d.Config.DataBits),
		serial.WithParity(serial.Parity(d.Config.Parity)),
		serial.WithReadTimeout(int(d.Config.ReadTimeout.Milliseconds())),
		serial.WithWriteTimeout(int(d.Config.ReadTimeout.Milliseconds())),
	)
}

func (d *SerialDev) init() {
	if d.Config == nil {
		d.Config = DefaultSerialConfig()
	}
}

//----------------------------------------------------------------------

type FileDev struct {
	Device string
	Config *DevConfig
}

func (d *FileDev) UseXonXoff() bool {
	return false
}

func (d *FileDev) Open() (io.ReadWriteCloser, error) {
	if d.Config == nil {
		d.Config = DefaultDevConfig()
	}
	log("Opening %s", d.Device)
	file, err := os.OpenFile(d.Device, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		return nil, err
	}
	err = file.SetWriteDeadline(time.Now().Add(d.Config.WriteTimeout))
	if err != nil {
		return nil, err
	}
	return &deadlined{file, d.Config}, nil
}

type deadlined struct {
	file *os.File
	conf *DevConfig
}

func (d *deadlined) Read(b []byte) (int, error) {
	dl := d.conf.ReadTimeout * time.Duration(len(b))
	err := d.file.SetReadDeadline(time.Now().Add(dl))
	if err != nil {
		return 0, err
	}

	n, err := d.file.Read(b)
	// Ignore deadline error and EOF
	if errors.Is(err, os.ErrDeadlineExceeded) || err == io.EOF {
		err = nil
	}
	return n, err
}

func (d *deadlined) Write(b []byte) (int, error) {
	dl := d.conf.WriteTimeout * time.Duration(len(b))
	err := d.file.SetWriteDeadline(time.Now().Add(dl))
	if err != nil {
		return 0, err
	}

	n, err := d.file.Write(b)
	// Ignore deadline error
	if errors.Is(err, os.ErrDeadlineExceeded) {
		err = nil
	}
	return n, err
}

func (d *deadlined) Close() error {
	return d.file.Close()
}
