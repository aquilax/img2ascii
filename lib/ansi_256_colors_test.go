package img2ascii

import (
	"image/color"
	"testing"

	"golang.org/x/image/colornames"
)

func TestANSI256Colors_GetCharForColor(t *testing.T) {
	type fields struct {
		lastColor int
	}
	type args struct {
		c color.Color
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := ANSI256Colors{
				lastColor: tt.fields.lastColor,
			}
			if got := ac.GetCharForColor(tt.args.c); got != tt.want {
				t.Errorf("ANSI256Colors.GetCharForColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestANSI256Colors_GetCharForColorCaching(t *testing.T) {
	ansi := &ANSI256Colors{}

	r1Want := "\x1b[48;5;196m "
	if r1 := ansi.GetCharForColor(colornames.Red); r1 != r1Want {
		t.Errorf("ANSI256Colors.GetCharForColorFirst() = %v, want %v", r1, r1Want)
	}
	r2Want := " "

	if r2 := ansi.GetCharForColor(colornames.Red); r2 != r2Want {
		t.Errorf("ANSI256Colors.GetCharForColorSecond() = %v, want %v", r2, r2Want)
	}
}
