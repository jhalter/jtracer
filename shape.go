package jtracer

// TODO: What should I name this interface?
type Shaper interface {
	Intersects(r Ray) Intersections
	NormalAt(worldPoint Tuple) Tuple
	GetMaterial() Material
	GetTransform() Matrix
}

type Shape struct {
	Material  Material
	Transform Matrix
}

func (s Shape) GetMaterial() Material {
	return s.Material
}

func (s Shape) GetTransform() Matrix {
	return s.Transform
}
