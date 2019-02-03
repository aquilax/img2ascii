package img2ascii

import (
	"image"
	"io"
)

// Converter is color to string interface
type Converter interface {
	Process(img image.Image, out io.WriteCloser) error
	GetFontRatio() float64
}
