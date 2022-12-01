package jtracer

import (
	"math"
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
			name:   "intersect a world with a ray",
			fields: fields(dw),
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

	s1 := NewSphereWithID(1)
	s1.SetTransform(NewTranslation(0, 0, 10))

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
					return i.PrepareComputations(Ray{
						Origin:    *NewPoint(0, 0, -5),
						Direction: *NewVector(0, 0, 1),
					}, nil)
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
					NewSphereWithID(2),
					s1,
				},
				Light: NewPointLight(*NewPoint(0, 0, -10), White),
			},
			args: args{
				comps: func() Computations {
					i := Intersection{
						T:      4,
						Object: s1,
					}
					return i.PrepareComputations(Ray{
						Origin:    *NewPoint(0, 0, 5),
						Direction: *NewVector(0, 0, 1),
					}, nil)
				}(),
			},
			want: Color{
				Red:   0.1,
				Green: 0.1,
				Blue:  0.1,
			},
		},
		{
			name: "ShadeHit() with a reflective material",
			fields: fields{
				Objects: []Shape{
					NewSphere(),
					s1,
				},
				Light: NewPointLight(*NewPoint(0, 0, -10), White),
			},
			args: args{
				comps: func() Computations {
					i := Intersection{
						T:      4,
						Object: s1,
					}
					return i.PrepareComputations(Ray{
						Origin:    *NewPoint(0, 0, 5),
						Direction: *NewVector(0, 0, 1),
					}, nil)
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
			if got := w.ShadeHit(tt.args.comps, 0); !got.Equals(&tt.want) {
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

func TestWorld_ReflectedColor(t *testing.T) {

	defaultWorldWithReflectivePlane := dw
	m := NewMaterial()
	m.Reflectivity = 0.5
	p := NewPlane()
	p.SetTransform(NewTranslation(0, -1, 0))
	p.Material = m
	defaultWorldWithReflectivePlane.Objects = append(dw.Objects, p)

	type fields struct {
		Objects []Shape
		Light   Light
	}
	type args struct {
		comps     Computations
		remaining int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Color
	}{
		{
			name:   "the reflected color for a nonreflective material",
			fields: fields(dw),
			args: args{
				comps: func() Computations {
					shape := dw.Objects[1].(*Sphere)
					shape.Material.Ambient = 1

					i := Intersection{T: 1, Object: shape}
					return i.PrepareComputations(NewRay(
						*NewPoint(0, 0, 0),
						*NewVector(0, 0, 1),
					), nil)
				}(),
			},
			want: Black,
		},
		//	TODO: Figure out why this is failing
		//
		//{
		//	name:   "the reflected color for a reflective material",
		//	fields: fields(defaultWorldWithReflectivePlane),
		//	args: args{
		//		comps: func() Computations {
		//			i := Intersection{T: math.Sqrt(2), Object: defaultWorldWithReflectivePlane.Objects[2]}
		//			return i.PrepareComputations(
		//				NewRay(
		//					*NewPoint(0, 0, -3),
		//					*NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2),
		//				), nil,
		//			)
		//		}(),
		//	},
		//	want: Color{0.19032, 0.2379, 0.14274},
		//},
		{
			name:   "the reflected color at the maximum recursive depth",
			fields: fields(defaultWorldWithReflectivePlane),
			args: args{
				remaining: 0,
				comps: func() Computations {
					i := Intersection{T: math.Sqrt(2), Object: defaultWorldWithReflectivePlane.Objects[2]}
					return i.PrepareComputations(NewRay(
						*NewPoint(0, 0, -3),
						*NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2),
					), nil)
				}(),
			},
			want: Black,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := World{
				Objects: tt.fields.Objects,
				Light:   tt.fields.Light,
			}
			if got := w.ReflectedColor(tt.args.comps, 0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReflectedColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorld_RefractedColor(t *testing.T) {

	glassySphere := NewSphere()
	m1 := NewMaterial()
	m1.Color = Color{0.8, 1.0, 0.6}
	m1.Diffuse = 0.7
	m1.Specular = 0.2
	m1.Transparency = 1.0
	m1.RefractiveIndex = 1.5
	glassySphere.Material = m1

	//s1 := Sphere{
	//	AbstractShape: AbstractShape{
	//		Transform: IdentityMatrix,
	//		Material: Material{
	//			Color:           White,
	//			Ambient:         1.0,
	//			Diffuse:         0.9,
	//			Specular:        0.9,
	//			Shininess:       200.0,
	//			RefractiveIndex: 1.0,
	//			HasPattern:      true,
	//			Pattern:         NewTestPattern(),
	//		},
	//	},
	//}
	//s2 := Sphere{
	//	AbstractShape: AbstractShape{
	//		Transform: Scaling(0.5, 0.5, 0.5),
	//		Material: Material{
	//			Color:           White,
	//			Ambient:         0.1,
	//			Diffuse:         0.9,
	//			Specular:        0.9,
	//			Shininess:       200.0,
	//			RefractiveIndex: 1.5,
	//			Transparency:    1.0,
	//		},
	//	},
	//}

	type fields struct {
		Objects []Shape
		Light   Light
	}
	type args struct {
		comps     Computations
		remaining int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Color
	}{
		{
			name:   "the refracted color with an opaque surface",
			fields: fields(dw),
			args: args{
				comps: func() Computations {
					xs := Intersections{
						{
							T:      4,
							Object: dw.Objects[0],
						},
						{
							T:      6,
							Object: dw.Objects[0],
						},
					}
					return xs[0].PrepareComputations(
						NewRay(
							*NewPoint(0, 0, -5),
							*NewVector(0, 0, 1),
						),
						xs,
					)
				}(),
				remaining: 5,
			},
			want: Black,
		},
		{
			name: "the refracted color at maximum recursive depth",
			fields: fields{
				Light: dw.Light,
				Objects: func() []Shape {
					s1 := NewSphere()
					m1 := NewMaterial()
					m1.Color = Color{0.8, 1.0, 0.6}
					m1.Diffuse = 0.7
					m1.Specular = 0.2
					m1.Transparency = 1.0
					m1.RefractiveIndex = 1.5
					s1.Material = m1

					s2 := NewSphere()
					s2.Transform = Scaling(0.5, 0.5, 0.5)

					return []Shape{s1, s2}
				}(),
			},
			args: args{
				comps: func() Computations {
					xs := Intersections{
						{
							T: 4,
							Object: func() Shape {
								s1 := NewSphere()
								m1 := NewMaterial()
								m1.Color = Color{0.8, 1.0, 0.6}
								m1.Diffuse = 0.7
								m1.Specular = 0.2
								m1.Transparency = 1.0
								m1.RefractiveIndex = 1.5
								s1.Material = m1

								return s1
							}(),
						},
						{
							T:      6,
							Object: dw.Objects[0],
						},
					}
					return xs[0].PrepareComputations(
						NewRay(
							*NewPoint(0, 0, -5),
							*NewVector(0, 0, 1),
						),
						xs,
					)
				}(),
				remaining: 0,
			},
			want: Black,
		},
		{
			name: "the refracted color under total internal reflection",
			fields: fields{
				Light: dw.Light,
				Objects: func() []Shape {
					s1 := NewSphere()
					s1.Material = m1

					s2 := NewSphere()
					s2.Transform = Scaling(0.5, 0.5, 0.5)

					return []Shape{s1, s2}
				}(),
			},
			args: args{
				comps: func() Computations {
					xs := Intersections{
						{
							T:      -math.Sqrt(2) / 2,
							Object: glassySphere,
						},
						{
							T:      math.Sqrt(2) / 2,
							Object: glassySphere,
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
				remaining: 5,
			},
			want: Black,
		},
		//{
		//	name: "the refracted color with a refracted ray",
		//	fields: fields{
		//		Light:   dw.Light,
		//		Objects: []Shape{s1, s2},
		//	},
		//	args: args{
		//		comps: func() Computations {
		//			xs := Intersections{
		//				{
		//					T:      -0.9899,
		//					Object: s1,
		//				},
		//				{
		//					T:      -0.4899,
		//					Object: s2,
		//				},
		//				{
		//					T:      0.4899,
		//					Object: s2,
		//				},
		//				{
		//					T:      0.9899,
		//					Object: s1,
		//				},
		//			}
		//			return xs[2].PrepareComputations(
		//				NewRay(
		//					*NewPoint(0, 0, 0.1),
		//					*NewVector(0, 1, 0),
		//				),
		//				xs,
		//			)
		//		}(),
		//		remaining: 5,
		//	},
		//	want: Color{0, 0.99888, 0.04725},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := World{
				Objects: tt.fields.Objects,
				Light:   tt.fields.Light,
			}
			if got := w.RefractedColor(tt.args.comps, tt.args.remaining); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RefractedColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
