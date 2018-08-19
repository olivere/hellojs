GO := go1.11rc1

build:
	GOOS=js GOARCH=wasm $(GO) build -o app/main.wasm
