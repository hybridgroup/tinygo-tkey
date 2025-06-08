package main

import (
	"github.com/hybridgroup/tinygo-tkey/pkg/proto"
)

func handleSigningCommand(rx []byte, tx []byte) (err error) {
	var response proto.Frame
	var invalidCommand bool
	if len(message) < messageSize {
		return errInvalidMessage // not enough data to sign
	}

	hdr, err := proto.ParseFramingHdr(rx[0])
	if err != nil {
		return err
	}

	switch rx[1] {
	case cmdGetSig.Code():
		sig := signMessage(privateKey, message[:messageSize])
		response, err = proto.NewFrame(rspGetSig, int(hdr.ID), append([]byte{proto.StatusOK}, sig...))

		clearBuffer(message)
		message = message[:0]
		currentState = stateStarted

	default:
		response, err = proto.AppErrorFrame(int(hdr.ID))
		invalidCommand = true
	}

	if err != nil {
		return err
	}

	// read response into tx buffer
	response.Read(tx)

	// write tx buffer with response
	uart.Write(tx[:response.Len()+1])

	if invalidCommand {
		return errInvalidCommand
	}

	return nil
}
