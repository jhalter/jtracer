package jtracer

import (
	"math"
	"math/rand"
)

type Sphere struct {
	ID        int
	Transform Matrix
	Material  Material
}

func NewSphere() Sphere {
	return Sphere{
		ID:        rand.Int(),
		Transform: IdentityMatrix,
		Material:  NewMaterial(),
	}
}

func (s Sphere) GetMaterial() Material {
	return s.Material
}

func (s Sphere) Intersects(r Ray) Intersections {
	ray := r.Transform(s.Transform.Inverse())

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

func (s Sphere) NormalAt(worldPoint Tuple) Tuple {
	objectPoint := s.Transform.Inverse().MultiplyByTuple(worldPoint)
	objectNormal := objectPoint.Subtract(NewPoint(0, 0, 0))
	worldNormal := s.Transform.Inverse().Transpose().MultiplyByTuple(*objectNormal)
	worldNormal.W = 0

	return *worldNormal.Normalize()
}
