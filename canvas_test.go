package jtracer

import (
	"reflect"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want *Canvas
	}{
		{
			name: "creating a canvas",
			args: args{
				width:  10,
				height: 20,
			},
			want: &Canvas{
				Width:  10,
				Height: 20,
				Data: func() [][]Color {
					data := make([][]Color, 20)
					for i := range data {
						data[i] = make([]Color, 10)
					}
					return data
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCanvas(tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCanvas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanvas_WritePixel(t *testing.T) {
	type fields struct {
		Width  int
		Height int
		Data   [][]Color
	}
	type args struct {
		x     int
		y     int
		color Color
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "writing pixels to a canvas",
			fields: fields{
				Width:  10,
				Height: 20,
				Data: func() [][]Color {
					data := make([][]Color, 20)
					for i := range data {
						data[i] = make([]Color, 10)
					}
					return data
				}(),
			},
			args: args{
				x:     2,
				y:     3,
				color: Red,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Canvas{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
				Data:   tt.fields.Data,
			}
			c.WritePixel(tt.args.x, tt.args.y, &tt.args.color)
			c.PixelAt(2, 3).Equals(&Red)
		})
	}
}
