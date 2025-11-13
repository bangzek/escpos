package main

import (
	"log"
	"os"

	"github.com/bangzek/escpos"
)

const (
	//SER_DEV = "/dev/ttyS1"
	//SER_DEV = "/dev/ttyUSB0"
	//SER_DEV  = "COM1"
	SER_DEV  = "/dev/ttyVPORT4"
	FILE_DEV = "/dev/usb/lp0"
)

func main() {
	log.SetOutput(os.Stdout)
	escpos.ErrorLogFunc = log.Printf
	escpos.InfoLogFunc = log.Printf
	escpos.DebugLogFunc = log.Printf

	dev := &escpos.SerialDev{
		Device: SER_DEV,
	}

	/*
		dev := &escpos.FileDev{
			Device: FILE_DEV,
		}
	*/

	con := &escpos.Controller{
		Dev: dev,
	}

	sc := &escpos.Scanner{Controller: con}
	ps, err := con.PrinterStatus()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ps)

	os, err := con.OfflineStatus()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(os)

	es, err := con.ErrorStatus()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(es)

	rs, err := con.RollStatus()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rs)

	cmd := make(chan []escpos.Cmd, 1)
	evt := sc.Scan(cmd)

	for list := range evt {
		for _, ev := range list {
			log.Printf("EVENT %s", ev)
			switch e := ev.(type) {
			case escpos.ConnectionEvent:
				if e.Connected {
					cmd <- []escpos.Cmd{
						escpos.PulseCmd{5, 250, 250},
						escpos.TypeIDCmd{},
						escpos.MakerCmd{},
						escpos.ModelCmd{},
						escpos.FirmwareCmd{},
						escpos.SerialNoCmd{},
						escpos.ModelIDCmd{},
						escpos.PulseCmd{2, 250, 250},
						escpos.PrintCmd("everything is OK\n\n\n"),
					}
				}
			case escpos.StartPrintEvent:
			case escpos.FinishPrintEvent:
			case escpos.CmdEvent:
				switch c := e.Cmd.(type) {
				case escpos.PrintCmd:
					log.Printf("DEBUG %s", c)
				}
			case escpos.StateEvent:
			case escpos.DrawerEvent:
				if !e.On {
					cmd <- []escpos.Cmd{
						escpos.PulseCmd{2, 25, 25},
						escpos.PulseCmd{2, 25, 25},
						escpos.PulseCmd{2, 25, 25},
						escpos.PulseCmd{2, 25, 25},
						escpos.PulseCmd{2, 25, 25},
					}
				}
			}
		}
	}
}
