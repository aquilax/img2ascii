package img2ascii

import (
	"image/color"
	"testing"

	"golang.org/x/image/colornames"
)

func TestAsciiNoColor_GetCharForColor(t *testing.T) {
	type args struct {
		c color.Color
	}
	tests := []struct {
		name string
		anc  Converter
		args args
		want string
	}{
		{
			"Black color returns M",
			AsciiNoColor{},
			args{
				colornames.Black,
			},
			"M",
		},
		{
			"White color returns space",
			AsciiNoColor{},
			args{
				colornames.Green,
			},
			"O",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			anc := AsciiNoColor{}
			if got := anc.GetCharForColor(tt.args.c); got != tt.want {
				t.Errorf("AsciiNoColor.GetCharForColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
