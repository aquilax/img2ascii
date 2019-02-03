package img2ascii

import (
	"bytes"
	"image"
	"image/color"
	"io"
)

// Pallette from http://paulbourke.net/dataformats/asciiart/
const palette = "@%#*+=-:. "

var palletteCount = uint(len(palette))

type AsciiNoColor struct{}

func (c AsciiNoColor) GetFontRatio() float64 {
	return doubleHeightRatio
}

func (c AsciiNoColor) Process(img image.Image, out io.WriteCloser) error {
	bounds := img.Bounds()
	height := bounds.Dy()
	width := bounds.Dx()

	var buffer bytes.Buffer
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			buffer.WriteString(c.GetCharForColor(img.At(x, y)))
		}
		buffer.WriteString("\n")
		io.WriteString(out, buffer.String())
		buffer.Reset()
	}
	return nil
}

func (anc AsciiNoColor) GetCharForColor(c color.Color) string {
	gray := color.GrayModel.Convert(c)
	y := uint(gray.(color.Gray).Y)
	pos := int(y * (palletteCount - 1) / 255)
	return string(palette[pos])
}
