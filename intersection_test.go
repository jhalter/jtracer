package jtracer

import (
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
					T: 1, Object: Sphere{id: 1},
				},
				{
					T: 2, Object: Sphere{id: 1},
				},
			},
			want: &Intersection{
				T: 1, Object: Sphere{id: 1},
			},
		},
		{
			name: "the hit, when some intersections have negative t",
			i: Intersections{
				{
					T: -1, Object: Sphere{id: 1},
				},
				{
					T: 2, Object: Sphere{id: 1},
				},
			},
			want: &Intersection{
				T: 2, Object: Sphere{id: 1},
			},
		},
		{
			name: "the hit, when all intersections have negative t",
			i: Intersections{
				{
					T: -2, Object: Sphere{id: 1},
				},
				{
					T: -1, Object: Sphere{id: 1},
				},
			},
			want: nil,
		},
		{
			name: "the hit is always the lowest non-negative intersection",
			i: Intersections{
				{
					T: 5, Object: Sphere{id: 1},
				},
				{
					T: 7, Object: Sphere{id: 1},
				},
				{
					T: -3, Object: Sphere{id: 1},
				},
				{
					T: 2, Object: Sphere{id: 1},
				},
			},
			want: &Intersection{
				T: 2, Object: Sphere{id: 1},
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
