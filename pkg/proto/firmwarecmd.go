package proto

var (
	RspFirmwareError    = FirmwareCmd{0x00, "rspFirmwareError", CmdLen1}
	CmdGetNameVersion   = FirmwareCmd{0x01, "cmdGetNameVersion", CmdLen1}
	RspGetNameVersion   = FirmwareCmd{0x02, "rspGetNameVersion", CmdLen32}
	CmdLoadApp          = FirmwareCmd{0x03, "cmdLoadApp", CmdLen128}
	RspLoadApp          = FirmwareCmd{0x04, "rspLoadApp", CmdLen4}
	CmdLoadAppData      = FirmwareCmd{0x05, "cmdLoadAppData", CmdLen128}
	RspLoadAppData      = FirmwareCmd{0x06, "rspLoadAppData", CmdLen4}
	RspLoadAppDataReady = FirmwareCmd{0x07, "rspLoadAppDataReady", CmdLen128}
	CmdGetUDI           = FirmwareCmd{0x08, "cmdGetUDI", CmdLen1}
	RspGetUDI           = FirmwareCmd{0x09, "rspGetUDI", CmdLen32}
)

// FirmwareCmd represents a command in the Tillitis firmware protocol.
type FirmwareCmd struct {
	code   byte
	name   string
	cmdLen CmdLen
}

// Code returns the command code of the FirmwareCmd.
func (c FirmwareCmd) Code() byte {
	return c.code
}

// CmdLen returns the length of the command.
func (c FirmwareCmd) CmdLen() CmdLen {
	return c.cmdLen
}

// Endpoint returns the endpoint for the FirmwareCmd, which is always DestFW.
func (c FirmwareCmd) Endpoint() Endpoint {
	return DestFW
}

// String returns the string representation of the FirmwareCmd.
func (c FirmwareCmd) String() string {
	return c.name
}
