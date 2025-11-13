package escpos

import "time"

type ASBReq byte

const (
	ASB_DRAWER = 0b0000_0001
	ASB_ONLINE = 0b0000_0010
	ASB_ERROR  = 0b0000_0100
	ASB_PAPER  = 0b0000_1000
	ASB_BUTTON = 0b0100_0000

	ASB_ALL = ASB_DRAWER | ASB_ONLINE | ASB_ERROR | ASB_PAPER | ASB_BUTTON

	AS_MASK_1    = 0b1001_0011
	AS_FIXED_1   = 0b0001_0000
	AS_1_DRAWER  = 0b0000_0100
	AS_1_OFFLINE = 0b0000_1000
	AS_1_COVER   = 0b0010_0000
	AS_1_FED     = 0b0100_0000

	AS_MASK_2         = 0b1001_0000
	AS_FIXED_2        = 0b0000_0000
	AS_2_RECOVERY     = 0b0000_0001
	AS_2_FEED_BTN     = 0b0000_0010
	AS_2_RECOVER      = 0b0000_0100
	AS_2_AUTOCUTTER   = 0b0000_1000
	AS_2_UNRECOVER    = 0b0010_0000
	AS_2_AUTO_RECOVER = 0b0100_0000

	AS_MASK_3     = 0b1001_0000
	AS_FIXED_3    = 0b0000_0000
	AS_3_NEAR_END = 0b0000_0011
	AS_3_NO_PAPER = 0b0000_1100

	AS_MASK_4  = 0b1001_0000
	AS_FIXED_4 = 0b0000_0000
)

type ASB struct {
	Time time.Time
	AS   AS
}

func (s ASB) String() (str string) {
	return s.Time.Format("15:04:05.000 ") + s.AS.String()
}

type AS []byte

func (s AS) Clone() AS {
	if len(s) == 0 {
		return nil
	}

	n := AS(make([]byte, 0, len(s)))
	return AS(append(n, s...))
}

func (s AS) IsValid() bool {
	return len(s) == 4 &&
		s[0]&AS_MASK_1 == AS_FIXED_1 &&
		s[1]&AS_MASK_2 == AS_FIXED_2 &&
		s[2]&AS_MASK_3 == AS_FIXED_3 &&
		(s[2]&0x01 == 0x01) == (s[2]&0x02 == 0x02) &&
		(s[2]&0x04 == 0x04) == (s[2]&0x08 == 0x08) &&
		s[3]&AS_MASK_4 == AS_FIXED_4
}

func (s AS) State() State {
	if s.IsCoverOpen() {
		return COVER_IS_OPEN
	} else if s.IsNoPaper() {
		return NO_PAPER
	} else if s.IsAutocutterError() {
		return AUTOCUTTER_ERR
	} else if s.IsUnrecoverableError() {
		return UNRECOVERABLE_ERR
	} else if s.IsAutoRecoverableError() {
		return AUTO_RECOVER_ERR
	} else if s.IsRecoverableError() {
		return RECOVERABLE_ERR
	} else if s.IsFedByButton() {
		return FED_BY_BUTTON
	} else if s.IsRecovery() {
		return WAIT_RECOVERY
	} else if s.IsOffline() {
		return OFFLINE
	} else {
		return ONLINE
	}
}

func (s AS) IsDrawerPin3() bool {
	return s[0]&AS_1_DRAWER == AS_1_DRAWER
}

func (s AS) IsOffline() bool {
	return s[0]&AS_1_OFFLINE == AS_1_OFFLINE
}

func (s AS) IsCoverOpen() bool {
	return s[0]&AS_1_COVER == AS_1_COVER
}

// Is paper being fed by the paper feed button
func (s AS) IsFedByButton() bool {
	return s[0]&AS_1_FED == AS_1_FED
}

// Is it waiting for online recovery?
func (s AS) IsRecovery() bool {
	return s[1]&AS_2_RECOVERY == AS_2_RECOVERY
}

// Is paper feed button being pressed
func (s AS) IsFeedButton() bool {
	return s[1]&AS_2_FEED_BTN == AS_2_FEED_BTN
}

func (s AS) IsRecoverableError() bool {
	return s[1]&AS_2_RECOVER == AS_2_RECOVER
}

func (s AS) IsAutocutterError() bool {
	return s[1]&AS_2_AUTOCUTTER == AS_2_AUTOCUTTER
}

func (s AS) IsUnrecoverableError() bool {
	return s[1]&AS_2_UNRECOVER == AS_2_UNRECOVER
}

func (s AS) IsAutoRecoverableError() bool {
	return s[1]&AS_2_AUTO_RECOVER == AS_2_AUTO_RECOVER
}

func (s AS) IsPaperNearEnd() bool {
	return s[2]&AS_3_NEAR_END == AS_3_NEAR_END
}

func (s AS) IsNoPaper() bool {
	return s[2]&AS_3_NO_PAPER == AS_3_NO_PAPER
}

func (s AS) String() (str string) {
	if s.IsDrawerPin3() {
		str = "DRAWER"
	}
	if s.IsOffline() {
		if str != "" {
			str += ", "
		}
		str += "OFFLINE"
	}
	if s.IsCoverOpen() {
		if str != "" {
			str += ", "
		}
		str += "COVER"
	}
	if s.IsFedByButton() {
		if str != "" {
			str += ", "
		}
		str += "FED BY"
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
	if s.IsRecoverableError() {
		if str != "" {
			str += ", "
		}
		str += "RECOVER"
	}
	if s.IsAutocutterError() {
		if str != "" {
			str += ", "
		}
		str += "AUTOCUTTER"
	}
	if s.IsUnrecoverableError() {
		if str != "" {
			str += ", "
		}
		str += "UNRECOVER"
	}
	if s.IsAutoRecoverableError() {
		if str != "" {
			str += ", "
		}
		str += "AUTO RECOVER"
	}
	if s.IsPaperNearEnd() {
		if str != "" {
			str += ", "
		}
		str += "NEAR END"
	}
	if s.IsNoPaper() {
		if str != "" {
			str += ", "
		}
		str += "NO PAPER"
	}
	if str == "" {
		str = "OK"
	}
	return str
}
