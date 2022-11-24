package jtracer

import "math"

type Plane struct {
	ID int
	Shape
}

func NewPlane() Plane {
	return Plane{}
}

func (p Plane) NormalAt(_ Tuple) Tuple {
	return *NewVector(0, 1, 0)
}

func (p Plane) Intersects(r Ray) Intersections {
	if math.Abs(r.Direction.Y) < epsilon {
		return Intersections{}
	}

	t := -r.Origin.Y / r.Direction.Y

	return Intersections{{t, p}}
}
