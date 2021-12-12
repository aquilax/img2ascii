package img2ascii

import (
	"image/color"
)

var subPixels = [...]rune{
	' ', '▘', '▝', '▀', '▖', '▌', '▞', '▛',
	'▗', '▚', '▐', '▜', '▄', '▙', '▟', '█',
}

func getFgBgColors(ul, ur, ll, lr color.Color) (fg color.Color, bg color.Color, index int) {
	return nil, nil, 1
}

func QuixelToChar(ul, ur, ll, lr color.Color) (string, color.Color, color.Color) {
	// get colors
	fg, bg, index := getFgBgColors(ul, ur, ll, lr)
	return string(subPixels[index]), fg, bg
}
