.PHONY: build wasm server

all: format wasm server

wasm:
	GOOS=js GOARCH=wasm go build -o assets/main.wasm wasm/main.go

server:
	go build -o main server/main.go

format:
	gofmt -s -w .

clean:
	-rm assets/main.wasm
	-rm main
