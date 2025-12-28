
wasm:
	GOOS=js GOARCH=wasm go build -o internal/app/server/http/resources/app.wasm ./cmd/wasm/

server: wasm
	go build -o build/server ./cmd/server/

all: server