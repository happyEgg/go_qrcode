package go_qrcode

import (
	"fmt"
	"image/png"
	"os"
	"testing"
)

func TestQrcode(t *testing.T) {
	img := Encode("test qrcode", 300, 300)

	file, err := os.Create("./test.png")
	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		t.Fatal(err)
	}

	value := Decode("./test.png")
	fmt.Println(value)
}
