package jtracer

import (
	"math"
	"math/rand"
)

type Sphere struct {
	ID int
	Shape
}

func NewSphere() Sphere {
	return Sphere{
		Shape: Shape{
			ID:        rand.Int(),
			Transform: IdentityMatrix,
			Material:  NewMaterial(),
		},
	}
}

func NewGlassSphere() Sphere {
	return Sphere{
		Shape: Shape{
			ID:        rand.Int(),
			Transform: IdentityMatrix,
			Material: Material{
				Transparency:    1.0,
				RefractiveIndex: 1.5,
			},
		},
	}
}

func (s Sphere) GetMaterial() Material {
	return s.Material
}

func (s Sphere) LocalIntersect(ray Ray) Intersections {
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

func (s Sphere) LocalNormalAt(point Tuple) Tuple {
	objectNormal := point.Subtract(NewPoint(0, 0, 0))

	return *objectNormal
}

//func (s Sphere) NormalAt(worldPoint Tuple) Tuple {
//	objectPoint := s.Transform.Inverse().MultiplyByTuple(worldPoint)
//	objectNormal := objectPoint.Subtract(NewPoint(0, 0, 0))
//	worldNormal := s.Transform.Inverse().Transpose().MultiplyByTuple(*objectNormal)
//	worldNormal.W = 0
//
//	return *worldNormal.Normalize()
//}

func Intersects(s Shaper, r Ray) Intersections {
	localRay := r.Transform(s.GetTransform().Inverse())
	return s.LocalIntersect(localRay)
}

func NormalAt(s Shaper, worldPoint Tuple) Tuple {
	localPoint := s.GetTransform().Inverse().MultiplyByTuple(worldPoint)
	objectNormal := s.LocalNormalAt(localPoint)
	worldNormal := s.GetTransform().Inverse().Transpose().MultiplyByTuple(objectNormal)
	worldNormal.W = 0

	return *worldNormal.Normalize()
}
