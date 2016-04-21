package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func usage() {
	fmt.Println(`
usage: fpanel_control {A|B|C|D} [delay seconds,default 1(s)]
`)
	os.Exit(-1)
}
func main() {

	if len(os.Args) != 3 {
		usage()
	}
	var pin PinName
	switch os.Args[1] {
	case "A":
		pin = PinCTS
	case "B":
		pin = PinTXD
	case "C":
		pin = PinRXD
	case "D":
		pin = PinDTR
	default:
		usage()
	}

	delay, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("E:", err)
		usage()
	}

	ctx, err := NewFtdiContext()
	if err != nil {
		fmt.Println("E:", err)
		return
	}
	defer ctx.Close()

	ctx.Enable(pin, true)
	<-time.After(time.Second * time.Duration(delay))
	ctx.Enable(pin, false)
}
