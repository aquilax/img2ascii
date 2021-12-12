package img2ascii

import (
	"bytes"
	"image"
	"image/color"
	"io"
)

type TrueColors struct {
	lastColor string
}

func NewTrueColors() *TrueColors {
	return &TrueColors{}
}

func (c TrueColors) GetFontRatio() float64 {
	return doubleHeightRatio
}

func (c *TrueColors) Encode(out io.Writer, img image.Image) error {
	bounds := img.Bounds()
	height := bounds.Dy()
	width := bounds.Dx()

	var buffer bytes.Buffer
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			buffer.WriteString(c.GetCharForColor(img.At(x, y)))
		}
		buffer.WriteString("\x1b[0m\n")
		c.lastColor = ""
		io.WriteString(out, buffer.String())
		buffer.Reset()
	}
	return nil
}

func (c *TrueColors) GetCharForColor(col color.Color) string {
	color := getTrueColor(col)
	if c.lastColor == color {
		// Keep the existing color
		return emptyChar
	}
	c.lastColor = color
	return "\x1b[48;2;" + c.lastColor + "m" + emptyChar
}
