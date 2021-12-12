package ansi

import (
	"bytes"
	"image"
	"image/color"
	"io"

	"github.com/aquilax/img2ascii"
)

// Image converts images to colored ANSI terminal text representations
type Image struct {
	colorPalette ColorPalette
	pixelHeight  PixelHeight
	colorGetter  colorConverter
	lastColor    string
}

// NewImage creates new Image converter
func NewImage(p ColorPalette, ph PixelHeight) *Image {
	return &Image{
		colorPalette: p,
		pixelHeight:  ph,
		colorGetter:  getColorGetter(p),
	}
}

// GetFontRatio returns the font ratio
func (i Image) GetFontRatio() float64 {
	if i.pixelHeight == SingleHeight {
		return img2ascii.SingleHeightRatio
	}
	return img2ascii.HalfHeightRatio
}

// Encode converts image to text and writes it to out
func (i *Image) Encode(out io.Writer, img image.Image) error {
	bounds := img.Bounds()
	height := bounds.Dy()
	width := bounds.Dx()

	var buffer bytes.Buffer
	if i.pixelHeight == SingleHeight {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				buffer.WriteString(i.getCharForColor(img.At(x, y)))
			}
			buffer.WriteString("\x1b[0m\n")
			i.lastColor = ""
			io.WriteString(out, buffer.String())
			buffer.Reset()
		}
	} else {
		for y := 0; y < height/2; y++ {
			for x := 0; x < width; x++ {
				buffer.WriteString(i.getCharForHalfColor(img.At(x, y*2), img.At(x, y*2+1)))
			}
			buffer.WriteString("\x1b[0m\n")
			io.WriteString(out, buffer.String())
			buffer.Reset()
		}
	}
	return nil
}

func (i *Image) getCharForColor(col color.Color) string {
	newColor := i.colorGetter(col)
	if i.lastColor == newColor {
		// Keep the existing color
		return emptyChar
	}
	i.lastColor = newColor
	if i.colorPalette == Palette256 {
		return "\x1b[48;5;" + i.lastColor + "m" + emptyChar
	}
	return "\x1b[48;2;" + i.lastColor + "m" + emptyChar
}

func (i Image) getCharForHalfColor(c1 color.Color, c2 color.Color) string {
	colorTop := i.colorGetter(c1)
	colorBottom := i.colorGetter(c2)
	if i.colorPalette == Palette256 {
		return "\x1b[38;5;" + colorTop + "m" + "\x1b[48;5;" + colorBottom + "m" + "▀"
	}
	return "\x1b[38;2;" + colorTop + "m" + "\x1b[48;2;" + colorBottom + "m" + "▀"
}
