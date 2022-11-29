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
				T:          4,
				Object:     s,
				Point:      *NewPoint(0, 0, -1),
				Eyev:       *NewVector(0, 0, -1),
				Normalv:    *NewVector(0, 0, -1),
				Inside:     false,
				OverPoint:  *NewPoint(0, 0, -1.00001),
				UnderPoint: *NewPoint(0, 0, -0.99999),

				Reflectv: *NewVector(0, 0, -1),
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
				T:          1,
				Object:     s,
				Point:      *NewPoint(0, 0, 1),
				Eyev:       *NewVector(0, 0, -1),
				Normalv:    *NewVector(0, 0, -1),
				Inside:     true,
				OverPoint:  *NewPoint(0, 0, 1),
				UnderPoint: *NewPoint(0, 0, 1.00001),
				Reflectv:   *NewVector(0, 0, -1),
			},
		},
		{
			name: "precomputing the reflection vector",
			fields: fields{
				T: math.Sqrt(2),
				Object: Plane{
					Shape: Shape{Transform: IdentityMatrix},
				},
			},
			args: args{
				r: Ray{
					Origin:    *NewPoint(0, 1, -1),
					Direction: *NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2),
				},
			},
			want: Computations{
				T: 1.4142135623730951,
				Object: Plane{
					Shape: Shape{Transform: IdentityMatrix},
				},
				Point:      *NewPoint(0, 0, 0),
				Eyev:       *NewVector(0, 0.7071067811865476, -0.7071067811865476),
				Normalv:    *NewVector(0, 1, 0),
				Inside:     false,
				OverPoint:  *NewPoint(0, 0, 0),
				UnderPoint: *NewPoint(0, -0.00001, 0),
				Reflectv:   *NewVector(0, math.Sqrt(2)/2, math.Sqrt(2)/2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Intersection{
				T:      tt.fields.T,
				Object: tt.fields.Object,
			}

			if got := i.PrepareComputations(tt.args.r, nil); !cmp.Equal(got, tt.want, float64Comparer) {
				fmt.Println(cmp.Diff(got, tt.want, float64Comparer))
				t.Errorf("PrepareComputations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersection_FindingN1AndN2(t *testing.T) {
	a := NewGlassSphere()
	a.Shape.ID = 1
	a.Transform = Scaling(2, 2, 2)
	a.Material.RefractiveIndex = 1.5

	b := NewGlassSphere()
	b.Shape.ID = 2
	b.Transform = NewTranslation(0, 0, -0.25)
	b.Material.RefractiveIndex = 2.0

	c := NewGlassSphere()
	c.Shape.ID = 3
	c.Transform = NewTranslation(0, 0, 0.25)
	c.Material.RefractiveIndex = 2.5

	r := NewRay(
		*NewPoint(0, 0, -4),
		*NewVector(0, 0, 1),
	)
	xs := Intersections{
		{
			T:      2,
			Object: a,
		},
		{
			T:      2.75,
			Object: b,
		},
		{
			T:      3.25,
			Object: c,
		},
		{
			T:      4.75,
			Object: b,
		},
		{
			T:      5.25,
			Object: c,
		},
		{
			T:      6,
			Object: a,
		},
	}

	tests := []struct {
		n1 float64
		n2 float64
	}{
		{1, 1.5},
		{1.5, 2},
		{2, 2.5},
		{2.5, 2.5},
		{2.5, 1.5},
		{1.5, 1.0},
	}

	for j, tt := range tests {
		t.Run("tt.name", func(t *testing.T) {
			got := xs[j].PrepareComputations(r, xs)
			if !cmp.Equal(got.N1, tt.n1, float64Comparer) {
				t.Errorf("PrepareComputations() N1 = %v, want %v", got.N1, tt.n1)
			}

			if !cmp.Equal(got.N2, tt.n2, float64Comparer) {
				t.Errorf("PrepareComputations() N2 = %v, want %v", got.N2, tt.n2)
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
				Object: NewPlane(),
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
			if got := i.PrepareComputations(tt.args.r, nil); !reflect.DeepEqual(got.Reflectv, tt.want) {
				t.Errorf("PrepareComputations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchlick(t *testing.T) {
	glassSphere := NewGlassSphere()

	type args struct {
		comps Computations
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "the schlick approximation under total internal reflection",
			args: args{
				comps: func() Computations {
					xs := Intersections{
						{
							T:      -(math.Sqrt(2) / 2),
							Object: glassSphere,
						},
						{
							T:      (math.Sqrt(2) / 2),
							Object: glassSphere,
						},
					}

					return xs[1].PrepareComputations(
						NewRay(
							*NewPoint(0, 0, math.Sqrt(2)/2),
							*NewVector(0, 1, 0),
						),
						xs,
					)
				}(),
			},
			want: 1.0,
		},
		{
			name: "the Schlick approximation with a perpendicular viewing angle",
			args: args{
				comps: func() Computations {
					xs := Intersections{
						{
							T:      -1,
							Object: glassSphere,
						},
						{
							T:      1,
							Object: glassSphere,
						},
					}

					return xs[1].PrepareComputations(
						NewRay(
							*NewPoint(0, 0, 0),
							*NewVector(0, 1, 0),
						),
						xs,
					)
				}(),
			},
			want: 0.04,
		},
		{
			name: "the Schlick approximation with a small angle and n2 > n1",
			args: args{
				comps: func() Computations {
					xs := Intersections{
						{
							T:      1.8589,
							Object: glassSphere,
						},
					}
					return xs[0].PrepareComputations(
						NewRay(
							*NewPoint(0, 0.99, -2),
							*NewVector(0, 0, 1),
						),
						xs,
					)
				}(),
			},
			want: 0.48873,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Schlick(tt.args.comps); !cmp.Equal(tt.want, got, float64Comparer) {
				t.Errorf("Schlick() = %v, want %v", got, tt.want)
			}
		})
	}
}
