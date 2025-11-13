package escpos

type Res interface {
	Error() error
	Res() any
	String() string
}

//----------------------------------------------------------------------

type CmdRes struct {
	Err error
}

func (r CmdRes) Error() error {
	return r.Err
}

func (r CmdRes) Res() any {
	return nil
}

func (r CmdRes) String() string {
	if r.Err == nil {
		return "<nil>"
	} else {
		return r.Err.Error()
	}
}

//----------------------------------------------------------------------

type ModelIDRes struct {
	ID  *ModelID
	Err error
}

func (r ModelIDRes) Error() error {
	return r.Err
}

func (r ModelIDRes) Res() any {
	return r.ID
}

func (r ModelIDRes) String() string {
	if r.ID == nil {
		if r.Err == nil {
			return "<nil>, <nil>"
		} else {
			return "<nil>, " + r.Err.Error()
		}
	} else {
		if r.Err == nil {
			return r.ID.String() + ", <nil>"
		} else {
			return r.ID.String() + ", " + r.Err.Error()
		}
	}

}

//----------------------------------------------------------------------

type TypeIDRes struct {
	ID  *TypeID
	Err error
}

func (r TypeIDRes) Error() error {
	return r.Err
}

func (r TypeIDRes) Res() any {
	return r.ID
}

func (r TypeIDRes) String() string {
	if r.ID == nil {
		if r.Err == nil {
			return "<nil>, <nil>"
		} else {
			return "<nil>, " + r.Err.Error()
		}
	} else {
		if r.Err == nil {
			return r.ID.String() + ", <nil>"
		} else {
			return r.ID.String() + ", " + r.Err.Error()
		}
	}

}

//----------------------------------------------------------------------

type InfoRes struct {
	Info string
	Err  error
}

func (r InfoRes) Error() error {
	return r.Err
}

func (r InfoRes) Res() any {
	return r.Info
}

func (r InfoRes) String() string {
	if r.Err == nil {
		return r.Info + ", <nil>"
	} else {
		return r.Info + ", " + r.Err.Error()
	}

}
