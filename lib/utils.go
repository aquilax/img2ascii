package img2ascii

import (
	"fmt"
	"image/color"
	"math"
)

const emptyChar = " "
const doubleHeightRatio = 9 / 22.
const halfHeightRatio = 9 / 11.

func getTrueColor(c color.Color) string {
	r, g, b, _ := c.RGBA()

	cR := int(r / 256)
	cG := int(g / 256)
	cB := int(b / 256)

	return fmt.Sprintf("%d;%d;%d", cR, cG, cB)
}

// https://github.com/hit9/img2txt/blob/gh-pages/ansi.py#L4
func getANSIColor(c color.Color) int {
	r, g, b, _ := c.RGBA()

	websafeR := int(math.Round((float64(r) / 65535.0) * 5))
	websafeG := int(math.Round((float64(g) / 65535.0) * 5))
	websafeB := int(math.Round((float64(b) / 65535.0) * 5))

	return int(((websafeR * 36) + (websafeG * 6) + websafeB) + 16)
}
