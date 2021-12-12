package ansi

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
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := NewImage(Palette256, SingleHeight)
			if got := ac.getCharForColor(tt.args.c); got != tt.want {
				t.Errorf("Image256Colors.GetCharForColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImage256Colors_GetCharForColorCaching(t *testing.T) {
	ansi := NewImage(Palette256, SingleHeight)

	r1Want := "\x1b[48;5;196m "
	if r1 := ansi.getCharForColor(colornames.Red); r1 != r1Want {
		t.Errorf("Image256Colors.GetCharForColorFirst() = %v, want %v", r1, r1Want)
	}
	r2Want := " "

	if r2 := ansi.getCharForColor(colornames.Red); r2 != r2Want {
		t.Errorf("Image256Colors.GetCharForColorSecond() = %v, want %v", r2, r2Want)
	}
}
