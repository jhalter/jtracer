package jtracer

// TODO: What should I name this interface?
type Shaper interface {
	GetMaterial() Material
	GetTransform() Matrix
	SetTransform(Matrix)
	GetInverse() Matrix
	GetInverseTranspose() Matrix

	GetID() int
	LocalNormalAt(Tuple) Tuple
	LocalIntersect(Ray) Intersections
}

type Shape struct {
	ID               int
	Material         Material
	Transform        Matrix
	Inverse          Matrix
	InverseTranspose Matrix
}

func (s *Shape) GetID() int {
	return s.ID
}

func (s *Shape) GetMaterial() Material {
	return s.Material
}

func (s *Shape) GetTransform() Matrix {
	return s.Transform
}

func (s *Shape) GetInverse() Matrix {
	return s.Inverse
}

func (s *Shape) GetInverseTranspose() Matrix {
	return s.InverseTranspose
}

func (s *Shape) SetTransform(t Matrix) {
	s.Transform = t
	s.Inverse = t.Inverse()
	s.InverseTranspose = s.Inverse.Transpose()
}
