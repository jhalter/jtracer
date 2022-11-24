package jtracer

import (
	"math"
)

type Material struct {
	Color        Color
	Ambient      float64
	Diffuse      float64
	Specular     float64
	Shininess    float64
	Pattern      Patterny
	HasPattern   bool
	Reflectivity float64
}

func NewMaterial() Material {
	return Material{
		Color:     Color{Red: 1, Green: 1, Blue: 1},
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}

func (m Material) Lighting(object Shaper, light Light, point, eyev, normalv Tuple, inShadow bool) Color {

	var color Color
	if m.HasPattern {
		color = m.Pattern.ColorAtObject(object, point)
	} else {
		color = m.Color
	}

	// combine the surface color with the light's color/intensity
	effectiveColor := color.Multiply(&light.Intensity)

	// find the direction to the light source
	lightv := light.Position.Subtract(&point).Normalize()

	// compute the ambient contribution
	ambient := effectiveColor.MultiplyByScalar(m.Ambient)

	// light_dot_normal represents the cosine of the angle between the
	// light vector and the normal vector. A negative number means the
	// light is on the other side of the surface.

	var diffuse, specular Color
	lightDotNormal := lightv.Dot(&normalv)
	if lightDotNormal < 0 || inShadow {
		diffuse = Black
		specular = Black
	} else {
		diffuse = *effectiveColor.MultiplyByScalar(m.Diffuse)
		diffuse = *diffuse.MultiplyByScalar(lightDotNormal)

		//  reflect_dot_eye represents the cosine of the angle between the
		//  reflection vector and the eye vector. A negative number means the
		//  light reflects away from the eye.
		reflectV := lightv.Multiply(-1).Reflect(normalv)
		reflectDotEye := reflectV.Dot(&eyev)

		if reflectDotEye <= 0 {
			specular = Black
		} else {
			// compute the specular contribution
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = light.Intensity
			specular = *specular.MultiplyByScalar(m.Specular)
			specular = *specular.MultiplyByScalar(factor)
		}
	}

	c := ambient.Add(&diffuse)
	c = c.Add(&specular)
	return *c
}
