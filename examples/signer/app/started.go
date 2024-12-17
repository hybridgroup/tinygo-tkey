package main

import (
	"encoding/binary"

	"github.com/hybridgroup/tinygo-tkey/pkg/proto"
)

func handleStartedCommand(rx []byte, tx []byte) (err error) {
	var response proto.Frame
	hdr, err := proto.ParseFramingHdr(rx[0])
	if err != nil {
		return err
	}

	if hdr.Endpoint == proto.DestFW {
		response, err = proto.FirmwareErrorFrame(int(hdr.ID))
	} else {
		switch rx[1] {
		case cmdGetPublicKey.Code():
			response, err = proto.NewFrame(rspGetPublicKey, int(hdr.ID), publicKey)

		case cmdSetSize.Code():
			messageSize = int(binary.LittleEndian.Uint32(rx[2:]))
			response, err = proto.NewFrame(rspSetSize, int(hdr.ID), []byte{proto.StatusOK})
			currentState = stateLoading

		case cmdGetNameVersion.Code():
			result := make([]byte, 32)
			copy(result[0:], []byte(app_name0))
			copy(result[4:], []byte(app_name1))
			binary.LittleEndian.PutUint32(result[8:], app_version)

			response, err = proto.NewFrame(rspGetNameVersion, int(hdr.ID), result)

		case cmdGetFirmwareHash.Code():
			// TODO; implement
			response, err = proto.NewFrame(rspGetFirmwareHash, int(hdr.ID), []byte{proto.StatusOK})

		default:
			response, err = proto.AppErrorFrame(int(hdr.ID))

		}
	}

	if err != nil {
		return err
	}

	// read response into tx buffer
	response.Read(tx)

	// write tx buffer with response
	uart.Write(tx[:response.Len()+1])

	return nil
}
