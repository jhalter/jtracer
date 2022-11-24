package jtracer

type Light struct {
	Position  Tuple
	Intensity Color
}

func NewPointLight(p Tuple, i Color) Light {
	return Light{Position: p, Intensity: i}
}
