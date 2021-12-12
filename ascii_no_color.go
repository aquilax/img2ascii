package img2ascii

import (
	"bytes"
	"image"
	"image/color"
	"io"
)

// Pallette from http://paulbourke.net/dataformats/asciiart/
const DefaultAsciiPalette = "@%#*+=-:. "

type AsciiNoColor struct {
	palette     string
	paletteSize int
}

// NewAsciiNoColor creates new AsciiNoColor converter
func NewAsciiNoColor(palette string) *AsciiNoColor {
	return &AsciiNoColor{
		palette,
		len(palette),
	}
}

func (anc AsciiNoColor) GetFontRatio() float64 {
	return doubleHeightRatio
}

// Encode converts image to AsciiNoColor and writes it to out
func (anc AsciiNoColor) Encode(out io.Writer, img image.Image) error {
	bounds := img.Bounds()
	height := bounds.Dy()
	width := bounds.Dx()

	var buffer bytes.Buffer
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			buffer.WriteString(anc.GetCharForColor(img.At(x, y)))
		}
		buffer.WriteString("\n")
		io.WriteString(out, buffer.String())
		buffer.Reset()
	}
	return nil
}

// GetCharForColor returns AsciiNoColor string representation of color
func (anc AsciiNoColor) GetCharForColor(c color.Color) string {
	gray := color.GrayModel.Convert(c)
	y := int(gray.(color.Gray).Y)
	pos := int(y * (anc.paletteSize - 1) / 255)
	return string(anc.palette[pos])
}
