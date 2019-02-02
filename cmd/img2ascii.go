package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/aquilax/img2ascii"
	"github.com/nfnt/resize"
)

// https://unix.stackexchange.com/questions/148569/standard-terminal-font-aspect-ratio
const fontRatio = 1 / 2.5
const maxWidth = 120

func main() {
	fileName := os.Args[1]
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
	width := uint(maxWidth)
	if bounds.Max.X < maxWidth {
		width = uint(bounds.Max.X)
	}
	height := uint(float64(bounds.Max.Y) * fontRatio * float64(width) / float64(bounds.Max.X))

	newImage := resize.Resize(width, height, img, resize.Lanczos3)

	c := img2ascii.AsciiNoColor{}
	img2ascii.Process(newImage, os.Stdout, c)

}
