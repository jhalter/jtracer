package jtracer

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"math"
	"testing"
)

func TestNewCamera(t *testing.T) {
	type args struct {
		hsize float64
		vsize float64
		fov   float64
	}
	tests := []struct {
		name string
		args args
		want Camera
	}{
		{
			name: "constructing a new camera",
			args: args{
				hsize: 160,
				vsize: 120,
				fov:   math.Pi / 2,
			},
			want: Camera{
				Hsize:      160,
				Vsize:      120,
				Fov:        math.Pi / 2,
				Transform:  IdentityMatrix,
				HalfWidth:  1,
				HalfHeight: 0.75,
				PixelSize:  0.0125,
			},
		},
		{
			name: "horizontal canvas",
			args: args{
				hsize: 200,
				vsize: 125,
				fov:   math.Pi / 2,
			},
			want: Camera{
				Hsize:      200,
				Vsize:      125,
				Fov:        math.Pi / 2,
				Transform:  IdentityMatrix,
				HalfWidth:  1,
				HalfHeight: 0.625,
				PixelSize:  0.01,
			},
		},
		{
			name: "vertical canvas",
			args: args{
				hsize: 200,
				vsize: 125,
				fov:   math.Pi / 2,
			},
			want: Camera{
				Hsize:      200,
				Vsize:      125,
				Fov:        math.Pi / 2,
				Transform:  IdentityMatrix,
				HalfWidth:  1,
				HalfHeight: 0.625,
				PixelSize:  0.01,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCamera(tt.args.hsize, tt.args.vsize, tt.args.fov); !cmp.Equal(got, tt.want, float64Comparer) {
				t.Errorf("NewCamera() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamera_RayForPixel(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	spew.Dump(c)
	type fields struct {
		Hsize      float64
		Vsize      float64
		Fov        float64
		Transform  Matrix
		HalfWidth  float64
		HalfHeight float64
		PixelSize  float64
	}
	type args struct {
		px float64
		py float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Ray
	}{
		{
			name: "constructing a ray through the center of the canvas",
			fields: fields{
				Hsize:      201,
				Vsize:      101,
				Fov:        math.Pi / 2,
				Transform:  IdentityMatrix,
				HalfWidth:  1,
				HalfHeight: 0.50248756,
				PixelSize:  0.00995024,
			},
			args: args{
				px: 100,
				py: 50,
			},
			want: Ray{
				Origin:    *NewPoint(0, 0, 0),
				Direction: *NewVector(0, 0, -1),
			},
		},
		{
			name: "constructing a ray through the corner of the canvas",
			fields: fields{
				Hsize:      201,
				Vsize:      101,
				Fov:        math.Pi / 2,
				Transform:  IdentityMatrix,
				HalfWidth:  1,
				HalfHeight: 0.50248756,
				PixelSize:  0.00995024,
			},
			args: args{
				px: 0,
				py: 0,
			},
			want: Ray{
				Origin:    *NewPoint(0, 0, 0),
				Direction: *NewVector(0.66519, 0.33259, -0.66851),
			},
		},
		{
			name: "constructing a ray when the camera is transformed",
			fields: fields{
				Hsize:      201,
				Vsize:      101,
				Fov:        math.Pi / 2,
				Transform:  RotationY(math.Pi / 4).Multiply(NewTranslation(0, -2, 5)),
				HalfWidth:  1,
				HalfHeight: 0.50248756,
				PixelSize:  0.00995024,
			},
			args: args{
				px: 100,
				py: 50,
			},
			want: Ray{
				Origin:    *NewPoint(0, 2, -5),
				Direction: *NewVector(math.Sqrt(2)/2, 0, -math.Sqrt(2)/2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Camera{
				Hsize:      tt.fields.Hsize,
				Vsize:      tt.fields.Vsize,
				Fov:        tt.fields.Fov,
				Transform:  tt.fields.Transform,
				HalfWidth:  tt.fields.HalfWidth,
				HalfHeight: tt.fields.HalfHeight,
				PixelSize:  tt.fields.PixelSize,
			}
			if got := c.RayForPixel(tt.args.px, tt.args.py); !cmp.Equal(got, tt.want, float64Comparer) {
				t.Errorf("RayForPixel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamera_Render(t *testing.T) {
	c := NewCamera(11, 11, math.Pi/2)
	spew.Dump(c)

	type fields struct {
		Hsize      float64
		Vsize      float64
		Fov        float64
		Transform  Matrix
		HalfWidth  float64
		HalfHeight float64
		PixelSize  float64
	}
	type args struct {
		w World
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Canvas
	}{
		{
			name: "rendering a world with a camera",
			fields: fields{
				Hsize: 11,
				Vsize: 11,
				Fov:   math.Pi / 2,
				Transform: ViewTransform(
					NewPoint(0, 0, -5),
					NewPoint(0, 0, 0),
					NewVector(0, 1, 0),
				),
				HalfWidth:  1,
				HalfHeight: 1,
				PixelSize:  0.181818,
			},
			args: args{
				w: DefaultWorld(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Camera{
				Hsize:      tt.fields.Hsize,
				Vsize:      tt.fields.Vsize,
				Fov:        tt.fields.Fov,
				Transform:  tt.fields.Transform,
				HalfWidth:  tt.fields.HalfWidth,
				HalfHeight: tt.fields.HalfHeight,
				PixelSize:  tt.fields.PixelSize,
			}

			got := c.Render(tt.args.w)
			want := &Color{0.38066, 0.47583, 0.2855}
			if !cmp.Equal(got.PixelAt(5, 5), want, float64Comparer) {
				t.Errorf("Render() = %v, want %v", got.PixelAt(5, 5), want)
			}

			// if got := c.Render(tt.args.w); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Render() = %v, want %v", got, tt.want)
			// }
		})
	}
}
