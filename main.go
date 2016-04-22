package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func usage() {
	fmt.Println(`
usage: fpanel_control {A|B|C|D} [delay seconds,default 1(s)]
`)
	os.Exit(-1)
}

var Delay = flag.Int("d", 1, "hold on time in seconds.")
var All = flag.Bool("a", false, "control all pin.")

var PinTable = map[string]PinName{
	"1": PinCTS,
	"2": PinTXD,
	"3": PinRXD,
	"4": PinDTR,
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s pin1 [pin2...]\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "Current support pins")
		fmt.Fprintln(os.Stderr, "  1  ->  CTS")
		fmt.Fprintln(os.Stderr, "  2  ->  TXD")
		fmt.Fprintln(os.Stderr, "  3  ->  RXD")
		fmt.Fprintln(os.Stderr, "  4  ->  DTR")
	}
	flag.Parse()

	var pins = make(map[PinName]struct{})

	if *All {
		for _, v := range PinTable {
			pins[v] = struct{}{}
		}
	} else {
		for _, v := range flag.Args() {
			pin, ok := PinTable[v]
			if !ok {
				fmt.Printf("W:invlid pin number :%v\n", v)
				continue
			}
			pins[pin] = struct{}{}
		}
	}

	var r []PinName
	for v := range pins {
		r = append(r, v)
	}
	err := Hold(*Delay, r...)
	if err != nil {
		fmt.Printf("Hold E:%v\n", err)
	}
}

func Hold(delay int, pins ...PinName) error {
	ctx, err := NewFtdiContext()
	if err != nil {
		return err
	}

	defer ctx.Close()

	for _, pin := range pins {
		ctx.Enable(pin, true)
	}

	<-time.After(time.Second * time.Duration(delay))

	for _, pin := range pins {
		ctx.Enable(pin, false)
	}

	return nil
}
