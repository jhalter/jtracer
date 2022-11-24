package jtracer

import (
	"reflect"
	"testing"
)

func TestPlane_NormalAt(t *testing.T) {
	type fields struct {
		ID    int
		Shape Shape
	}
	type args struct {
		in0 Tuple
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Tuple
	}{
		{
			name: "the normal of a plane is constant everywhere",
			fields: fields{
				ID:    0,
				Shape: Shape{},
			},
			args: args{
				*NewPoint(0, 0, 0),
			},
			want: *NewVector(0, 1, 0),
		},
		{
			name: "the normal of a plane is constant everywhere",
			fields: fields{
				ID:    0,
				Shape: Shape{},
			},
			args: args{
				*NewPoint(10, 0, -10),
			},
			want: *NewVector(0, 1, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Plane{
				ID:    tt.fields.ID,
				Shape: tt.fields.Shape,
			}
			if got := p.NormalAt(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NormalAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlane_Intersects(t *testing.T) {
	type fields struct {
		ID    int
		Shape Shape
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
			name: "intersect with a ray parallel to the plane",
			fields: fields{
				Shape: Shape{},
			},
			args: args{
				Ray{
					Origin:    *NewPoint(0, 10, 0),
					Direction: *NewVector(0, 0, 1),
				},
			},
			want: Intersections{},
		},
		{
			name: "intersect with a coplanar ray",
			fields: fields{
				Shape: Shape{},
			},
			args: args{
				Ray{
					Origin:    *NewPoint(0, 0, 0),
					Direction: *NewVector(0, 0, 1),
				},
			},
			want: Intersections{},
		},
		{
			name: "a ray intersecting a plane from above",
			fields: fields{
				ID:    1,
				Shape: Shape{},
			},
			args: args{
				Ray{
					Origin:    *NewPoint(0, 1, 0),
					Direction: *NewVector(0, -1, 0),
				},
			},
			want: Intersections{
				{
					T:      1,
					Object: Plane{ID: 1},
				},
			},
		},
		{
			name: "a ray intersecting a plane from below",
			fields: fields{
				ID:    1,
				Shape: Shape{},
			},
			args: args{
				Ray{
					Origin:    *NewPoint(0, -1, 0),
					Direction: *NewVector(0, 1, 0),
				},
			},
			want: Intersections{
				{
					T:      1,
					Object: Plane{ID: 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Plane{
				ID:    tt.fields.ID,
				Shape: tt.fields.Shape,
			}
			if got := p.Intersects(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersects() = %v, want %v", got, tt.want)
			}
		})
	}
}
