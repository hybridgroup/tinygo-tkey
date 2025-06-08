package main

import (
	"github.com/hybridgroup/tinygo-tkey/pkg/proto"
)

func handleLoadingCommand(rx []byte, tx []byte) (err error) {
	var response proto.Frame
	var invalidCommand bool

	hdr, err := proto.ParseFramingHdr(rx[0])
	if err != nil {
		return err
	}

	switch rx[1] {
	case cmdLoadData.Code():
		message = append(message, rx[2:cmdLoadData.CmdLen().Bytelen()+1]...)
		response, err = proto.NewFrame(rspLoadData, int(hdr.ID), []byte{proto.StatusOK})
		if len(message) >= messageSize {
			currentState = stateSigning
		}

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
