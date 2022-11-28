package jtracer

import (
	"math"
	"math/rand"
)

type Plane struct {
	ID int
	Shape
}

func NewPlane() Plane {
	return Plane{
		Shape: Shape{
			ID:        rand.Int(),
			Transform: IdentityMatrix,
			Material:  NewMaterial(),
		},
	}
}

func (p Plane) LocalNormalAt(_ Tuple) Tuple {
	return *NewVector(0, 1, 0)
}

func (p Plane) LocalIntersect(r Ray) Intersections {
	if math.Abs(r.Direction.Y) < epsilon {
		return Intersections{}
	}

	t := -r.Origin.Y / r.Direction.Y

	return Intersections{{t, p}}
}
