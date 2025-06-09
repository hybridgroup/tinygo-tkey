// Package proto provides the protocol definitions for the [Tillitis TKey-1](https://github.com/tillitis/tillitis-key1)
// an open source, open hardware FPGA-based USB security token.
//
// This package implements the [Tillitis framing protocol for communication](https://dev.tillitis.se/protocol/)
// for use by device applications written using TinyGo.
//
// It defines the commands, endpoints, and framing headers used in the protocol.
// It also provides types and functions for working with commands and framing headers.
//
// How To Use:
//
// - Import the package in your TinyGo application.
//
// - Use the framing headers to parse and construct frames for communication.
//
// Example of defining application commands:
//
//	cmdSetLED    = proto.NewAppCmd(0x01, "cmdSetLED", proto.CmdLen32)
//	rspSetLED    = proto.NewAppCmd(0x02, "rspSetLED", proto.CmdLen4)
//
// Once you have defined your commands, you can use them to create frames for sending over UART or other communication channels.
// For example, to create a response frame for setting an LED:
//
//		response, err = proto.NewFrame(rspSetLED, 2, []byte{proto.StatusOK}
//		if err != nil {
//			// handle error
//		}
//		// now send response via UART or other communication channel
//	    ....
//
// You can also create an error frame like this:
//
//	response, err = proto.NewFrame(proto.NewAppCmd(0x00, "cmdUnknown", proto.CmdLen1), 2, []byte{proto.StatusBad})
//
// Example of parsing a framing header:
//
//		hdr, err := proto.ParseFramingHdr(rx[0])
//		if err != nil {
//			// handle error
//		}
//	    // did we get a valid header?
//		if hdr.Len() > 0 {
//			handleCommand(rx, tx)
//			...
//
// A typical command handler might look like this:
//
//	func handleCommand(rx []byte, tx []byte) (err error) {
//		var response proto.Frame
//
//		switch rx[1] {
//		case cmdSetLED.Code():
//			changeLED(rx[2])
//
//			response, err = proto.NewFrame(rspSetLED, 2, []byte{proto.StatusOK})
//			...
package proto
