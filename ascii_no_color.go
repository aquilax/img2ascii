package img2ascii

import (
	"image"
	"image/color"
)

const palette = "MND8OZ$7I?+=~:,.."

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
