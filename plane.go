package jtracer

import (
	"math"
	"math/rand"
)

type Plane struct {
	Shape
}

func NewPlane() *Plane {
	p := &Plane{
		Shape: Shape{
			ID:       rand.Int(),
			Material: NewMaterial(),
		},
	}
	p.SetTransform(IdentityMatrix)
	return p
}

func NewPlaneWithID(id int) *Plane {
	p := NewPlane()
	p.Shape.ID = id
	return p
}

func (p *Plane) LocalNormalAt(_ Tuple) Tuple {
	return *NewVector(0, 1, 0)
}

func (p *Plane) LocalIntersect(r Ray) Intersections {
	if math.Abs(r.Direction.Y) < epsilon {
		return Intersections{}
	}

	t := -r.Origin.Y / r.Direction.Y

	return Intersections{{t, p}}
}
