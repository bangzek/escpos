package escpos

import (
	"fmt"
	"time"
)

const (
	event_TAG = 31
)

type Event interface {
	eventTag() int
	String() string
}

//----------------------------------------------------------------------

type ConnectionEvent struct {
	Time      time.Time
	Connected bool
}

func (e ConnectionEvent) eventTag() int {
	return event_TAG
}

func (e ConnectionEvent) String() string {
	return fmt.Sprintf("%s CONN %t", e.Time.Format(time.TimeOnly), e.Connected)
}

//----------------------------------------------------------------------

type StartPrintEvent struct {
	Time  time.Time
	Print []byte
}

func (e StartPrintEvent) eventTag() int {
	return event_TAG
}

func (e StartPrintEvent) String() string {
	return fmt.Sprintf("%s SPRN %d", e.Time.Format(time.TimeOnly), len(e.Print))
}

//----------------------------------------------------------------------

type FinishPrintEvent struct {
	Time  time.Time
	Print []byte
}

func (e FinishPrintEvent) eventTag() int {
	return event_TAG
}

func (e FinishPrintEvent) String() string {
	return fmt.Sprintf("%s FPRN %d", e.Time.Format(time.TimeOnly), len(e.Print))
}

//----------------------------------------------------------------------

type CmdEvent struct {
	Time time.Time
	Cmd  Cmd
	Res  Res
}

func (e CmdEvent) eventTag() int {
	return event_TAG
}

func (e CmdEvent) String() string {
	return fmt.Sprintf("%s  CMD %s -> %s",
		e.Time.Format(time.TimeOnly), e.Cmd, e.Res)
}

//----------------------------------------------------------------------

type StateEvent struct {
	Time  time.Time
	State State
}

func (e StateEvent) eventTag() int {
	return event_TAG
}

func (e StateEvent) String() string {
	return fmt.Sprintf("%s STAT %s", e.Time.Format(time.TimeOnly), e.State)
}

//----------------------------------------------------------------------

type DrawerEvent struct {
	Time time.Time
	On   bool
}

func (e DrawerEvent) eventTag() int {
	return event_TAG
}

func (e DrawerEvent) String() string {
	return fmt.Sprintf("%s DRWR %t", e.Time.Format(time.TimeOnly), e.On)
}

//----------------------------------------------------------------------

type FeedButtonEvent struct {
	Time time.Time
	On   bool
}

func (e FeedButtonEvent) eventTag() int {
	return event_TAG
}

func (e FeedButtonEvent) String() string {
	return fmt.Sprintf("%s FBTN %t", e.Time.Format(time.TimeOnly), e.On)
}

//----------------------------------------------------------------------

type PaperNearEndEvent struct {
	Time time.Time
	On   bool
}

func (e PaperNearEndEvent) eventTag() int {
	return event_TAG
}

func (e PaperNearEndEvent) String() string {
	return fmt.Sprintf("%s NEND %t", e.Time.Format(time.TimeOnly), e.On)
}
