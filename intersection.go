package jtracer

import "math"

type Intersection struct {
	T      float64
	Object Shaper
}

type Intersections []Intersection

func (i Intersections) Len() int {
	return len(i)
}

func (i Intersections) Less(a, b int) bool {
	return i[a].T < i[b].T
}

func (i Intersections) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

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

type Computations struct {
	T         float64
	Object    Shaper
	Point     Tuple
	Eyev      Tuple
	Normalv   Tuple
	Inside    bool
	OverPoint Tuple
	Reflectv  Tuple
}

func (i Intersection) PrepareComputations(r Ray) Computations {
	comps := Computations{
		T:      i.T,
		Object: i.Object,
		Inside: false,
	}

	comps.Point = *r.Position(comps.T)
	comps.Eyev = *r.Direction.Negate()
	comps.Normalv = comps.Object.NormalAt(comps.Point)

	if comps.Normalv.Dot(&comps.Eyev) < 0 {
		comps.Inside = true
		comps.Normalv = *comps.Normalv.Negate()
	}

	comps.OverPoint = *comps.Point.Add(comps.Normalv.Multiply(epsilon))
	comps.Reflectv = r.Direction.Reflect(comps.Normalv)

	return comps
}
