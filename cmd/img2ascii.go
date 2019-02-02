package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/aquilax/img2ascii"
)

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	c := img2ascii.AsciiNoColor{}
	img2ascii.Process(img, os.Stdout, c)

}
