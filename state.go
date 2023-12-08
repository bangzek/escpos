package escpos

import "fmt"

type State byte

const (
	OnlineState State = iota
	OfflineState
	WaitRecoveryState
	FedByButtonState
	RecoverableErrState
	AutoRecoverErrState
	UnrecoverableErrState
	AutocutterErrState
	NoPaperState
	CoverOpenState
)

func (s State) IsValid() bool {
	switch s {
	case OnlineState, OfflineState,
		WaitRecoveryState, FedByButtonState,
		RecoverableErrState, AutoRecoverErrState,
		UnrecoverableErrState, AutocutterErrState,
		NoPaperState, CoverOpenState:
		return true
	default:
		return false
	}
}

func (s State) String() string {
	switch s {
	case OnlineState:
		return "ONLINE"
	case OfflineState:
		return "OFFLINE"
	case WaitRecoveryState:
		return "WAIT RECOVERY"
	case FedByButtonState:
		return "FED BY BUTTON"
	case RecoverableErrState:
		return "RECOVERABLE ERR"
	case AutoRecoverErrState:
		return "AUTO-RECOVER ERR"
	case UnrecoverableErrState:
		return "UNRECOVERABLE ERR"
	case AutocutterErrState:
		return "AUTOCUTTER ERR"
	case NoPaperState:
		return "NO PAPER"
	case CoverOpenState:
		return "COVER OPEN"
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
		*s = OnlineState
	case "OFFLINE":
		*s = OfflineState
	case "WAIT RECOVERY":
		*s = WaitRecoveryState
	case "FED BY BUTTON":
		*s = FedByButtonState
	case "RECOVERABLE ERR":
		*s = RecoverableErrState
	case "AUTO-RECOVER ERR":
		*s = AutoRecoverErrState
	case "UNRECOVERABLE ERR":
		*s = UnrecoverableErrState
	case "AUTOCUTTER ERR":
		*s = AutocutterErrState
	case "NO PAPER":
		*s = NoPaperState
	case "COVER OPEN":
		*s = CoverOpenState
	default:
		return fmt.Errorf("Invalid State from %q", b)
	}
	return nil
}
