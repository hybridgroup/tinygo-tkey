// The client for the TKey Blinker example.
// It allows you to control the LED blinking on a TKey device.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/pflag"
	"github.com/tillitis/tkeyclient"
)

func main() {
	var devPath string
	var led, timing, speed int
	var blinking, verbose, helpOnly bool
	pflag.CommandLine.SortFlags = false
	pflag.StringVar(&devPath, "port", "",
		"Set serial port device `PATH`. If this is not passed, auto-detection will be attempted.")
	pflag.IntVar(&speed, "speed", tkeyclient.SerialSpeed,
		"Set serial port speed in `BPS` (bits per second).")
	pflag.BoolVar(&verbose, "verbose", false,
		"Enable verbose output.")
	pflag.IntVar(&led, "led", 0,
		"Set LED")
	pflag.IntVar(&timing, "timing", 500,
		"Set blink timing (in ms).")
	pflag.BoolVar(&blinking, "blinking", true,
		"Set blinking on or off.")
	pflag.BoolVar(&helpOnly, "help", false, "Output this help.")
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n%s", os.Args[0],
			pflag.CommandLine.FlagUsagesWrapped(80))
	}
	pflag.Parse()

	if helpOnly {
		pflag.Usage()
		os.Exit(0)
	}

	if !verbose {
		tkeyclient.SilenceLogging()
	}

	if devPath == "" {
		var err error
		devPath, err = tkeyclient.DetectSerialPort(true)
		if err != nil {
			os.Exit(1)
		}
	}

	tk := tkeyclient.New()
	fmt.Printf("Connecting to device on serial port %s ...\n", devPath)
	if err := tk.Connect(devPath, tkeyclient.WithSpeed(speed)); err != nil {
		fmt.Printf("Could not open %s: %v\n", devPath, err)
		os.Exit(1)
	}
	exit := func(code int) {
		if err := tk.Close(); err != nil {
			fmt.Printf("tk.Close: %v\n", err)
		}
		os.Exit(code)
	}
	handleSignals(func() { exit(1) }, os.Interrupt, syscall.SIGTERM)

	bl := NewBlinker(tk)

	if err := bl.SetLED(led); err != nil {
		fmt.Printf("SetLED: %v\n", err)
		exit(1)
	}

	if err := bl.SetTiming(timing); err != nil {
		fmt.Printf("SetTiming: %v\n", err)
		exit(1)
	}

	if err := bl.Blinking(blinking); err != nil {
		fmt.Printf("Blinking: %v\n", err)
		exit(1)
	}

	exit(0)
}

func handleSignals(action func(), sig ...os.Signal) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, sig...)
	go func() {
		for {
			<-ch
			action()
		}
	}()
}
