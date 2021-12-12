package img2ascii

import (
	"image/color"
	"reflect"
	"testing"

	"golang.org/x/image/colornames"
)

func Test_getFgBgColors(t *testing.T) {
	t.Skip()
	type args struct {
		ul color.Color
		ur color.Color
		ll color.Color
		lr color.Color
	}
	tests := []struct {
		name      string
		args      args
		wantFg    color.Color
		wantBg    color.Color
		wantIndex int
	}{
		{
			"test1",
			args{
				colornames.Red,
				colornames.Red,
				colornames.Blue,
				colornames.Blue,
			},
			colornames.Red,
			colornames.Blue,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFg, gotBg, gotIndex := getFgBgColors(tt.args.ul, tt.args.ur, tt.args.ll, tt.args.lr)
			if !reflect.DeepEqual(gotFg, tt.wantFg) {
				t.Errorf("getFgBgColors() gotFg = %v, want %v", gotFg, tt.wantFg)
			}
			if !reflect.DeepEqual(gotBg, tt.wantBg) {
				t.Errorf("getFgBgColors() gotBg = %v, want %v", gotBg, tt.wantBg)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("getFgBgColors() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}
