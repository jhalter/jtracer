package jtracer

import (
	"math"
)

// Tuple describes a point in 3 dimensional space
type Tuple struct {
	X, Y, Z float64
	W       float64
}

// NewPoint creates a new tuple which is a point
func NewPoint(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}

// NewVector creates a new tuple which is a vector
func NewVector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0.0}
}

const epsilon = 0.00001

func floatEquals(a, b float64) bool {
	return (a-b) < epsilon && (b-a) < epsilon
}

// Equals checks if this tuple is mostly equal to t
func (t *Tuple) Equals(a *Tuple) bool {
	return floatEquals(a.X, t.X) &&
		floatEquals(a.Y, t.Y) &&
		floatEquals(a.Z, t.Z) &&
		floatEquals(a.W, t.W)
}

// Add adds a tuple to this tuple
func (t *Tuple) Add(a *Tuple) *Tuple {
	return &Tuple{
		t.X + a.X,
		t.Y + a.Y,
		t.Z + a.Z,
		t.W + a.W,
	}
}

// Subtract subtracts a tuple from this tuple
func (t *Tuple) Subtract(a *Tuple) *Tuple {
	return &Tuple{
		t.X - a.X,
		t.Y - a.Y,
		t.Z - a.Z,
		t.W - a.W,
	}
}

// Negate negates this tuple, subtracting it from the zero tuple
func (t *Tuple) Negate() *Tuple {
	return &Tuple{
		0 - t.X,
		0 - t.Y,
		0 - t.Z,
		0 - t.W,
	}
}

// Multiply multiplies a tuple from this tuple
func (t *Tuple) Multiply(a float64) *Tuple {
	return &Tuple{
		t.X * a,
		t.Y * a,
		t.Z * a,
		t.W * a,
	}
}

// Divide multiplies a tuple from this tuple
func (t *Tuple) Divide(a float64) *Tuple {
	return &Tuple{
		t.X / a,
		t.Y / a,
		t.Z / a,
		t.W / a,
	}
}

// Magnitude calculates the magnitude of the vector described by t
func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

// Normalize normalizes a vector
func (t *Tuple) Normalize() *Tuple {
	mag := t.Magnitude()
	return &Tuple{
		t.X / mag,
		t.Y / mag,
		t.Z / mag,
		t.W / mag,
	}
}

// Dot calculates the dot product with another tuple
func (t *Tuple) Dot(a *Tuple) float64 {
	return a.X*t.X + a.Y*t.Y + a.Z*t.Z + a.W*t.W
}

func (t *Tuple) Cross(b *Tuple) *Tuple {
	return NewVector(
		t.Y*b.Z-t.Z*b.Y,
		t.Z*b.X-t.X*b.Z,
		t.X*b.Y-t.Y*b.X)
}

func (t *Tuple) Reflect(normal Tuple) Tuple {
	return *t.Subtract(normal.Multiply(2.0).Multiply(t.Dot(&normal)))
}
