package jtracer

// TODO: What should I name this interface?
type Shaper interface {
	Intersects(r Ray) Intersections
	NormalAt(worldPoint Tuple) Tuple
	GetMaterial() Material
	GetTransform() Matrix
	GetID() int
}

type Shape struct {
	ID        int
	Material  Material
	Transform Matrix
}

func (s Shape) GetID() int {
	return s.ID
}

func (s Shape) GetMaterial() Material {
	return s.Material
}

func (s Shape) GetTransform() Matrix {
	return s.Transform
}
