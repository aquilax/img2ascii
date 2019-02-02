package main

import (
	"flag"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/aquilax/img2ascii"
	"github.com/nfnt/resize"
)

// https://unix.stackexchange.com/questions/148569/standard-terminal-font-aspect-ratio
const fontRatio = 9 / 22.

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
	height := uint(float64(bounds.Max.Y) * fontRatio * float64(width) / float64(bounds.Max.X))

	newImage := resize.Resize(width, height, img, resize.Lanczos3)

	var c img2ascii.Converter
	if *convertor == "ascii" {
		c = img2ascii.AsciiNoColor{}
	} else {
		c = img2ascii.ANSI256Colors{}
	}
	img2ascii.Process(newImage, os.Stdout, c)
}
