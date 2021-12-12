package img2ascii

import (
	"image"
	"io"
)

// Converter is the interface implemented by all converters
type Converter interface {
	Encode(w io.Writer, m image.Image) error
	GetFontRatio() float64
}
