package main

import (
	"bytes"
	"encoding/base64"
	"github.com/kblz808/gother/filters"
	_ "image"
	"image/jpeg"
	"log"
	"syscall/js"
)

func main() {
	c := make(chan int)

	js.Global().Set("decodeImage", js.FuncOf(decodeImage))

	<-c
}

func decodeImage(this js.Value, args []js.Value) interface{} {
	buffer := args[0]
	length := buffer.Get("byteLength").Int()
	data := make([]byte, length)
	js.CopyBytesToGo(data, buffer)

	img, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		return nil
	}

	img = filters.Floyd(img)
	// img = filters.Ordered(img)

	var d_buffer bytes.Buffer
	err = jpeg.Encode(&d_buffer, img, nil)
	if err != nil {
		log.Println("error encoding image:", err)
		return nil
	}

	encodedImage := base64.StdEncoding.EncodeToString(d_buffer.Bytes())
	println(encodedImage)
	js.Global().Call("displayImage", encodedImage)

	return length
}
