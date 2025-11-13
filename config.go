package escpos

import "time"

const (
	PING      = time.Second
	ERR_DELAY = 5 * time.Second
	TIMEOUT   = time.Second
	WAIT      = 100 * time.Millisecond
	ASB_CACHE = 25 * time.Millisecond

	READ_TIMEOUT  = 50 * time.Millisecond
	WRITE_TIMEOUT = 50 * time.Millisecond

	BAUDRATE  = 9600
	DATA_BITS = 8
	PARITY    = NoParity
)

//----------------------------------------------------------------------

type ScannerConfig struct {
	Ping     time.Duration
	ErrDelay time.Duration
}

func DefaultScannerConfig() *ScannerConfig {
	return &ScannerConfig{
		Ping:     PING,
		ErrDelay: ERR_DELAY,
	}
}

//----------------------------------------------------------------------

type ControllerConfig struct {
	Timeout  time.Duration
	Wait     time.Duration
	ASBCache time.Duration
}

func DefaultControllerConfig() *ControllerConfig {
	return &ControllerConfig{
		Timeout:  TIMEOUT,
		Wait:     WAIT,
		ASBCache: ASB_CACHE,
	}
}

//----------------------------------------------------------------------

type DevConfig struct {
	ReadTimeout  time.Duration // This is per byte timeout
	WriteTimeout time.Duration // This is per byte timeout
}

func DefaultDevConfig() *DevConfig {
	return &DevConfig{
		ReadTimeout:  READ_TIMEOUT,
		WriteTimeout: WRITE_TIMEOUT,
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
		Baudrate:   BAUDRATE,
		DataBits:   DATA_BITS,
		Parity:     PARITY,
		DevConfig: DevConfig{
			ReadTimeout:  READ_TIMEOUT,
			WriteTimeout: WRITE_TIMEOUT,
		},
	}
}
