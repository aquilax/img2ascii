package img2ascii

import (
	"bytes"
	"image"
	"image/color"
	"io"
	"strconv"
)

type ANSI256ColorsDoubleHeight struct{}

func NewANSI256ColorsDoubleHeight() *ANSI256ColorsDoubleHeight {
	return &ANSI256ColorsDoubleHeight{}
}

func (c ANSI256ColorsDoubleHeight) GetFontRatio() float64 {
	return halfHeightRatio
}

func (c *ANSI256ColorsDoubleHeight) Encode(out io.Writer, img image.Image) error {
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

func (c ANSI256ColorsDoubleHeight) GetCharForColor(c1 color.Color, c2 color.Color) string {
	colorTop := getANSIColor(c1)
	colorBottom := getANSIColor(c2)
	return "\x1b[38;5;" + strconv.Itoa(colorTop) + "m" + "\x1b[48;5;" + strconv.Itoa(colorBottom) + "m" + "▀"
}
