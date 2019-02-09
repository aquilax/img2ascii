package main

import (
	"flag"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/aquilax/img2ascii/lib"
	"github.com/nfnt/resize"
)

var maxWidth = flag.Int("width", 120, "Maximum width")
var convertor = flag.String("converter", "ansi256", "Converter")

func main() {
	flag.Parse()
	fileName := flag.Arg(0)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	width := uint(*maxWidth)
	if bounds.Max.X < *maxWidth {
		width = uint(bounds.Max.X)
	}

	var c img2ascii.Converter
	if *convertor == "ascii" {
		c = &img2ascii.AsciiNoColor{}
	} else if *convertor == "24bit" {
		c = &img2ascii.TrueColors{}
	} else if *convertor == "24bit2x" {
		c = &img2ascii.TrueColorsDoubleHeight{}
	} else if *convertor == "ansi2562x" {
		c = &img2ascii.ANSI256ColorsDoubleHeight{}
	} else {
		c = &img2ascii.ANSI256Colors{}
	}
	height := uint(float64(bounds.Max.Y) * c.GetFontRatio() * float64(width) / float64(bounds.Max.X))

	newImage := resize.Resize(width, height, img, resize.Lanczos3)

	if err := c.Process(newImage, os.Stdout); err != nil {
		panic(err)
	}
}
