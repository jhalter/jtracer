package jtracer

import "math"

type Patterny interface {
	ColorAt(Tuple) Color
	ColorAtObject(Shaper, Tuple) Color
}

type StripePattern struct {
	A         Color
	B         Color
	Transform Matrix
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

func (s StripePattern) ColorAtObject(object Shaper, worldPoint Tuple) Color {
	objectPoint := object.GetTransform().Inverse().MultiplyByTuple(worldPoint)
	patternPoint := s.Transform.Inverse().MultiplyByTuple(objectPoint)

	return s.ColorAt(patternPoint)
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

func (s CheckersPattern) ColorAtObject(object Shaper, worldPoint Tuple) Color {
	objectPoint := object.GetTransform().Inverse().MultiplyByTuple(worldPoint)
	patternPoint := s.Transform.Inverse().MultiplyByTuple(objectPoint)

	return s.ColorAt(patternPoint)
}
