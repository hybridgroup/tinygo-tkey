package main

import (
	"crypto/ed25519"
	"errors"
	"time"

	"github.com/hybridgroup/tinygo-tkey/pkg/proto"
)

type state int

const (
	stateStarted state = iota
	stateLoading
	stateSigning
	stateFailed
)

var (
	currentState state = stateStarted
	publicKey    ed25519.PublicKey
	privateKey   ed25519.PrivateKey

	messageSize int
	message     []byte
)

func main() {
	generateKeys()

	rx := make([]byte, 256)
	tx := make([]byte, 256)
	message = make([]byte, 0, 4096)

	i := 0

	for {
		for uart.Buffered() > 0 {
			data, err := uart.ReadByte()
			if err != nil {
				i = 0
			}

			rx[i] = data
			i++

			hdr, err := proto.ParseFramingHdr(rx[0])
			if err != nil {
				// reset, and wait for next command
				i = 0
			}

			// did we receive a full command?
			len := int(hdr.Len())
			if i > len {
				if err := handleCommand(rx, tx); err != nil {
					response, _ := proto.AppErrorFrame(0)
					// read response into tx buffer
					response.Read(tx)

					// write tx buffer with response
					uart.Write(tx[:response.Len()+1])
				}

				// reset, and wait for next command
				clearBuffer(rx)
				i = 0
			}

			// wait for more data
			time.Sleep(10 * time.Millisecond)
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func handleCommand(rx []byte, tx []byte) (err error) {
	clearBuffer(tx)

	switch currentState {
	case stateStarted:
		return handleStartedCommand(rx, tx)
	case stateLoading:
		return handleLoadingCommand(rx, tx)
	case stateSigning:
		return handleSigningCommand(rx, tx)
	case stateFailed:
	}

	return errors.ErrUnsupported
}

func clearBuffer(buf []byte) {
	for i := 0; i < len(buf); i++ {
		buf[i] = 0
	}
}
