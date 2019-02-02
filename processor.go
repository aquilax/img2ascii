package img2ascii

import (
	"image"
	"image/color"
	"io"
)

// Converter is color to string interface
type Converter interface {
	GetImage(img image.Image) (image.Image, error)
	GetCharForColor(c color.Color) string
	GetEndLine() string
	GetEndOfImage() string
}

func Process(img image.Image, out io.WriteCloser, c Converter) error {
	// Allow the converter to pre-process the image
	image, err := c.GetImage(img)
	if err != nil {
		return err
	}
	bounds := image.Bounds()
	height := bounds.Max.Y
	width := bounds.Max.X

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			io.WriteString(out, c.GetCharForColor(image.At(x, y)))
		}
		io.WriteString(out, c.GetEndLine())
	}
	io.WriteString(out, c.GetEndOfImage())
	return nil
}
