package jtracer

import (
	"github.com/davecgh/go-spew/spew"
	"math"
)

type Intersection struct {
	T      float64
	Object Shape
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
	T          float64
	Object     Shape
	Point      Tuple
	Eyev       Tuple
	Normalv    Tuple
	Inside     bool
	OverPoint  Tuple
	UnderPoint Tuple
	Reflectv   Tuple
	N1         float64 // n1 is the refractive index belonging to the material being exited
	N2         float64 // n2 is the refractive index belonging to the material being entered
}

type container []Shape

func (c container) contains(shape Shape) bool {
	for _, i := range c {
		if i.GetID() == shape.GetID() {
			return true
		}
	}
	return false
}

func (c container) remove(shape Shape) container {
	newContainer := container{}

	for _, i := range c {
		if i.GetID() != shape.GetID() {
			newContainer = append(newContainer, i)
		}
	}

	return newContainer
}

func (i Intersection) Equal(i2 Intersection) bool {
	return i.T == i2.T && i.Object.GetID() == i2.Object.GetID()
}

func (i Intersection) PrepareComputations(r Ray, xs Intersections) Computations {
	comps := Computations{
		T:      i.T,
		Object: i.Object,
		Inside: false,
	}

	// containers will record which objects have been entered but not yet exited
	// these objects must contain the subsequent intersection
	var containers container
	for _, intersection := range xs {

		if i.Equal(intersection) {
			if len(containers) == 0 {
				comps.N1 = 1.0
			} else {
				comps.N1 = containers[len(containers)-1].GetMaterial().RefractiveIndex
			}
		}

		if containers.contains(intersection.Object) {
			containers = containers.remove(intersection.Object)
		} else {
			containers = append(containers, intersection.Object)
		}

		if i.Equal(intersection) {
			if len(containers) == 0 {
				comps.N2 = 1.0
			} else {
				comps.N2 = containers[len(containers)-1].GetMaterial().RefractiveIndex
			}
		}
	}

	comps.Point = *r.Position(comps.T)
	comps.Eyev = *r.Direction.Negate()
	comps.Normalv = NormalAt(comps.Object, comps.Point)

	if comps.Normalv.Dot(&comps.Eyev) < 0 {
		comps.Inside = true
		comps.Normalv = *comps.Normalv.Negate()
	}

	comps.OverPoint = *comps.Point.Add(comps.Normalv.Multiply(epsilon))
	comps.UnderPoint = *comps.Point.Subtract(comps.Normalv.Multiply(epsilon))

	comps.Reflectv = r.Direction.Reflect(comps.Normalv)

	return comps
}

func Schlick(comps Computations) float64 {
	// find the cosine of the angle between the eye and normal vectors
	cos := comps.Eyev.Dot(&comps.Normalv)

	// total internal reflection can only occur if n1 > n2
	if comps.N1 > comps.N2 {
		n := comps.N1 / comps.N2
		sin2T := (n * n) * (1.0 - (cos * cos))
		if sin2T > 1.0 {
			return 1.0
		}

		// compute cosine of theta t using trig identity
		cosT := math.Sqrt(1.0 - sin2T)

		// when n1 > n2, use cos(theta t) instead
		cos = cosT
	}

	temp := (comps.N1 - comps.N2) / (comps.N1 + comps.N2)
	r0 := temp * temp
	tmp2 := r0 + (1-r0)*math.Pow(1-cos, 5)
	if math.IsNaN(tmp2) {
		//return 0.0
		spew.Dump(tmp2, comps)
	}
	return tmp2
}
