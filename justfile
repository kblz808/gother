build:
	GOOS=js GOARCH=wasm go build -o main.wasm wasm.go
	rm dist/main.wasm
	mv main.wasm dist/