package usb

// Standar classes defined by USB Spec
// https://www.usb.org/defined-class-codes

type Class uint8

const (
	ClassPrinter Class = 0x07
)

func Printers() {
	// search for usb devices

}
