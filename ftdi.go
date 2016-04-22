package main

// #cgo LDFLAGS: -lftdi1
// #include "ftdi.h"
import "C"
import "fmt"
import "os"

type PinName int

const (
	PinTXD = 0x01
	PinRXD = 0x02
	PinCTS = 0x08
	PinDTR = 0x10
)

type FtdiContext struct {
	core        *C.struct_ftdi_context
	cachedValue int
}

func NewFtdiContext() (*FtdiContext, error) {
	core := C.ftdi_new()
	if core == nil {
		return nil, fmt.Errorf("ftdi_new failed")
	}
	ctx := &FtdiContext{
		core: core,
	}
	err := ctx.open()
	if err != nil {
		ctx.Close()
	}
	err = ctx.enableBitBang()

	ctx.load()

	return ctx, err
}

func (ctx *FtdiContext) Enable(pin PinName, enable bool) error {
	if enable {
		ctx.cachedValue |= int(pin)
	} else {
		ctx.cachedValue &= ^int(pin)
	}
	return ctx.update()
}

func (ctx *FtdiContext) load() error {
	var d C.uchar
	C.ftdi_read_data(ctx.core, &d, 1)
	ctx.cachedValue = int(d)
	return nil
}

func (ctx *FtdiContext) update() error {
	d := C.uchar(ctx.cachedValue)
	if C.ftdi_write_data(ctx.core, &d, 1) < 0 {
		return fmt.Errorf("Write failed for 0x%x: %s\n", ctx.cachedValue,
			C.GoString(C.ftdi_get_error_string(ctx.core)))
	}
	return nil
}

func (ctx *FtdiContext) Test(pin PinName) bool {
	return ctx.cachedValue&int(pin) == 0
}

func (ctx *FtdiContext) fatal(fmtStr string, args ...interface{}) {
	ctx.Close()
	fmt.Fprintf(os.Stderr, fmtStr, args...)
}

func (ctx *FtdiContext) enableBitBang() error {
	if C.ftdi_set_bitmode(ctx.core, 0xFF, C.BITMODE_BITBANG) < 0 {
		return fmt.Errorf("Can't enable bitbang")
	}
	return nil
}

const (
	VendorID  = 0x0403
	ProductID = 0x6001
)

func (c *FtdiContext) open() error {
	if C.ftdi_usb_open(c.core, VendorID, ProductID) < 0 {
		return fmt.Errorf("Can't open ftdi device")
	}
	return nil
}

func (c *FtdiContext) Close() {
	C.ftdi_usb_close(c.core)
	C.ftdi_free(c.core)
}
