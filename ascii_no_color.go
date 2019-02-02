package img2ascii

import (
	"image"
	"image/color"
)

// Pallette from http://paulbourke.net/dataformats/asciiart/
const palette = "@%#*+=-:. "

var palletteCount = uint(len(palette))

type AsciiNoColor struct{}

func (anc AsciiNoColor) GetImage(img image.Image) (image.Image, error) {
	return img, nil
}

func (anc AsciiNoColor) GetCharForColor(c color.Color) string {
	gray := color.GrayModel.Convert(c)
	y := uint(gray.(color.Gray).Y)
	pos := int(y * (palletteCount - 1) / 255)
	return string(palette[pos])
}

func (anc AsciiNoColor) GetEndLine() string {
	return "\n"
}

func (anc AsciiNoColor) GetEndOfImage() string {
	return ""
}
