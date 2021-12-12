package ascii

import (
	"image/color"
	"testing"

	"golang.org/x/image/colornames"
)

func TestImage_GetCharForColor(t *testing.T) {
	type args struct {
		c color.Color
	}
	tests := []struct {
		name string
		i    *Image
		args args
		want string
	}{
		{
			"Black color returns M",
			NewImage(),
			args{
				colornames.Black,
			},
			"@",
		},
		{
			"White color returns space",
			NewImage(),
			args{
				colornames.Green,
			},
			"#",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.getCharForColor(tt.args.c); got != tt.want {
				t.Errorf("Image.getCharForColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
