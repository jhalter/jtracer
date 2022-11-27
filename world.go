package jtracer

import (
	"sort"
)

type World struct {
	Objects []Shaper
	Light   Light
}

func NewWorld() World {
	return World{}
}

// DefaultWorld returns a World that contains two concentric spheres, where the outermost is a unit sphere and the
// innermost has a radius of 0.5. Both lie at the origin.
func DefaultWorld() World {
	s1 := NewSphere()
	m1 := NewMaterial()
	m1.Color = Color{0.8, 1.0, 0.6}
	m1.Diffuse = 0.7
	m1.Specular = 0.2
	s1.Material = m1

	s2 := NewSphere()
	s2.Transform = Scaling(0.5, 0.5, 0.5)

	w := NewWorld()
	w.Objects = []Shaper{s1, s2}

	w.Light = NewPointLight(
		*NewPoint(-10, 10, -10),
		Color{1, 1, 1},
	)

	return w
}

func (w World) Intersect(r Ray) Intersections {
	xs := Intersections{}

	for _, object := range w.Objects {
		newxs := object.Intersects(r)
		xs = append(xs, newxs...)
	}
	sort.Sort(xs)

	return xs
}

func (w World) ColorAt(r Ray, remaining int) Color {
	xs := w.Intersect(r)
	hit := xs.Hit()
	if hit == nil {
		//spew.Dump("NOHIT")
		return Black
	}

	//spew.Dump("HIT color", hit.Object.GetMaterial().Color)
	comps := hit.PrepareComputations(r, nil)
	return w.ShadeHit(comps, remaining)
}

func (w World) ShadeHit(comps Computations, remaining int) Color {
	shadowed := w.IsShadowed(comps.OverPoint)

	surface := comps.Object.GetMaterial().Lighting(comps.Object, w.Light, comps.OverPoint, comps.Eyev, comps.Normalv, shadowed)
	reflected := w.ReflectedColor(comps, remaining)
	//
	//spew.Dump(comps.Object.GetMaterial())
	//spew.Dump(w.Light)

	return *surface.Add(&reflected)
}

func (w World) IsShadowed(p Tuple) bool {
	v := w.Light.Position.Subtract(&p)
	distance := v.Magnitude()
	direction := v.Normalize()

	r := NewRay(p, *direction)
	intersections := w.Intersect(r)

	h := intersections.Hit()

	if h != nil && h.T < distance {
		return true
	}

	return false
}

func (w World) ReflectedColor(comps Computations, remaining int) Color {
	//spew.Dump(w)
	if comps.Object.GetMaterial().Reflectivity == 0 || remaining == 0 {
		return Black
	}

	reflectRay := Ray{Origin: comps.OverPoint, Direction: comps.Reflectv}
	color := w.ColorAt(reflectRay, remaining-1)
	//
	//spew.Dump("OrigRay", NewVector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	//spew.Dump("Reflectv", comps.Reflectv)
	//spew.Dump("Color at the result of reflected ray", color)
	//spew.Dump("Reflectivity of this material", comps.Object.GetMaterial().Reflectivity)

	return *color.MultiplyByScalar(comps.Object.GetMaterial().Reflectivity)
}

func (w World) RefractedColor(comps Computations, remaining int) Color {
	if comps.Object.GetMaterial().Transparency == 0 || remaining == 0 {
		return Black
	}

	// Start check for total internal reflection
	// find the ratio of the first index of refraction to the second
	nRatio := comps.N1 / comps.N2
	// cos(theta_i) is the same as the dot product of the two vectors
	cosI := comps.Eyev.Dot(&comps.Normalv)
	sin2T := (nRatio * nRatio) * (1 - (cosI * cosI))
	if sin2T > 1 {
		return Black
	}
	// End check for total internal reflection TODO: move to function?

	return White

}
