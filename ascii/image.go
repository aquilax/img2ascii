package ascii

import (
	"bytes"
	"image"
	"image/color"
	"io"

	"github.com/aquilax/img2ascii"
)

// DefaultASCIIPalette contains the default palette used for the ascii encoding
//
// Source from http://paulbourke.net/dataformats/asciiart/
const DefaultASCIIPalette = "@%#*+=-:. "

// Image converts images to pseudo greyscale text reprosentation
type Image struct {
	palette     string
	paletteSize int
}

// NewImage creates new ASCII image encoder using the default palette
func NewImage() *Image {
	return &Image{
		DefaultASCIIPalette,
		len(DefaultASCIIPalette),
	}
}

// NewImageCustomPalette creates new ASCII image encoder using custom palette
func NewImageCustomPalette(palette string) *Image {
	return &Image{
		palette,
		len(palette),
	}
}

// GetFontRatio returns the font ratio
func (i Image) GetFontRatio() float64 {
	return img2ascii.SingleHeightRatio
}

// Encode converts image to text and writes it to out
func (i Image) Encode(out io.Writer, img image.Image) error {
	bounds := img.Bounds()
	height := bounds.Dy()
	width := bounds.Dx()

	var buffer bytes.Buffer
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			buffer.WriteString(i.getCharForColor(img.At(x, y)))
		}
		buffer.WriteString("\n")
		io.WriteString(out, buffer.String())
		buffer.Reset()
	}
	return nil
}

// GetCharForColor returns Image string representation of color
func (i Image) getCharForColor(c color.Color) string {
	gray := color.GrayModel.Convert(c)
	y := int(gray.(color.Gray).Y)
	pos := int(y * (i.paletteSize - 1) / 255)
	return string(i.palette[pos])
}
