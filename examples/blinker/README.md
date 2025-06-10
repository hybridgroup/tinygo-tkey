# Blinker

![tkey led](../../images/tkey-led.gif)

Example application for TKey written using TinyGo for the device application and Go for the client application.

It lets you control the built-in LED on the TKey, by sending commands via USB using the CLI application on your computer.

## Device application

The files in the "app" directory contain the device application that runs on the TKey device.

To compile and flash the TKey with the device application:

```shell
tinygo flash -size short -target=tkey ./examples/blinker/app
```

The LED should start blinking green every half second.

## Client application

The files in the "cmd" directory contain the client application that runs on your computer and communicates with the TKey device.


Now you can run the command line client application on your computer:

```shell
go run ./examples/blinker/cmd --led 0 --timing 250
```

The LED should now be blinking blue every 250 ms.
