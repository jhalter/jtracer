package jtracer

// A quick and dirty scene yaml parser

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Scene struct {
	Camera  Camera
	Light   Light
	Objects []Shaper
}

func LoadSceneFile(path string) (*Scene, error) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var test []map[string]interface{}
	err = yaml.Unmarshal(yamlFile, &test)
	if err != nil {
		return nil, err
	}

	var scene Scene

	defines := make(map[string]interface{})

	// first look for defines
	for _, k := range test {
		if k["define"] != nil {
			defines[k["define"].(string)] = k["value"]
		}
	}

	for _, k := range test {
		switch k["add"] {
		case "camera":
			scene.Camera = NewCamera(
				float64(k["width"].(int)),
				float64(k["height"].(int)),
				k["field-of-view"].(float64),
			)

			from := ConvertToFloat64(k["from"].([]interface{}))
			up := ConvertToFloat64(k["up"].([]interface{}))
			to := ConvertToFloat64(k["to"].([]interface{}))
			scene.Camera.Transform = ViewTransform(
				NewPoint(from[0], from[1], from[2]),
				NewPoint(to[0], to[1], to[2]),
				NewVector(up[0], up[1], up[2]),
			)
		case "light":
			at := ConvertToFloat64(k["at"].([]interface{}))
			intensity := ConvertToFloat64(k["intensity"].([]interface{}))
			scene.Light = NewPointLight(
				*NewPoint(at[0], at[1], at[2]),
				Color{Red: intensity[0], Green: intensity[1], Blue: intensity[2]},
			)
		case "plane":
			p := NewPlane()
			if k["transform"] != nil {
				tf := ParseTransforms(k["transform"].([]interface{}))
				p.SetTransform(tf)
			}

			if k["material"] != nil {
				if _, ok := k["material"].(string); ok {
					k["material"] = defines[k["material"].(string)]
				}

				p.Material = ParseMaterial(p.Material, k["material"].(map[string]interface{}))
			}

			scene.Objects = append(scene.Objects, p)
		case "sphere":
			s := NewSphere()
			if k["transform"] != nil {
				tf := ParseTransforms(k["transform"].([]interface{}))
				s.SetTransform(tf)
			}

			if k["material"] != nil {
				s.Material = ParseMaterial(s.Material, k["material"].(map[string]interface{}))
			}

			scene.Objects = append(scene.Objects, s)
		default:
			fmt.Printf("unknown type %v\n", k["add"])
		}
	}

	return &scene, nil
}

func ParseMaterial(m Material, cfg map[string]interface{}) Material {
	for k, v := range cfg {
		switch k {
		case "color":
			rgb := ConvertToFloat64(v.([]interface{}))
			m.Color = Color{rgb[0], rgb[1], rgb[2]}
		case "shininess":
			f := ConvertToFloat64([]interface{}{v})
			m.Shininess = f[0]
		case "specular":
			f := ConvertToFloat64([]interface{}{v})
			m.Specular = f[0]
		case "ambient":
			f := ConvertToFloat64([]interface{}{v})
			m.Ambient = f[0]
		case "diffuse":
			f := ConvertToFloat64([]interface{}{v})
			m.Diffuse = f[0]
		case "reflective":
			f := ConvertToFloat64([]interface{}{v})
			m.Reflectivity = f[0]
		case "transparency":
			f := ConvertToFloat64([]interface{}{v})
			m.Transparency = f[0]
		case "refractive-index":
			f := ConvertToFloat64([]interface{}{v})
			m.RefractiveIndex = f[0]
		case "pattern":
			pDef := v.(map[string]interface{})

			switch pDef["type"] {
			case "stripes":
				colors := pDef["colors"].([]interface{})
				a := ConvertToFloat64(colors[0].([]interface{}))
				b := ConvertToFloat64(colors[1].([]interface{}))
				p := NewStripePattern(
					Color{a[0], a[1], a[2]},
					Color{b[0], b[1], b[2]},
				)
				if pDef["transform"] != nil {
					tf := ParseTransforms(pDef["transform"].([]interface{}))
					p.SetTransform(tf)
				}
				m.HasPattern = true
				m.Pattern = &p

			case "checkers":
				colors := pDef["colors"].([]interface{})
				a := ConvertToFloat64(colors[0].([]interface{}))
				b := ConvertToFloat64(colors[1].([]interface{}))
				p := NewCheckersPattern(
					Color{a[0], a[1], a[2]},
					Color{b[0], b[1], b[2]},
				)
				if pDef["transform"] != nil {
					tf := ParseTransforms(pDef["transform"].([]interface{}))
					p.SetTransform(tf)
				}
				m.HasPattern = true
				m.Pattern = &p

			}

			//m.Pattern = f[0]
		}
	}

	return m
}

func ParseTransforms(transforms []interface{}) Matrix {
	result := IdentityMatrix

	// Apply transforms in reverse order
	for i := len(transforms) - 1; i >= 0; i-- {
		transform := transforms[i].([]interface{})
		switch transform[0] {
		case "translate":
			xyz := ConvertToFloat64(transform[1:])
			result = result.Multiply(NewTranslation(xyz[0], xyz[1], xyz[2]))
		case "rotate-x":
			result = result.Multiply(RotationX(transform[1].(float64)))
		case "rotate-y":
			result = result.Multiply(RotationY(transform[1].(float64)))
		case "rotate-z":
			result = result.Multiply(RotationZ(transform[1].(float64)))
		case "scale":
			f := ConvertToFloat64(transform[1:])
			result = result.Multiply(
				Scaling(f[0], f[1], f[2]),
			)
		default:
			fmt.Println("Unknown transform: " + transform[0].(string))
		}

	}

	return result
}

func ConvertToFloat64(values []interface{}) (results []float64) {
	for _, v := range values {
		if _, ok := v.(int); ok {
			results = append(results, float64(v.(int)))
		} else {
			results = append(results, v.(float64))
		}
	}
	return results
}
