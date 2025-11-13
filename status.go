package escpos

const (
	PRN_MASK     = 0b1001_0011
	PRN_FIXED    = 0b0001_0010
	PRN_DRAWER   = 0b0000_0100
	PRN_OFFLINE  = 0b0000_1000
	PRN_RECOVERY = 0b0010_0000
	PRN_FEED_BTN = 0b0100_0000

	OFF_MASK  = 0b1001_0011
	OFF_FIXED = 0b0001_0010
	OFF_COVER = 0b0000_0100
	OFF_FED   = 0b0000_1000
	OFF_PAPER = 0b0010_0000
	OFF_ERROR = 0b0100_0000

	ERR_MASK         = 0b1001_0011
	ERR_FIXED        = 0b0001_0010
	ERR_RECOVER      = 0b0000_0100
	ERR_AUTOCUTTER   = 0b0000_1000
	ERR_UNRECOVER    = 0b0010_0000
	ERR_AUTO_RECOVER = 0b0100_0000

	ROLL_MASK     = 0b1001_0011
	ROLL_FIXED    = 0b0001_0010
	ROLL_NEAR_END = 0b0000_1100
	ROLL_NO_PAPER = 0b0110_0000

	PAPER_MASK     = 0b1001_0000
	PAPER_FIXED    = 0b0000_0000
	PAPER_NEAR_END = 0b0000_0011
	PAPER_END      = 0b0000_1100

	DRAWER_MASK  = 0b1001_0000
	DRAWER_FIXED = 0b0000_0000
	DRAWER_PIN_3 = 0b0000_0001
)

//----------------------------------------------------------------------

type PrinterStatus byte

func (s PrinterStatus) IsValid() bool {
	return s&PRN_MASK == PRN_FIXED
}

func (s PrinterStatus) IsDrawerPin3() bool {
	return s&PRN_DRAWER == PRN_DRAWER
}

func (s PrinterStatus) IsOffline() bool {
	return s&PRN_OFFLINE == PRN_OFFLINE
}

// Is the printer waiting for online recovery?
func (s PrinterStatus) IsRecovery() bool {
	return s&PRN_RECOVERY == PRN_RECOVERY
}

// Is paper feed button being pressed
func (s PrinterStatus) IsFeedButton() bool {
	return s&PRN_FEED_BTN == PRN_FEED_BTN
}

func (s PrinterStatus) String() (str string) {
	if s.IsDrawerPin3() {
		str = "DRAWER"
	}
	if s.IsOffline() {
		if str != "" {
			str += ", "
		}
		str += "OFFLINE"
	}
	if s.IsRecovery() {
		if str != "" {
			str += ", "
		}
		str += "RECOVERY"
	}
	if s.IsFeedButton() {
		if str != "" {
			str += ", "
		}
		str += "FEED BTN"
	}
	if str == "" {
		str = "OK"
	}
	return str
}

//----------------------------------------------------------------------

type OfflineStatus byte

func (s OfflineStatus) IsValid() bool {
	return s&OFF_MASK == OFF_FIXED
}

func (s OfflineStatus) IsCoverOpen() bool {
	return s&OFF_COVER == OFF_COVER
}

// Is paper being fed by the paper feed button
func (s OfflineStatus) IsFedByButton() bool {
	return s&OFF_FED == OFF_FED
}

// Is printing stops due to paper-end
func (s OfflineStatus) IsPaperEnd() bool {
	return s&OFF_PAPER == OFF_PAPER
}

func (s OfflineStatus) IsError() bool {
	return s&OFF_ERROR == OFF_ERROR
}

func (s OfflineStatus) String() (str string) {
	if s.IsCoverOpen() {
		str = "COVER"
	}
	if s.IsFedByButton() {
		if str == "" {
			str += ", "
		}
		str += "FED BY"
	}
	if s.IsPaperEnd() {
		if str == "" {
			str += ", "
		}
		str += "PAPER"
	}
	if s.IsError() {
		if str == "" {
			str += ", "
		}
		str += "ERROR"
	}
	if str == "" {
		str = "ONLINE"
	}
	return str
}

//----------------------------------------------------------------------

type ErrorStatus byte

func (s ErrorStatus) IsValid() bool {
	return s&ERR_MASK == ERR_FIXED
}

func (s ErrorStatus) IsRecoverable() bool {
	return s&ERR_RECOVER == ERR_RECOVER
}

func (s ErrorStatus) IsAutocutter() bool {
	return s&ERR_AUTOCUTTER == ERR_AUTOCUTTER
}

func (s ErrorStatus) IsUnrecoverable() bool {
	return s&ERR_UNRECOVER == ERR_UNRECOVER
}

func (s ErrorStatus) IsAutoRecoverable() bool {
	return s&ERR_AUTO_RECOVER == ERR_AUTO_RECOVER
}

func (s ErrorStatus) String() (str string) {
	if s.IsRecoverable() {
		str = "RECOVER"
	}
	if s.IsAutocutter() {
		if str == "" {
			str += ", "
		}
		str += "AUTOCUTTER"
	}
	if s.IsUnrecoverable() {
		if str == "" {
			str += ", "
		}
		str += "UNRECOVER"
	}
	if s.IsAutoRecoverable() {
		if str == "" {
			str += ", "
		}
		str += "AUTO RECOVER"
	}
	if str == "" {
		str = "NO ERROR"
	}
	return str
}

//----------------------------------------------------------------------

type RollStatus byte

func (s RollStatus) IsValid() bool {
	return s&ROLL_MASK == ROLL_FIXED &&
		(s&0x04 == 0x04) == (s&0x08 == 0x08) &&
		(s&0x20 == 0x20) == (s&0x40 == 0x40)
}

func (s RollStatus) IsNearEnd() bool {
	return s&ROLL_NEAR_END == ROLL_NEAR_END
}

func (s RollStatus) IsNoPaper() bool {
	return s&ROLL_NO_PAPER == ROLL_NO_PAPER
}

func (s RollStatus) String() (str string) {
	if s.IsNearEnd() {
		str = "NEAR END"
	}
	if s.IsNoPaper() {
		if str == "" {
			str += ", "
		}
		str += "NO PAPER"
	}
	if str == "" {
		str = "ROLL OK"
	}
	return str
}

//----------------------------------------------------------------------

type PaperStatus byte

func (s PaperStatus) IsValid() bool {
	return s&PAPER_MASK == PAPER_FIXED &&
		(s&0x01 == 0x01) == (s&0x02 == 0x02) &&
		(s&0x04 == 0x04) == (s&0x08 == 0x08)
}

func (s PaperStatus) IsNearEnd() bool {
	return s&PAPER_NEAR_END == PAPER_NEAR_END
}

func (s PaperStatus) IsNoPaper() bool {
	return s&PAPER_END == PAPER_END
}

func (s PaperStatus) String() (str string) {
	if s.IsNearEnd() {
		str = "NEAR END"
	}
	if s.IsNoPaper() {
		if str == "" {
			str += ", "
		}
		str += "NO PAPER"
	}
	if str == "" {
		str = "PAPER OK"
	}
	return str
}

//----------------------------------------------------------------------

type DrawerStatus byte

func (s DrawerStatus) IsValid() bool {
	return s&DRAWER_MASK == DRAWER_FIXED
}

func (s DrawerStatus) IsPin3() bool {
	return s&DRAWER_PIN_3 == DRAWER_PIN_3
}

func (s DrawerStatus) String() string {
	if s.IsPin3() {
		return "DRAWER ON"
	} else {
		return "DRAWER OFF"
	}
}
