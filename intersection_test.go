package jtracer

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"math"
	"reflect"
	"testing"
)

func TestIntersections_Hit(t *testing.T) {
	tests := []struct {
		name string
		i    Intersections
		want *Intersection
	}{
		{
			name: "the hit, when all intersections have positive t",
			i: Intersections{
				{
					T: 1, Object: Sphere{ID: 1},
				},
				{
					T: 2, Object: Sphere{ID: 1},
				},
			},
			want: &Intersection{
				T: 1, Object: Sphere{ID: 1},
			},
		},
		{
			name: "the hit, when some intersections have negative t",
			i: Intersections{
				{
					T: -1, Object: Sphere{ID: 1},
				},
				{
					T: 2, Object: Sphere{ID: 1},
				},
			},
			want: &Intersection{
				T: 2, Object: Sphere{ID: 1},
			},
		},
		{
			name: "the hit, when all intersections have negative t",
			i: Intersections{
				{
					T: -2, Object: Sphere{ID: 1},
				},
				{
					T: -1, Object: Sphere{ID: 1},
				},
			},
			want: nil,
		},
		{
			name: "the hit is always the lowest non-negative intersection",
			i: Intersections{
				{
					T: 5, Object: Sphere{ID: 1},
				},
				{
					T: 7, Object: Sphere{ID: 1},
				},
				{
					T: -3, Object: Sphere{ID: 1},
				},
				{
					T: 2, Object: Sphere{ID: 1},
				},
			},
			want: &Intersection{
				T: 2, Object: Sphere{ID: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.i.Hit()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersection_PrepareComputations(t *testing.T) {
	s := NewSphere()
	type fields struct {
		T      float64
		Object Shaper
	}
	type args struct {
		r Ray
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Computations
	}{
		{
			name: "precomputing the state of an intersection",
			fields: fields{
				T:      4,
				Object: s,
			},
			args: args{
				r: Ray{
					Origin:    *NewPoint(0, 0, -5),
					Direction: *NewVector(0, 0, 1),
				},
			},
			want: Computations{
				T:         4,
				Object:    s,
				Point:     *NewPoint(0, 0, -1),
				Eyev:      *NewVector(0, 0, -1),
				Normalv:   *NewVector(0, 0, -1),
				Inside:    false,
				OverPoint: *NewPoint(0, 0, -1.00001),
				Reflectv:  *NewVector(0, 0, -1),
			},
		},
		{
			name: "the hit, when an intersection occurs on the inside",
			fields: fields{
				T:      1,
				Object: s,
			},
			args: args{
				r: Ray{
					Origin:    *NewPoint(0, 0, 0),
					Direction: *NewVector(0, 0, 1),
				},
			},
			want: Computations{
				T:         1,
				Object:    s,
				Point:     *NewPoint(0, 0, 1),
				Eyev:      *NewVector(0, 0, -1),
				Normalv:   *NewVector(0, 0, -1),
				Inside:    true,
				OverPoint: *NewPoint(0, 0, 1),
				Reflectv:  *NewVector(0, 0, -1),
			},
		},
		{
			name: "precomputing the reflection vector",
			fields: fields{
				T:      math.Sqrt(2),
				Object: Plane{},
			},
			args: args{
				r: Ray{
					Origin:    *NewPoint(0, 1, -1),
					Direction: *NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2),
				},
			},
			want: Computations{
				T:         1.4142135623730951,
				Object:    Plane{},
				Point:     *NewPoint(0, 0, 0),
				Eyev:      *NewVector(0, 0.7071067811865476, -0.7071067811865476),
				Normalv:   *NewVector(0, 1, 0),
				Inside:    false,
				OverPoint: *NewPoint(0, 0, 0),
				Reflectv:  *NewVector(0, math.Sqrt(2)/2, math.Sqrt(2)/2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Intersection{
				T:      tt.fields.T,
				Object: tt.fields.Object,
			}

			if got := i.PrepareComputations(tt.args.r); !cmp.Equal(got, tt.want, float64Comparer) {
				fmt.Println(cmp.Diff(got, tt.want))
				t.Errorf("PrepareComputations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersection_PrecomputingReflectionVector(t *testing.T) {
	type fields struct {
		T      float64
		Object Shaper
	}
	type args struct {
		r Ray
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Tuple
	}{
		{
			name: "precomputing the reflection vector",
			fields: fields{
				T:      math.Sqrt(2),
				Object: Plane{},
			},
			args: args{
				r: Ray{
					Origin:    *NewPoint(0, 1, -1),
					Direction: *NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2),
				},
			},
			want: *NewVector(0, math.Sqrt(2)/2, math.Sqrt(2)/2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Intersection{
				T:      tt.fields.T,
				Object: tt.fields.Object,
			}
			if got := i.PrepareComputations(tt.args.r); !reflect.DeepEqual(got.Reflectv, tt.want) {
				t.Errorf("PrepareComputations() = %v, want %v", got, tt.want)
			}
		})
	}
}
