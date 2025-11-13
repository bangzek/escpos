package escpos

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MODEL_MASK  = 0b1001_0000
	MODEL_FIXED = 0b0000_0000

	TYPE_MASK       = 0b1001_0000
	TYPE_FIXED      = 0b0000_0000
	TYPE_MULTI_BYTE = 0b0000_0001
	TYPE_AUTOCUTTER = 0b0000_0010
	TYPE_DISPLAY    = 0b0000_0100
)

//----------------------------------------------------------------------

type ModelID byte

func (id ModelID) IsValid() bool {
	return id&MODEL_MASK == MODEL_FIXED
}

func (id ModelID) String() string {
	return strconv.Itoa(int(id))
}

//----------------------------------------------------------------------

type TypeID byte

func (id TypeID) IsValid() bool {
	return id&TYPE_MASK == TYPE_FIXED
}

func (id TypeID) HasMultiByte() bool {
	return id&TYPE_MULTI_BYTE == TYPE_MULTI_BYTE
}

func (id TypeID) HasAutocutter() bool {
	return id&TYPE_AUTOCUTTER == TYPE_AUTOCUTTER
}

func (id TypeID) HasDisplay() bool {
	return id&TYPE_DISPLAY == TYPE_DISPLAY
}

func (id TypeID) String() (str string) {
	if id.HasMultiByte() {
		str = "MULTI BYTE"
	}
	if id.HasAutocutter() {
		if str != "" {
			str += ", "
		}
		str += "AUTOCUTTER"
	}
	if id.HasDisplay() {
		if str != "" {
			str += ", "
		}
		str += "DISPLAY"
	}
	if str == "" {
		str = "NONE"
	}
	return str
}

func (id TypeID) MarshalText() ([]byte, error) {
	if id.IsValid() {
		return []byte(id.String()), nil
	} else {
		return nil, fmt.Errorf("Invalid TypeID: %d", id)
	}
}

func (id *TypeID) UnmarshalText(b []byte) error {
	if len(b) == 0 {
		return fmt.Errorf("Invalid TypeID from %q", b)
	}
	a := strings.Split(string(b), ", ")
	for _, s := range a {
		switch s {
		case "MULTI BYTE", "AUTOCUTTER", "DISPLAY", "NONE": // do nothing
		default:
			return fmt.Errorf("Invalid TypeID from %q", b)
		}
	}
	*id = 0
	for _, s := range a {
		switch s {
		case "MULTI BYTE":
			*id |= TYPE_MULTI_BYTE
		case "AUTOCUTTER":
			*id |= TYPE_AUTOCUTTER
		case "DISPLAY":
			*id |= TYPE_DISPLAY
		}
	}
	return nil
}
