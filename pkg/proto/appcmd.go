package proto

// AppCmd represents an application command in the Tillitis framing protocol.
// It contains a command code, a name, and the length of the command.
type AppCmd struct {
	code   byte
	name   string
	cmdLen CmdLen
}

// NewAppCmd creates a new AppCmd with the specified code, name, and command length.
func NewAppCmd(code byte, name string, cmdLen CmdLen) AppCmd {
	return AppCmd{code, name, cmdLen}
}

// Code returns the command code of the AppCmd.
func (c AppCmd) Code() byte {
	return c.code
}

// CmdLen returns the length of the command.
func (c AppCmd) CmdLen() CmdLen {
	return c.cmdLen
}

// Endpoint returns the endpoint for the AppCmd, which is always DestApp.
func (c AppCmd) Endpoint() Endpoint {
	return DestApp
}

// String returns the string representation of the AppCmd.
func (c AppCmd) String() string {
	return c.name
}
