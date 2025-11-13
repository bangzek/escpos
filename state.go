package escpos

import "fmt"

type State byte

const (
	ONLINE State = iota
	OFFLINE
	WAIT_RECOVERY
	FED_BY_BUTTON
	RECOVERABLE_ERR
	AUTO_RECOVER_ERR
	UNRECOVERABLE_ERR
	AUTOCUTTER_ERR
	NO_PAPER
	COVER_IS_OPEN
)

func (s State) IsValid() bool {
	switch s {
	case ONLINE, OFFLINE, WAIT_RECOVERY, FED_BY_BUTTON,
		RECOVERABLE_ERR, AUTO_RECOVER_ERR, UNRECOVERABLE_ERR, AUTOCUTTER_ERR,
		NO_PAPER, COVER_IS_OPEN:
		return true
	default:
		return false
	}
}

func (s State) String() string {
	switch s {
	case ONLINE:
		return "ONLINE"
	case OFFLINE:
		return "OFFLINE"
	case WAIT_RECOVERY:
		return "WAIT RECOVERY"
	case FED_BY_BUTTON:
		return "FED BY BUTTON"
	case RECOVERABLE_ERR:
		return "RECOVERABLE ERR"
	case AUTO_RECOVER_ERR:
		return "AUTO-RECOVER ERR"
	case UNRECOVERABLE_ERR:
		return "UNRECOVERABLE ERR"
	case AUTOCUTTER_ERR:
		return "AUTOCUTTER ERR"
	case NO_PAPER:
		return "NO PAPER"
	case COVER_IS_OPEN:
		return "COVER IS OPEN"
	default:
		return fmt.Sprintf("ERR(%d)", byte(s))
	}
}

func (s State) MarshalText() ([]byte, error) {
	if s.IsValid() {
		return []byte(s.String()), nil
	} else {
		return nil, fmt.Errorf("Invalid State: %d", s)
	}
}

func (s *State) UnmarshalText(b []byte) error {
	switch string(b) {
	case "ONLINE":
		*s = ONLINE
	case "OFFLINE":
		*s = OFFLINE
	case "WAIT RECOVERY":
		*s = WAIT_RECOVERY
	case "FED BY BUTTON":
		*s = FED_BY_BUTTON
	case "RECOVERABLE ERR":
		*s = RECOVERABLE_ERR
	case "AUTO-RECOVER ERR":
		*s = AUTO_RECOVER_ERR
	case "UNRECOVERABLE ERR":
		*s = UNRECOVERABLE_ERR
	case "AUTOCUTTER ERR":
		*s = AUTOCUTTER_ERR
	case "NO PAPER":
		*s = NO_PAPER
	case "COVER IS OPEN":
		*s = COVER_IS_OPEN
	default:
		return fmt.Errorf("Invalid State from %q", b)
	}
	return nil
}
