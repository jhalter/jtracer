package jtracer

import (
	"github.com/google/go-cmp/cmp"
	"math"
	"reflect"
	"testing"
)

func TestSphere_Intersects(t *testing.T) {
	type fields struct {
		id        int
		Transform Matrix
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
			name: "a ray intersects a sphere at two points",
			fields: fields{
				id:        1,
				Transform: IdentityMatrix,
			},
			args: args{
				r: Ray{
					Origin:    *NewPoint(0, 0, -5),
					Direction: *NewVector(0, 0, 1),
				},
			},
			want: Intersections{
				{T: 4.0, Object: Sphere{ID: 1, Transform: IdentityMatrix}},
				{T: 6.0, Object: Sphere{ID: 1, Transform: IdentityMatrix}},
			},
		},
		{
			name: "a ray intersects a sphere at a tangent",
			fields: fields{
				id:        1,
				Transform: IdentityMatrix,
			},
			args: args{
				r: Ray{
					Origin:    *NewPoint(0, 1, -5),
					Direction: *NewVector(0, 0, 1),
				},
			},
			want: Intersections{
				{T: 5, Object: Sphere{ID: 1, Transform: IdentityMatrix}},
				{T: 5, Object: Sphere{ID: 1, Transform: IdentityMatrix}},
			},
		},
		{
			name: "a ray misses a sphere",
			fields: fields{
				id:        1,
				Transform: IdentityMatrix,
			},
			args: args{
				r: Ray{
					Origin:    *NewPoint(0, 2, -5),
					Direction: *NewVector(0, 0, 1),
				},
			},
			want: Intersections{},
		},
		{
			name: "a ray originates inside a sphere",
			fields: fields{
				id:        1,
				Transform: IdentityMatrix,
			},
			args: args{
				r: Ray{
					Origin:    *NewPoint(0, 0, 0),
					Direction: *NewVector(0, 0, 1),
				},
			},
			want: Intersections{
				{T: -1, Object: Sphere{ID: 1, Transform: IdentityMatrix}},
				{T: 1, Object: Sphere{ID: 1, Transform: IdentityMatrix}},
			},
		},
		{
			name: "a sphere is behind a ray",
			fields: fields{
				id:        1,
				Transform: IdentityMatrix,
			},
			args: args{
				r: Ray{
					Origin:    *NewPoint(0, 0, 5),
					Direction: *NewVector(0, 0, 1),
				},
			},
			want: Intersections{
				{T: -6, Object: Sphere{ID: 1, Transform: IdentityMatrix}},
				{T: -4, Object: Sphere{ID: 1, Transform: IdentityMatrix}},
			},
		},
		{
			name: "intersecting a scaled sphere with a ray",
			fields: fields{
				id:        1,
				Transform: Scaling(2, 2, 2),
			},
			args: args{
				r: Ray{
					Origin:    *NewPoint(0, 0, -5),
					Direction: *NewVector(0, 0, 1),
				},
			},
			want: Intersections{
				{T: 3, Object: Sphere{ID: 1, Transform: Scaling(2, 2, 2)}},
				{T: 7, Object: Sphere{ID: 1, Transform: Scaling(2, 2, 2)}},
			},
		},
		{
			name: "intersecting a translated sphere with a ray",
			fields: fields{
				id:        1,
				Transform: NewTranslation(5, 0, 0),
			},
			args: args{
				r: Ray{
					Origin:    *NewPoint(0, 0, -5),
					Direction: *NewVector(0, 0, 1),
				},
			},
			want: Intersections{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sphere{
				ID:        tt.fields.id,
				Transform: tt.fields.Transform,
			}
			if got := s.Intersects(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSphere_NormalAt(t *testing.T) {
	type fields struct {
		id        int
		Transform Matrix
	}
	type args struct {
		worldPoint Tuple
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Tuple
	}{
		{
			name: "the normal on a sphere at a point on the x axis",
			fields: fields{
				Transform: IdentityMatrix,
			},
			args: args{*NewPoint(1, 0, 0)},
			want: *NewVector(1, 0, 0),
		},
		{
			name: "the normal on a sphere at a point on the y axis",
			fields: fields{
				Transform: IdentityMatrix,
			},
			args: args{*NewPoint(0, 1, 0)},
			want: *NewVector(0, 1, 0),
		},
		{
			name: "the normal on a sphere at a point on the z axis",
			fields: fields{
				Transform: IdentityMatrix,
			},
			args: args{*NewPoint(0, 0, 1)},
			want: *NewVector(0, 0, 1),
		},
		{
			name: "the normal on a sphere at a non-axial point",
			fields: fields{
				Transform: IdentityMatrix,
			},
			args: args{*NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)},
			want: *NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
		},
		{
			name: "computing the normal on a translated sphere",
			fields: fields{
				Transform: NewTranslation(0, 1, 0),
			},
			args: args{*NewPoint(0, 1.70711, -0.70711)},
			want: *NewVector(0, 0.70711, -0.70711),
		},
		{
			name: "computing the normal on a transformed sphere",
			fields: fields{
				Transform: Scaling(1, 0.5, 1).Multiply(RotationZ(math.Pi / 5)),
			},

			args: args{*NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)},
			want: *NewVector(0, 0.97014, -0.24254),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sphere{
				ID:        tt.fields.id,
				Transform: tt.fields.Transform,
			}
			if got := s.NormalAt(tt.args.worldPoint); !cmp.Equal(got, tt.want, float64Comparer) {
				t.Errorf("NormalAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
