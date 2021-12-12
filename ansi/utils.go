package ansi

import (
	"fmt"
	"image/color"
	"math"
	"strconv"
)

const emptyChar = " "

// PixelHeight defines the vertical density fot the output
type PixelHeight int

const (
	// SingleHeight uses one row per vertical pixel
	SingleHeight PixelHeight = iota
	// HalfHeight uses half row per vertical pixel
	HalfHeight
)

// ColorPalette is the color palette used for the output
type ColorPalette int

type colorConverter = func(c color.Color) string

const (
	// Palette256 is 256 colors ANSI palette
	Palette256 ColorPalette = iota
	// PaletteTrueColor is True Color ANSI palette
	PaletteTrueColor
)

func getColorGetter(p ColorPalette) colorConverter {
	switch p {
	case Palette256:
		return getANSIColor
	case PaletteTrueColor:
		return getTrueColor
	default:
		return getANSIColor
	}
}

// https://github.com/hit9/img2txt/blob/gh-pages/ansi.py#L4
func getANSIColor(c color.Color) string {
	r, g, b, _ := c.RGBA()

	websafeR := int(math.Round((float64(r) / 65535.0) * 5))
	websafeG := int(math.Round((float64(g) / 65535.0) * 5))
	websafeB := int(math.Round((float64(b) / 65535.0) * 5))

	return strconv.Itoa(int(((websafeR * 36) + (websafeG * 6) + websafeB) + 16))
}

func getTrueColor(c color.Color) string {
	r, g, b, _ := c.RGBA()

	cR := int(r / 256)
	cG := int(g / 256)
	cB := int(b / 256)

	return fmt.Sprintf("%d;%d;%d", cR, cG, cB)
}
