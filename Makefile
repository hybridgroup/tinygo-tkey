
build:
	mkdir -p build

.PHONY: build-blinker-app
build-blinker-app: build
	tinygo build -o ./build/blinker-app.hex -size short -target=tkey ./examples/blinker/app

.PHONY: flash-blinker-app
flash-blinker-app:
	tinygo flash -size short -target=tkey ./examples/blinker/app

.PHONY: blinker-cmd
blinker-cmd:
	go run ./examples/blinker/cmd

.PHONY: build-signer-app
build-signer-app:
	tinygo build -o ./build/signer-app.hex -size short -target=tkey ./examples/signer/app

.PHONY: flash-signer-app
flash-signer-app:
	tinygo flash -size short -target=tkey ./examples/signer/app

.PHONY: test
test:
	go test ./...
