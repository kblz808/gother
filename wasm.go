package main

import (
	"bytes"
	"encoding/base64"
	"image"
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

	// process(img)

	// println(img)

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

func process(img image.Image) {
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}

	var buffer bytes.Buffer
	err := jpeg.Encode(&buffer, grayImg, nil)
	if err != nil {
		log.Println("error encoding the image", err)
		return
	}
	encodedImage := base64.StdEncoding.EncodeToString(buffer.Bytes())
	log.Println("encoded image:", encodedImage)
}

// func decodeImage(this js.Value, args []js.Value) interface{} {
// 	jsBytes := args[0]
// 	length := jsBytes.Get("byteLength").Int()
// 	data := make([]byte, length)
// 	js.CopyBytesToGo(data, jsBytes)

// 	img, err := jpeg.Decode(bytes.NewReader(data))
// 	if err != nil {
// 		return nil
// 	}

// 	println(img)

// 	return nil
// }
