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
