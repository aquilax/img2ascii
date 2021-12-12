package img2ascii

import (
	"bytes"
	"image"
	"image/color"
	"io"
)

type TrueColorsDoubleHeight struct {
}

func (c TrueColorsDoubleHeight) GetFontRatio() float64 {
	return halfHeightRatio
}

func (c *TrueColorsDoubleHeight) Process(img image.Image, out io.WriteCloser) error {
	bounds := img.Bounds()
	height := bounds.Dy()
	width := bounds.Dx()

	var buffer bytes.Buffer
	for y := 0; y < height/2; y++ {
		for x := 0; x < width; x++ {
			buffer.WriteString(c.GetCharForColor(img.At(x, y*2), img.At(x, y*2+1)))
		}
		buffer.WriteString("\x1b[0m\n")
		io.WriteString(out, buffer.String())
		buffer.Reset()
	}
	return nil
}

func (c TrueColorsDoubleHeight) GetCharForColor(c1 color.Color, c2 color.Color) string {
	colorTop := getTrueColor(c1)
	colorBottom := getTrueColor(c2)
	return "\x1b[38;2;" + colorTop + "m" + "\x1b[48;2;" + colorBottom + "m" + "▀"
}
