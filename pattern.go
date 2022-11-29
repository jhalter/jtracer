package jtracer

import (
	"math"
)

type Patterny interface {
	ColorAt(Tuple) Color
	GetTransform() Matrix
}

type StripePattern struct {
	A         Color
	B         Color
	Transform Matrix
}

func PatternAtShape(patterny Patterny, shape Shaper, worldPoint Tuple) Color {
	objectPoint := shape.GetTransform().Inverse().MultiplyByTuple(worldPoint)
	patternPoint := patterny.GetTransform().Inverse().MultiplyByTuple(objectPoint)

	return patterny.ColorAt(patternPoint)
}

func NewTestPattern() TestPattern {
	return TestPattern{Transform: IdentityMatrix}
}

type TestPattern struct {
	Transform Matrix
}

func (p TestPattern) ColorAt(point Tuple) Color {
	return Color{point.X, point.Y, point.Z}
}

func (p TestPattern) GetTransform() Matrix {
	return p.Transform
}

func NewStripePattern(a Color, b Color) StripePattern {
	return StripePattern{A: a, B: b, Transform: IdentityMatrix}
}

func (s StripePattern) ColorAt(p Tuple) Color {
	if int(math.Floor(p.X))%2 == 0 {
		return s.A
	}

	return s.B
}
func (s StripePattern) GetTransform() Matrix {
	return s.Transform
}

type CheckersPattern struct {
	A         Color
	B         Color
	Transform Matrix
}

func NewCheckersPattern(a Color, b Color) CheckersPattern {
	return CheckersPattern{A: a, B: b, Transform: IdentityMatrix}
}

func (s CheckersPattern) ColorAt(p Tuple) Color {
	if (int(math.Floor(p.X))+int(math.Floor(p.Y))+int(math.Floor(p.Z)))%2 == 0 {
		return s.A
	}

	return s.B
}
func (s CheckersPattern) GetTransform() Matrix {
	return s.Transform
}
