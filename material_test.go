package jtracer

import (
	"github.com/davecgh/go-spew/spew"
	"math"
	"testing"
)

func TestMaterial_Lighting(t *testing.T) {
	m := NewMaterial()
	p := NewPoint(0, 0, 0)

	spew.Dump(m, p)

	type fields struct {
		Color        Color
		Ambient      float64
		Diffuse      float64
		Specular     float64
		Shininess    float64
		HasPattern   bool
		Reflectivity float64
	}
	type args struct {
		object   Sphere
		light    Light
		point    Tuple
		eyev     Tuple
		normalv  Tuple
		inShadow bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Color
	}{
		{
			name: "lighting with the eye between the light and the surface",
			fields: fields{
				Color{1, 1, 1},
				0.1,
				0.9,
				0.9,
				200.0,
				false,
				0.0,
			},
			args: args{
				object: Sphere{},
				light: Light{
					Position:  *NewPoint(0, 0, -10),
					Intensity: Color{1, 1, 1},
				},
				point:    *NewPoint(0, 0, 0),
				eyev:     *NewVector(0, 0, -1),
				normalv:  *NewVector(0, 0, -1),
				inShadow: false,
			},
			want: Color{1.9, 1.9, 1.9},
		},
		{
			name: "lighting with the eye between the light and the surface, eye offset 45 degrees",
			fields: fields{
				Color{1, 1, 1},
				0.1,
				0.9,
				0.9,
				200.0,
				false,
				0.0,
			},
			args: args{
				object: Sphere{},
				light: Light{
					Position:  *NewPoint(0, 0, -10),
					Intensity: Color{1, 1, 1},
				},
				point:    *NewPoint(0, 0, 0),
				eyev:     *NewVector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
				normalv:  *NewVector(0, 0, -1),
				inShadow: false,
			},
			want: Color{1, 1, 1},
		},
		{
			name: "lighting with the eye opposite surface, light offet 45 degrees",
			fields: fields{
				Color{1, 1, 1},
				0.1,
				0.9,
				0.9,
				200.0,
				false,
				0.0,
			},
			args: args{
				object: Sphere{},
				light: Light{
					Position:  *NewPoint(0, 10, -10),
					Intensity: Color{1, 1, 1},
				},
				point:    *NewPoint(0, 0, 0),
				eyev:     *NewVector(0, 0, -1),
				normalv:  *NewVector(0, 0, -1),
				inShadow: false,
			},
			want: Color{0.7364, 0.7364, 0.7364},
		},
		{
			name: "lighting with the eye in the path of the reflection vector",
			fields: fields{
				Color{1, 1, 1},
				0.1,
				0.9,
				0.9,
				200.0,
				false,
				0.0,
			},
			args: args{
				object: Sphere{},
				light: Light{
					Position:  *NewPoint(0, 10, -10),
					Intensity: Color{1, 1, 1},
				},
				point:    *NewPoint(0, 0, 0),
				eyev:     *NewVector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2),
				normalv:  *NewVector(0, 0, -1),
				inShadow: false,
			},
			want: Color{1.6364, 1.6364, 1.6364},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Material{
				Color:        tt.fields.Color,
				Ambient:      tt.fields.Ambient,
				Diffuse:      tt.fields.Diffuse,
				Specular:     tt.fields.Specular,
				Shininess:    tt.fields.Shininess,
				HasPattern:   tt.fields.HasPattern,
				Reflectivity: tt.fields.Reflectivity,
			}
			if got := m.Lighting(tt.args.light, tt.args.point, tt.args.eyev, tt.args.normalv, tt.args.inShadow); !got.Equals(&tt.want) {
				t.Errorf("Lighting() = %v, want %v", got, tt.want)
			}
		})
	}
}
