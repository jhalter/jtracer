package jtracer

// TODO: What should I name this interface?
type Shaper interface {
	GetMaterial() Material
	GetTransform() Matrix
	GetID() int
	LocalNormalAt(Tuple) Tuple
	LocalIntersect(Ray) Intersections
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
