package jtracer

import (
	"math"
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
	return World{
		Objects: []Shaper{
			&Sphere{
				Shape: Shape{
					Transform: IdentityMatrix,
					Material: Material{
						Color:           Color{0.8, 1.0, 0.6},
						Ambient:         0.1,
						Diffuse:         0.7,
						Specular:        0.2,
						Shininess:       200.0,
						RefractiveIndex: 1.0,
					},
				},
			},
			&Sphere{
				Shape: Shape{
					Transform: Scaling(0.5, 0.5, 0.5),
					Material:  NewMaterial(),
				},
			},
		},
		Light: NewPointLight(
			*NewPoint(-10, 10, -10),
			Color{1, 1, 1},
		),
	}
}

func (w World) Intersect(r Ray) Intersections {
	xs := Intersections{}

	for _, object := range w.Objects {
		newxs := Intersects(object, r)
		xs = append(xs, newxs...)
	}
	sort.Sort(xs)

	return xs
}

func (w World) ColorAt(r Ray, remaining int) Color {
	xs := w.Intersect(r)
	hit := xs.Hit()
	if hit == nil {
		return Black
	}

	comps := hit.PrepareComputations(r, xs)
	return w.ShadeHit(comps, remaining)
}

func (w World) ShadeHit(comps Computations, remaining int) Color {
	shadowed := w.IsShadowed(comps.OverPoint)

	surface := comps.Object.GetMaterial().Lighting(comps.Object, w.Light, comps.OverPoint, comps.Eyev, comps.Normalv, shadowed)
	reflected := w.ReflectedColor(comps, remaining)
	refracted := w.RefractedColor(comps, remaining)

	material := comps.Object.GetMaterial()
	if material.Reflectivity > 0.0 && material.Transparency > 0.0 {
		reflectance := Schlick(comps)

		// surface + reflected

		//spew.Dump(reflectance)
		//

		//return surface + reflected * reflectance + refracted * (1 - reflectance)

		m1 := reflected.MultiplyByScalar(reflectance)
		m2 := refracted.MultiplyByScalar(1 - reflectance)

		tmp1 := surface.Add(m1).Add(m2)
		//spew.Dump(reflectance, reflected, refracted, tmp1)

		return *tmp1
	}

	//
	//spew.Dump(comps.Object.GetMaterial())
	//spew.Dump(w.Light)

	return *surface.Add(&reflected).Add(&refracted)
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

	// Find cos(theta_t) via trigonometric identity
	//cos_t ← sqrt(1.0 - sin2_t)

	cosT := math.Sqrt(1.0 - sin2T)

	// Compute the direction of the refracted ray
	//direction ← comps.normalv * (n_ratio * cos_i - cos_t) -
	//             comps.eyev * n_ratio

	//foo := (nRatio * cosI) - cosT
	//foo2 := comps.Eyev.Multiply(nRatio)

	baz1 := comps.Normalv.Multiply((nRatio * cosI) - cosT)
	baz2 := comps.Eyev.Multiply(nRatio)

	direction := baz1.Subtract(baz2)

	// Create the refracted ray
	//refract_ray ← ray(comps.under_point, direction)
	refractRay := NewRay(comps.UnderPoint, *direction)

	//
	//# Find the color of the refracted ray, making sure to multiply
	//
	//# by the transparency value to account for any opacity
	//color ← color_at(world, refract_ray, remaining - 1) *
	//         comps.object.material.transparency

	color := w.ColorAt(refractRay, remaining-1)
	//spew.Dump("colorAt", comps.Object.GetID(), color)
	color = *color.MultiplyByScalar(comps.Object.GetMaterial().Transparency)

	//spew.Dump("refractedColor", color)

	return color

}
