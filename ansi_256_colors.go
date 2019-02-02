package img2ascii

import (
	"image"
	"image/color"
	"math"
	"strconv"
)

// TODO: Different shading can allow for 4x more colors
const solidChar = " "

// https://github.com/hit9/img2txt/blob/gh-pages/ansi.py#L4
func getANSIColor(c color.Color) int {
	r, g, b, _ := c.RGBA()

	websafeR := int(math.Round((float64(r) / 65535.0) * 5))
	websafeG := int(math.Round((float64(g) / 65535.0) * 5))
	websafeB := int(math.Round((float64(b) / 65535.0) * 5))

	return int(((websafeR * 36) + (websafeG * 6) + websafeB) + 16)
}

type ANSI256Colors struct {
	lastColor int
}

func (ac ANSI256Colors) GetImage(img image.Image) (image.Image, error) {
	return img, nil
}

func (ac ANSI256Colors) GetCharForColor(c color.Color) string {
	color := getANSIColor(c)
	if ac.lastColor == color {
		// Keep the existing color
		return solidChar
	}
	ac.lastColor = color
	return "\x1b[48;5;" + strconv.Itoa(ac.lastColor) + "m" + solidChar
}

func (ac ANSI256Colors) GetEndLine() string {
	return "\x1b[0m\n"
}

func (ac ANSI256Colors) GetEndOfImage() string {
	return "\x1b[0m"
}
