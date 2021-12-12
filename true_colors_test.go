package img2ascii

import (
	"image/color"
	"testing"

	"golang.org/x/image/colornames"
)

func Test_getTrueColor(t *testing.T) {
	type args struct {
		c color.Color
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"red",
			args{
				c: colornames.Red,
			},
			"255;0;0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTrueColor(tt.args.c); got != tt.want {
				t.Errorf("getTrueColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
