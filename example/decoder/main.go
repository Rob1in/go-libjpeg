package main

import (
	"bytes"
	"fmt"
	"github.com/viam-labs/go-libjpeg/jpeg"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
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
	img := pngToImage("./data/landscape.png")
	fmt.Println("tout va bien")
	if img == nil {
		log.Fatalln("Got nil")
	}
	img.ColorModel()
	//bf := new(bytes.Buffer)
	f, err := os.Create(filepath.Clean("./data/encoded_normal2.jpeg"))
	err = jpeg.Encode(f, img, &jpeg.EncoderOptions{Quality: 2})
	if err != nil {
		log.Fatalf("Got Error: %v", err)
	}
	fmt.Println(img.Bounds().Dx())
	fmt.Println(img.Bounds().Dy())
}
