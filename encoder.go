package img2ascii

import (
	"image"
	"io"
)

// SingleHeightRatio represents the width/height ratio of a character for single row pixel height
const SingleHeightRatio = 9 / 22.

// HalfHeightRatio represents the width/height ratio of a character for half row pixel height
const HalfHeightRatio = 9 / 11.

// Encoder is the interface implemented by all converters
type Encoder interface {
	Encode(w io.Writer, m image.Image) error
	GetFontRatio() float64
}
