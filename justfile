build:
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o main.wasm wasm.go
	rm dist/main.wasm
	mv main.wasm dist/