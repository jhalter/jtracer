package jtracer

type Shape interface {
	GetMaterial() Material
	GetTransform() Matrix
	SetTransform(Matrix)
	GetInverse() Matrix
	GetInverseTranspose() Matrix
	GetID() int
	LocalNormalAt(Tuple) Tuple
	LocalIntersect(Ray) Intersections
}

type AbstractShape struct {
	ID               int
	Material         Material
	Transform        Matrix
	Inverse          Matrix
	InverseTranspose Matrix
}

func (s *AbstractShape) GetID() int {
	return s.ID
}

func (s *AbstractShape) GetMaterial() Material {
	return s.Material
}

func (s *AbstractShape) GetTransform() Matrix {
	return s.Transform
}

func (s *AbstractShape) GetInverse() Matrix {
	return s.Inverse
}

func (s *AbstractShape) GetInverseTranspose() Matrix {
	return s.InverseTranspose
}

func (s *AbstractShape) SetTransform(t Matrix) {
	s.Transform = t
	s.Inverse = t.Inverse()
	s.InverseTranspose = s.Inverse.Transpose()
}
