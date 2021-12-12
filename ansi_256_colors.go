package img2ascii

import (
	"bytes"
	"image"
	"image/color"
	"io"
	"strconv"
)

type ANSI256Colors struct {
	lastColor int
}

func NewANSI256Colors() *ANSI256Colors {
	return &ANSI256Colors{}
}

func (c ANSI256Colors) GetFontRatio() float64 {
	return doubleHeightRatio
}

func (c *ANSI256Colors) Encode(out io.Writer, img image.Image) error {
	bounds := img.Bounds()
	height := bounds.Dy()
	width := bounds.Dx()

	var buffer bytes.Buffer
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			buffer.WriteString(c.GetCharForColor(img.At(x, y)))
		}
		buffer.WriteString("\x1b[0m\n")
		c.lastColor = -1
		io.WriteString(out, buffer.String())
		buffer.Reset()
	}
	return nil
}

func (c *ANSI256Colors) GetCharForColor(col color.Color) string {
	color := getANSIColor(col)
	if c.lastColor == color {
		// Keep the existing color
		return emptyChar
	}
	c.lastColor = color
	return "\x1b[48;5;" + strconv.Itoa(c.lastColor) + "m" + emptyChar
}
