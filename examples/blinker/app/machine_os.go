//go:build !tinygo

// This file contains various functions stubs to be able to run the tests using "big" Go on a full OS
// without the need for a real device.
// This file is not used when building for tinygo, as it will use the file machine_tkey.go file instead
// and the UART implementation there.
package main

var (
	uart = &UART{}
)

type UART struct {
}

func (u *UART) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (u *UART) ReadByte() (byte, error) {
	return 0, nil
}

func (u *UART) Buffered() int {
	return 0
}

func ledSet(on bool) {
}

func changeLED(p uint8) {
}
