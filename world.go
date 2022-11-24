package jtracer

import "sort"

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

func (w World) ColorAt(r Ray) Color {
	xs := w.Intersect(r)
	hit := xs.Hit()
	if hit == nil {
		return Black
	}
	comps := hit.PrepareComputations(r)

	return w.ShadeHit(comps)
}

func (w World) ShadeHit(comps Computations) Color {
	shadowed := w.IsShadowed(comps.OverPoint)

	return comps.Object.GetMaterial().Lighting(comps.Object, w.Light, comps.OverPoint, comps.Eyev, comps.Normalv, shadowed)
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
