package jpeg

import (
	"bytes"
	"fmt"
	"go.viam.com/test"
	"image"
	"image/color"
	"image/draw"
	nativeJpeg "image/jpeg"
	"image/png"
	"os"
	"testing"
)

func TestCompressNRGBA(t *testing.T) {
	//TODO changer ce path de merde
	img, err := pngToImage("../example/decoder/data/landscape.png")
	fmt.Println(img.ColorModel())
	test.That(t, img.ColorModel(), test.ShouldResemble, color.NRGBAModel)
	test.That(t, err, test.ShouldBeNil)
	bf := new(bytes.Buffer)
	err = Encode(bf, img, &EncoderOptions{Quality: 2})
	test.That(t, err, test.ShouldBeNil)
}

func TestCompressYCbCr(t *testing.T) {
	//TODO changer ce path de merde
	img, err := jpegToImage("../example/decoder/data/landscape.jpeg")
	test.That(t, img.ColorModel(), test.ShouldResemble, color.YCbCrModel)
	test.That(t, err, test.ShouldBeNil)
	bf := new(bytes.Buffer)
	err = Encode(bf, img, &EncoderOptions{Quality: 2})
	test.That(t, err, test.ShouldBeNil)
}

func TestCompressGray(t *testing.T) {
	//TODO changer ce path de merde
	img, err := jpegToImage("../example/decoder/data/landscape.jpeg")
	result := image.NewGray(img.Bounds())
	draw.Draw(result, result.Bounds(), img, img.Bounds().Min, draw.Src)
	test.That(t, result.ColorModel(), test.ShouldResemble, color.GrayModel)
	test.That(t, err, test.ShouldBeNil)
	bf := new(bytes.Buffer)
	err = Encode(bf, result, &EncoderOptions{Quality: 2})
	test.That(t, err, test.ShouldBeNil)
}

func TestCompressRGBA(t *testing.T) {
	//TODO changer ce path de merde
	img, err := jpegToImage("../example/decoder/data/landscape.jpeg")
	result := image.NewRGBA(img.Bounds())
	draw.Draw(result, result.Bounds(), img, img.Bounds().Min, draw.Src)
	test.That(t, result.ColorModel(), test.ShouldResemble, color.RGBAModel)
	test.That(t, err, test.ShouldBeNil)
	bf := new(bytes.Buffer)
	err = Encode(bf, result, &EncoderOptions{Quality: 2})
	test.That(t, err, test.ShouldBeNil)
}

func pngToImage(loc string) (image.Image, error) {
	openBytes, err := os.ReadFile(loc)
	if err != nil {
		return nil, err
	}
	return png.Decode(bytes.NewReader(openBytes))
}

func jpegToImage(loc string) (image.Image, error) {
	openBytes, err := os.ReadFile(loc)
	if err != nil {
		return nil, err
	}
	return nativeJpeg.Decode(bytes.NewReader(openBytes))
}
