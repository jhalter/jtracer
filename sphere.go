package jtracer

import (
	"math"
	"math/rand"
)

type Sphere struct {
	AbstractShape
}

func NewSphere() *Sphere {
	s := &Sphere{
		AbstractShape: AbstractShape{
			ID:       rand.Int(),
			Material: NewMaterial(),
		},
	}
	s.SetTransform(IdentityMatrix)
	return s
}

// NewSphereWithID returns a Sphere with a specific ID for use in test assertions
func NewSphereWithID(id int) *Sphere {
	s := NewSphere()
	s.AbstractShape.ID = id
	return s
}

func NewGlassSphere() *Sphere {
	s := &Sphere{
		AbstractShape: AbstractShape{
			ID: rand.Int(),
			Material: Material{
				Transparency:    1.0,
				RefractiveIndex: 1.5,
			},
		},
	}
	s.SetTransform(IdentityMatrix)
	return s
}

func (s *Sphere) GetMaterial() Material {
	return s.Material
}

func (s *Sphere) LocalIntersect(ray Ray) Intersections {
	xs := Intersections{}

	sphereToRay := ray.Origin.Subtract(NewPoint(0, 0, 0))

	a := ray.Direction.Dot(&ray.Direction)
	b := 2 * ray.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := (b * b) - 4*a*c

	if discriminant < 0 {
		return xs
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	xs = append(xs, Intersection{T: t1, Object: s})
	xs = append(xs, Intersection{T: t2, Object: s})
	return xs
}

func (s *Sphere) LocalNormalAt(point Tuple) Tuple {
	return *point.Subtract(NewPoint(0, 0, 0))
}

func Intersects(s Shape, r Ray) Intersections {
	return s.LocalIntersect(r.Transform(s.GetInverse()))
}

func NormalAt(s Shape, worldPoint Tuple) Tuple {
	localPoint := s.GetInverse().MultiplyByTuple(worldPoint)
	objectNormal := s.LocalNormalAt(*localPoint)
	worldNormal := s.GetInverseTranspose().MultiplyByTuple(objectNormal)
	worldNormal.W = 0

	return *worldNormal.Normalize()
}
