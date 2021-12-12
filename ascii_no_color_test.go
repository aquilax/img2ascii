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
		anc  *AsciiNoColor
		args args
		want string
	}{
		{
			"Black color returns M",
			NewAsciiNoColor(DefaultAsciiPalette),
			args{
				colornames.Black,
			},
			"@",
		},
		{
			"White color returns space",
			NewAsciiNoColor(DefaultAsciiPalette),
			args{
				colornames.Green,
			},
			"#",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			anc := tt.anc
			if got := anc.GetCharForColor(tt.args.c); got != tt.want {
				t.Errorf("AsciiNoColor.GetCharForColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
