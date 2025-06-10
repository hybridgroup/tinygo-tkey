package main

import (
	"errors"
	"fmt"

	"github.com/tillitis/tkeyclient"
)

var errInvalidCommand = errors.New("invalid command")

// commands for the blinker app running on the TKey.
var (
	cmdSetLED    = appCmd{0x01, "cmdSetLED", tkeyclient.CmdLen32}
	rspSetLED    = appCmd{0x02, "rspSetLED", tkeyclient.CmdLen4}
	cmdSetTiming = appCmd{0x03, "cmdSetTiming", tkeyclient.CmdLen32}
	rspSetTiming = appCmd{0x04, "rspSetTiming", tkeyclient.CmdLen4}
	cmdBlinking  = appCmd{0x05, "cmdBlinking", tkeyclient.CmdLen4}
	rspBlinking  = appCmd{0x06, "rspBlinking", tkeyclient.CmdLen4}
)

// appCmd represents a command in the Tillitis application protocol.
type appCmd struct {
	code   byte
	name   string
	cmdLen tkeyclient.CmdLen
}

func (c appCmd) Code() byte {
	return c.code
}

func (c appCmd) CmdLen() tkeyclient.CmdLen {
	return c.cmdLen
}

func (c appCmd) Endpoint() tkeyclient.Endpoint {
	return tkeyclient.DestApp
}

func (c appCmd) String() string {
	return c.name
}

type Blinker struct {
	tk *tkeyclient.TillitisKey // A connection to a TKey
}

// New allocates a struct for communicating with the timer app running
// on the TKey. You're expected to pass an existing connection to it,
// so use it like this:
//
//	tk := tkeyclient.New()
//	err := tk.Connect(port)
//	blinker := NewBlinker(tk)
func NewBlinker(tk *tkeyclient.TillitisKey) Blinker {
	var blinker Blinker

	blinker.tk = tk

	return blinker
}

// SetLED sets the LED on the TKey to the specified value.
func (b Blinker) SetLED(led int) error {
	return b.sendIntCommand(cmdSetLED, rspSetLED, led)
}

// SetTiming sets the timing for the blinking in milliseconds.
func (b Blinker) SetTiming(ms int) error {
	return b.sendIntCommand(cmdSetTiming, rspSetTiming, ms)
}

// Blinking sets the blinking state of the LED on the TKey.
func (b Blinker) Blinking(on bool) error {
	return b.sendBoolCommand(cmdBlinking, rspBlinking, on)
}

// sendBoolCommand sends a boolean command to the TKey and expects a response.
func (b Blinker) sendBoolCommand(sendCmd appCmd, expectedReceiveCmd appCmd, on bool) error {
	id := 2
	tx, err := tkeyclient.NewFrameBuf(sendCmd, id)
	if err != nil {
		return fmt.Errorf("error on NewFrameBuf: %w", err)
	}

	// The boolean
	if on {
		tx[2] = 1
	} else {
		tx[2] = 0
	}
	tkeyclient.Dump("tx", tx)
	if err = b.tk.Write(tx); err != nil {
		return fmt.Errorf("error on write: %w", err)
	}

	rx, _, err := b.tk.ReadFrame(expectedReceiveCmd, id)
	tkeyclient.Dump("rx", rx)
	if err != nil {
		return fmt.Errorf("ReadFrame: %w", err)
	}

	if rx[2] != tkeyclient.StatusOK {
		return errInvalidCommand
	}

	return nil
}

// sendIntCommand sends an integer command to the TKey and expects a response.
func (b Blinker) sendIntCommand(sendCmd appCmd, expectedReceiveCmd appCmd, i int) error {
	id := 2
	tx, err := tkeyclient.NewFrameBuf(sendCmd, id)
	if err != nil {
		return fmt.Errorf("NewFrameBuf: %w", err)
	}

	// The integer
	tx[2] = byte(i)
	tx[3] = byte(i >> 8)
	tx[4] = byte(i >> 16)
	tx[5] = byte(i >> 24)
	tkeyclient.Dump("tx", tx)
	if err = b.tk.Write(tx); err != nil {
		return fmt.Errorf("error on write: %w", err)
	}

	rx, _, err := b.tk.ReadFrame(expectedReceiveCmd, id)
	tkeyclient.Dump("rx", rx)
	if err != nil {
		return fmt.Errorf("ReadFrame: %w", err)
	}

	if rx[2] != tkeyclient.StatusOK {
		return errInvalidCommand
	}

	return nil
}
