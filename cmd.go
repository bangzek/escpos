package escpos

import (
	"crypto/md5"
	"fmt"
	"regexp"
)

type Cmd interface {
	exec(*Controller) Res
	String() string
}

//----------------------------------------------------------------------

type PrintCmd []byte

func (c PrintCmd) exec(con *Controller) Res {
	return CmdRes{con.Send([]byte(c))}
}

func (c PrintCmd) String() string {
	return fmt.Sprintf("PRINT: %d %x", len(c), md5.Sum(c))
}

var (
	pulseLvlRe = regexp.MustCompile(string([]byte{DLE, DC4, '[', 0, 1, ']'}))
)

func (c PrintCmd) NeedDisablePulseLevel() bool {
	return pulseLvlRe.Match(c)
}

//----------------------------------------------------------------------

type RawCmd []byte

func (c RawCmd) exec(con *Controller) Res {
	return CmdRes{con.Send([]byte(c))}
}

func (c RawCmd) String() string {
	return "RAW: " + string([]byte(c))
}

//----------------------------------------------------------------------

type PressFeedButtonCmd struct{}

func (c PressFeedButtonCmd) exec(con *Controller) Res {
	return CmdRes{con.PressFeedButton()}
}

func (c PressFeedButtonCmd) String() string {
	return "PRESS FEED BTN"
}

//----------------------------------------------------------------------

type RecoverCmd struct{}

func (c RecoverCmd) exec(con *Controller) Res {
	return CmdRes{con.Recover()}
}

func (c RecoverCmd) String() string {
	return "RECOVER"
}

//----------------------------------------------------------------------

type ClearAndRecoverCmd struct{}

func (c ClearAndRecoverCmd) exec(con *Controller) Res {
	return CmdRes{con.ClearAndRecover()}
}

func (c ClearAndRecoverCmd) String() string {
	return "CLR N RECOVER"
}

//----------------------------------------------------------------------

type PulseCmd struct {
	Pin int8
	On  byte
	Off byte
}

func (c PulseCmd) Validate() error {
	if c.Pin != 2 && c.Pin != 5 {
		return fmt.Errorf("Invalid pin: %d", c.Pin)
	}
	return nil
}

func (c PulseCmd) exec(con *Controller) Res {
	if err := c.Validate(); err != nil {
		return CmdRes{err}
	}
	if c.Pin == 5 {
		return CmdRes{con.PulseDrawerPin5(c.On, c.Off)}
	} else {
		return CmdRes{con.PulseDrawerPin2(c.On, c.Off)}
	}
}

func (c PulseCmd) String() string {
	return fmt.Sprintf("PULSE #%d [%d]%d", c.Pin, c.On, c.Off)
}

//----------------------------------------------------------------------

type PulseLevelCmd struct {
	Pin   int8
	Level int8
}

func (c PulseLevelCmd) Validate() error {
	if c.Pin != 2 && c.Pin != 5 {
		return fmt.Errorf("Invalid pin: %d", c.Pin)
	}
	if c.Level < 1 || c.Level > 8 {
		return fmt.Errorf("Invalid level: %d", c.Level)
	}
	return nil
}

func (c PulseLevelCmd) exec(con *Controller) Res {
	if err := c.Validate(); err != nil {
		return CmdRes{err}
	}
	if c.Pin == 5 {
		switch c.Level {
		case 1:
			return CmdRes{con.PulseDrawerPin5Level1()}
		case 2:
			return CmdRes{con.PulseDrawerPin5Level2()}
		case 3:
			return CmdRes{con.PulseDrawerPin5Level3()}
		case 4:
			return CmdRes{con.PulseDrawerPin5Level4()}
		case 5:
			return CmdRes{con.PulseDrawerPin5Level5()}
		case 6:
			return CmdRes{con.PulseDrawerPin5Level6()}
		case 7:
			return CmdRes{con.PulseDrawerPin5Level7()}
		default:
			return CmdRes{con.PulseDrawerPin5Level8()}
		}
	} else {
		switch c.Level {
		case 1:
			return CmdRes{con.PulseDrawerPin2Level1()}
		case 2:
			return CmdRes{con.PulseDrawerPin2Level2()}
		case 3:
			return CmdRes{con.PulseDrawerPin2Level3()}
		case 4:
			return CmdRes{con.PulseDrawerPin2Level4()}
		case 5:
			return CmdRes{con.PulseDrawerPin2Level5()}
		case 6:
			return CmdRes{con.PulseDrawerPin2Level6()}
		case 7:
			return CmdRes{con.PulseDrawerPin2Level7()}
		default:
			return CmdRes{con.PulseDrawerPin2Level8()}
		}
	}
}

func (c PulseLevelCmd) String() string {
	return fmt.Sprintf("PULSE #%d L%d", c.Pin, c.Level)
}

//----------------------------------------------------------------------

type ModelIDCmd struct{}

func (c ModelIDCmd) exec(con *Controller) Res {
	id, err := con.ModelID()
	return ModelIDRes{id, err}
}

func (c ModelIDCmd) String() string {
	return "MODEL ID"
}

func (c ModelIDCmd) Result(r Res) *ModelID {
	return r.(ModelIDRes).ID
}

//----------------------------------------------------------------------

type TypeIDCmd struct{}

func (c TypeIDCmd) exec(con *Controller) Res {
	id, err := con.TypeID()
	return TypeIDRes{id, err}
}

func (c TypeIDCmd) String() string {
	return "TYPE ID"
}

func (c TypeIDCmd) Result(r Res) *TypeID {
	return r.(TypeIDRes).ID
}

//----------------------------------------------------------------------

type FirmwareCmd struct{}

func (c FirmwareCmd) exec(con *Controller) Res {
	info, err := con.Firmware()
	return InfoRes{info.String(), err}
}

func (c FirmwareCmd) String() string {
	return "FIRMWARE"
}

func (c FirmwareCmd) Result(r Res) string {
	return r.(InfoRes).Info
}

//----------------------------------------------------------------------

type MakerCmd struct{}

func (c MakerCmd) exec(con *Controller) Res {
	info, err := con.Maker()
	return InfoRes{info.String(), err}
}

func (c MakerCmd) String() string {
	return "MAKER"
}

func (c MakerCmd) Result(r Res) string {
	return r.(InfoRes).Info
}

//----------------------------------------------------------------------

type ModelCmd struct{}

func (c ModelCmd) exec(con *Controller) Res {
	info, err := con.Model()
	return InfoRes{info.String(), err}
}

func (c ModelCmd) String() string {
	return "MODEL"
}

func (c ModelCmd) Result(r Res) string {
	return r.(InfoRes).Info
}

//----------------------------------------------------------------------

type SerialNoCmd struct{}

func (c SerialNoCmd) exec(con *Controller) Res {
	info, err := con.SerialNo()
	return InfoRes{info.String(), err}
}

func (c SerialNoCmd) String() string {
	return "SERIAL NO"
}

func (c SerialNoCmd) Result(r Res) string {
	return r.(InfoRes).Info
}
