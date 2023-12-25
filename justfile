build:
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o main.wasm wasm.go
	rm dist/main.wasm
	mv main.wasm dist/

tiny:
	tinygo build -o main.wasm -target wasm ./wasm.go
	rm tiny_dist/main.wasm
	mv main.wasm tiny_dist
