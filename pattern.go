package jtracer

import (
	"math"
)

type Patterny interface {
	ColorAt(Tuple) Color
	GetTransform() Matrix

	SetTransform(Matrix)
	GetInverse() Matrix
	GetInverseTranspose() Matrix
}

type AbstractPattern struct {
	Transform        Matrix
	Inverse          Matrix
	InverseTranspose Matrix
}

func (s *AbstractPattern) GetTransform() Matrix {
	return s.Transform
}

func (s *AbstractPattern) SetTransform(t Matrix) {
	s.Transform = t
	s.Inverse = t.Inverse()
	s.InverseTranspose = s.Inverse.Transpose()
}

func (s *AbstractPattern) GetInverse() Matrix {
	return s.Inverse
}

func (s *AbstractPattern) GetInverseTranspose() Matrix {
	return s.InverseTranspose
}

type StripePattern struct {
	A Color
	B Color
	AbstractPattern
}

func PatternAtShape(patterny Patterny, shape Shaper, worldPoint Tuple) Color {
	objectPoint := shape.GetInverse().MultiplyByTuple(worldPoint)
	patternPoint := patterny.GetInverse().MultiplyByTuple(objectPoint)

	return patterny.ColorAt(patternPoint)
}

func NewTestPattern() *TestPattern {
	p := &TestPattern{Transform: IdentityMatrix}
	p.SetTransform(IdentityMatrix)
	return p
}

func NewTestPatternWithTransform(t Matrix) *TestPattern {
	p := &TestPattern{Transform: IdentityMatrix}
	p.SetTransform(t)
	return p
}

type TestPattern struct {
	Transform Matrix
	AbstractPattern
}

func (p *TestPattern) ColorAt(point Tuple) Color {
	return Color{point.X, point.Y, point.Z}
}

func NewStripePattern(a Color, b Color) *StripePattern {
	p := StripePattern{A: a, B: b}
	p.SetTransform(IdentityMatrix)
	return &p
}

func (s *StripePattern) ColorAt(p Tuple) Color {
	if int(math.Floor(p.X))%2 == 0 {
		return s.A
	}

	return s.B
}

type CheckersPattern struct {
	A Color
	B Color
	AbstractPattern
}

func NewCheckersPattern(a Color, b Color) CheckersPattern {
	p := CheckersPattern{A: a, B: b}
	p.SetTransform(IdentityMatrix)
	return p
}

func (s *CheckersPattern) ColorAt(p Tuple) Color {
	if (int(math.Floor(p.X))+int(math.Floor(p.Y))+int(math.Floor(p.Z)))%2 == 0 {
		return s.A
	}

	return s.B
}
