package go_qrcode

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/happyEgg/go_qrcode/decode"
)

const (
	VERSION = 1.0
)

func Decode(imgPath string) string {
	var body string
	var img image.Image
	var err error
	file, err := os.Open(imgPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()

	imageTypeArr := strings.Split(file.Name(), ".")
	if len(imageTypeArr) <= 1 {
		fmt.Println("Image file format error")
		os.Exit(-1)
	}

	imageType := imageTypeArr[len(imageTypeArr)-1]

	switch imageType {
	case "jpeg", "jpg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	default:
		fmt.Println("Image file format error")
		os.Exit(-1)
	}

	if err != nil {
		fmt.Println("decode failed:", err)
		os.Exit(-1)
	}

	newImg := decode.NewImage(img)
	scanner := decode.NewScanner().SetEnabledAll(true)

	symbols, _ := scanner.ScanImage(newImg)
	for _, s := range symbols {
		body += s.Data
	}

	return body
}

func Encode(value string, width, height int) image.Image {
	code, err := qr.Encode(value, qr.L, qr.Unicode)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if value != code.Content() {
		fmt.Println("data differs")
		os.Exit(-1)
	}

	codeImg, err := barcode.Scale(code, width, height)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	return codeImg
}
