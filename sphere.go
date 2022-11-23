package jtracer

import (
	"math"
	"math/rand"
)

type Sphere struct {
	id        int
	Transform Matrix
}

func NewSphere() Sphere {
	return Sphere{
		id:        rand.Int(),
		Transform: IdentityMatrix,
	}
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
