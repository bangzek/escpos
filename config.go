package escpos

import "time"

const (
	Interval = time.Second / 2
	ErrDelay = 5 * time.Second
	Timeout  = time.Second
	Wait     = 100 * time.Millisecond

	ReadTimeout  = 50 * time.Millisecond
	WriteTimeout = 50 * time.Millisecond

	Baudrate = 9600
	DataBits = 8
)

//----------------------------------------------------------------------

type ScannerConfig struct {
	Interval time.Duration
	ErrDelay time.Duration
}

func DefaultScannerConfig() *ScannerConfig {
	return &ScannerConfig{
		Interval: Interval,
		ErrDelay: ErrDelay,
	}
}

//----------------------------------------------------------------------

type ControllerConfig struct {
	Timeout time.Duration
	Wait    time.Duration
}

func DefaultControllerConfig() *ControllerConfig {
	return &ControllerConfig{
		Timeout: Timeout,
		Wait:    Wait,
	}
}

//----------------------------------------------------------------------

type DevConfig struct {
	ReadTimeout  time.Duration // This is per byte timeout
	WriteTimeout time.Duration // This is per byte timeout
}

func DefaultDevConfig() *DevConfig {
	return &DevConfig{
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
	}
}

//----------------------------------------------------------------------

type SerialConfig struct {
	UseXonXoff bool
	Baudrate   int
	DataBits   int
	Parity     Parity
	DevConfig
}

func DefaultSerialConfig() *SerialConfig {
	return &SerialConfig{
		UseXonXoff: true,
		Baudrate:   Baudrate,
		DataBits:   DataBits,
		Parity:     NoParity,
		DevConfig: DevConfig{
			ReadTimeout:  ReadTimeout,
			WriteTimeout: WriteTimeout,
		},
	}
}
