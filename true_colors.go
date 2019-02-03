package img2ascii

import (
	"fmt"
	"image"
	"image/color"
)

// https://github.com/hit9/img2txt/blob/gh-pages/ansi.py#L4
func getTrueColor(c color.Color) string {
	r, g, b, _ := c.RGBA()

	cR := int(r / 256)
	cG := int(g / 256)
	cB := int(b / 256)

	return fmt.Sprintf("%d;%d;%d", cR, cG, cB)
}

type TrueColors struct {
	lastColor string
}

func (ac TrueColors) GetImage(img image.Image) (image.Image, error) {
	return img, nil
}

func (ac *TrueColors) GetCharForColor(c color.Color) string {
	color := getTrueColor(c)
	if ac.lastColor == color {
		// Keep the existing color
		return solidChar
	}
	ac.lastColor = color
	return "\x1b[48;2;" + ac.lastColor + "m" + solidChar
}

func (ac *TrueColors) GetEndLine() string {
	ac.lastColor = ""
	return "\x1b[0m\n"
}

func (ac TrueColors) GetEndOfImage() string {
	return ""
}
