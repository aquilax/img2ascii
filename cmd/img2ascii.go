package main

import (
	"flag"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/aquilax/img2ascii"
	"github.com/aquilax/img2ascii/ansi"
	"github.com/aquilax/img2ascii/ascii"
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

	var e img2ascii.Encoder
	if *convertor == "ascii" {
		e = ascii.NewImage()
	} else if *convertor == "24bit" {
		e = ansi.NewImage(ansi.PaletteTrueColor, ansi.SingleHeight)
	} else if *convertor == "24bit2x" {
		e = ansi.NewImage(ansi.PaletteTrueColor, ansi.HalfHeight)
	} else if *convertor == "ansi2562x" {
		e = ansi.NewImage(ansi.Palette256, ansi.HalfHeight)
	} else {
		e = ansi.NewImage(ansi.Palette256, ansi.SingleHeight)
	}
	height := uint(float64(bounds.Max.Y) * e.GetFontRatio() * float64(width) / float64(bounds.Max.X))

	newImage := resize.Resize(width, height, img, resize.Lanczos3)

	if err := e.Encode(os.Stdout, newImage); err != nil {
		panic(err)
	}
}
