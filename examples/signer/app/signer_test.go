package main

import (
	"testing"

	"github.com/hybridgroup/tinygo-tkey/pkg/proto"
)

func TestHandleStartedCommands(t *testing.T) {
	tests := []struct {
		name     string
		cmd      proto.AppCmd
		id       int
		data     []byte
		hasError bool
	}{
		{
			name:     "cmdGetPublicKey",
			cmd:      cmdGetPublicKey,
			id:       2,
			data:     []byte{0x00},
			hasError: false,
		},
		{
			name:     "cmdGetNameVersion",
			cmd:      cmdGetNameVersion,
			id:       2,
			data:     []byte{0x00},
			hasError: false,
		},
		{
			name:     "cmdSetSize",
			cmd:      cmdSetSize,
			id:       2,
			data:     []byte{0x00},
			hasError: false,
		},
		{
			name:     "cmdGetSig",
			cmd:      cmdGetSig,
			id:       2,
			data:     []byte{0x00},
			hasError: true,
		},
		{
			name:     "cmdLoadData",
			cmd:      cmdLoadData,
			id:       2,
			data:     []byte{0x00},
			hasError: true,
		},
	}

	for _, tt := range tests {
		rx := make([]byte, 256)
		tx := make([]byte, 256)

		t.Run(tt.name, func(t *testing.T) {
			currentState = stateStarted // Reset state for next test

			frame, _ := proto.NewFrame(tt.cmd, tt.id, tt.data)
			frame.Read(rx)

			err := handleCommand(rx, tx)
			if err != nil && !tt.hasError || (err == nil && tt.hasError) {
				t.Errorf("handleCommand(%v) = %v, want %v", tt.cmd, err, tt.hasError)
			}
		})
	}
}

func TestHandleLoadingCommands(t *testing.T) {
	generateKeys()

	tests := []struct {
		name     string
		cmd      proto.AppCmd
		id       int
		data     []byte
		hasError bool
	}{
		{
			name:     "cmdLoadData",
			cmd:      cmdLoadData,
			id:       2,
			data:     []byte{0x00},
			hasError: false,
		},
		{
			name:     "cmdGetPublicKey",
			cmd:      cmdGetPublicKey,
			id:       2,
			data:     []byte{0x00},
			hasError: true,
		},
		{
			name:     "cmdGetNameVersion",
			cmd:      cmdGetNameVersion,
			id:       2,
			data:     []byte{0x00},
			hasError: true,
		},
	}

	for _, tt := range tests {
		rx := make([]byte, 256)
		tx := make([]byte, 256)

		t.Run(tt.name, func(t *testing.T) {
			currentState = stateLoading // Reset state for next test
			frame, _ := proto.NewFrame(tt.cmd, tt.id, tt.data)
			frame.Read(rx)

			err := handleCommand(rx, tx)
			if err != nil && !tt.hasError || (err == nil && tt.hasError) {
				t.Errorf("handleCommand(%v) = %v, want %v", tt.cmd, err, tt.hasError)
			}
		})
	}
}

func TestHandleSigningCommands(t *testing.T) {
	generateKeys()

	tests := []struct {
		name     string
		cmd      proto.AppCmd
		id       int
		data     []byte
		hasError bool
	}{
		{
			name:     "cmdGetSig",
			cmd:      cmdGetSig,
			id:       2,
			data:     []byte{0x00},
			hasError: false,
		},
		{
			name:     "cmdGetPublicKey",
			cmd:      cmdGetPublicKey,
			id:       2,
			data:     []byte{0x00},
			hasError: true,
		},
		{
			name:     "cmdGetNameVersion",
			cmd:      cmdGetNameVersion,
			id:       2,
			data:     []byte{0x00},
			hasError: true,
		},
		{
			name:     "cmdLoadData",
			cmd:      cmdLoadData,
			id:       2,
			data:     []byte{0x00},
			hasError: true,
		},
	}

	for _, tt := range tests {
		rx := make([]byte, 256)
		tx := make([]byte, 256)

		t.Run(tt.name, func(t *testing.T) {
			currentState = stateSigning // Reset state for next test
			frame, _ := proto.NewFrame(tt.cmd, tt.id, tt.data)
			frame.Read(rx)

			err := handleCommand(rx, tx)
			if err != nil && !tt.hasError || (err == nil && tt.hasError) {
				t.Errorf("handleCommand(%v) = %v, want %v", tt.cmd, err, tt.hasError)
			}
		})
	}
}
