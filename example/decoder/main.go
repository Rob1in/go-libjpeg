package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/viam-labs/go-libjpeg/jpeg"
)

func pngToImage(path string) image.Image {
	openBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	img, err := png.Decode(bytes.NewReader(openBytes))
	if err != nil {
		panic(err.Error())
	}
	return img
}

func main() {
	img := pngToImage("data/landscape.png")
	if img == nil {
		log.Fatalln("Got nil")
	}
	img.ColorModel()
	//bf := new(bytes.Buffer)
	f, err := os.Create(filepath.Clean("data/encoded_ouais.jpeg"))
	err = jpeg.Encode(f, img, &jpeg.EncoderOptions{Quality: 2})
	if err != nil {
		log.Fatalf("Got Error: %v", err)
	}
	fmt.Println(img.Bounds().Dx())
	fmt.Println(img.Bounds().Dy())
	//switch img.(type) {
	//case *image.YCbCr:
	//	log.Println("decoded YCbCr")
	//case *image.Gray:
	//	log.Println("decoded Gray")
	//default:
	//	log.Println("unknown format")
	//}
}
