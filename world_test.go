package jtracer

import (
	"reflect"
	"testing"
)

var dw = DefaultWorld()

func TestWorld_Intersect(t *testing.T) {

	type fields struct {
		Objects []Shape
		Light   Light
	}
	type args struct {
		r Ray
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Intersections
	}{
		{
			name: "intersect a world with a ray",
			fields: fields{
				Objects: dw.Objects,
				Light:   dw.Light,
			},
			args: args{
				r: Ray{
					Origin:    *NewPoint(0, 0, -5),
					Direction: *NewVector(0, 0, 1),
				},
			},
			want: Intersections{
				{
					T:      4,
					Object: dw.Objects[0],
				},
				{
					T:      4.5,
					Object: dw.Objects[1],
				},
				{
					T:      5.5,
					Object: dw.Objects[1],
				},
				{
					T:      6,
					Object: dw.Objects[0],
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := World{
				Objects: tt.fields.Objects,
				Light:   tt.fields.Light,
			}
			if got := w.Intersect(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorld_ShadeHit(t *testing.T) {
	type fields struct {
		Objects []Shape
		Light   Light
	}
	type args struct {
		comps Computations
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Color
	}{
		{
			name: "shading an intersection",
			fields: fields{
				Objects: dw.Objects,
				Light:   dw.Light,
			},
			args: args{
				comps: func() Computations {
					i := Intersection{
						T:      4,
						Object: dw.Objects[0],
					}
					return i.PrepareComputations(
						Ray{
							Origin:    *NewPoint(0, 0, -5),
							Direction: *NewVector(0, 0, 1),
						},
					)
				}(),
			},
			want: Color{
				Red:   0.38066,
				Green: 0.47583,
				Blue:  0.2855,
			},
		},
		{
			name: "ShadeHit() is given an intersection in shadow",
			fields: fields{
				Objects: []Shape{
					NewSphere(),
					Sphere{
						Transform: NewTranslation(0, 0, 10),
						Material:  NewMaterial(),
					},
				},
				Light: NewPointLight(*NewPoint(0, 0, -10), White),
			},
			args: args{
				comps: func() Computations {
					i := Intersection{
						T: 4,
						Object: Sphere{
							Transform: NewTranslation(0, 0, 10),
							Material:  NewMaterial(),
						},
					}
					return i.PrepareComputations(
						Ray{
							Origin:    *NewPoint(0, 0, 5),
							Direction: *NewVector(0, 0, 1),
						},
					)
				}(),
			},
			want: Color{
				Red:   0.1,
				Green: 0.1,
				Blue:  0.1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := World{
				Objects: tt.fields.Objects,
				Light:   tt.fields.Light,
			}
			if got := w.ShadeHit(tt.args.comps); !got.Equals(&tt.want) {
				t.Errorf("ShadeHit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorld_IsShadowed(t *testing.T) {
	type fields struct {
		Objects []Shape
		Light   Light
	}
	type args struct {
		p Tuple
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "there is no shadow when nothing is collinear with the point and light",
			fields: fields{
				Objects: dw.Objects,
				Light:   dw.Light,
			},
			args: args{p: *NewPoint(0, 10, 0)},
			want: false,
		},
		{
			name: "the shadow when an object is between the point and light",
			fields: fields{
				Objects: dw.Objects,
				Light:   dw.Light,
			},
			args: args{p: *NewPoint(10, -10, 10)},
			want: true,
		},
		{
			name: "there is no shadow when an object is behind the light",
			fields: fields{
				Objects: dw.Objects,
				Light:   dw.Light,
			},
			args: args{p: *NewPoint(-20, 20, -20)},
			want: false,
		},
		{
			name: "there is no shadow when an object is behind the point",
			fields: fields{
				Objects: dw.Objects,
				Light:   dw.Light,
			},
			args: args{p: *NewPoint(-2, 2, -2)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := World{
				Objects: tt.fields.Objects,
				Light:   tt.fields.Light,
			}
			if got := w.IsShadowed(tt.args.p); got != tt.want {
				t.Errorf("IsShadowed() = %v, want %v", got, tt.want)
			}
		})
	}
}
