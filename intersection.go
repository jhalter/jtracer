package jtracer

import "math"

type Intersection struct {
	T      float64
	Object Shape
}

type Intersections []Intersection

func (i Intersections) Hit() *Intersection {
	lowestNonNegative := Intersection{T: math.MaxFloat64}
	for _, intersection := range i {
		if intersection.T > 0 && intersection.T < lowestNonNegative.T {
			lowestNonNegative = intersection
		}
	}

	if lowestNonNegative.T < math.MaxFloat64 {
		return &lowestNonNegative
	}

	return nil
}
