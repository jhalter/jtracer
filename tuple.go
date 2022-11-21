package jtracer

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

// Subtract substracts a tuple from this tuple
func (t *Tuple) Subtract(a *Tuple) *Tuple {
	return &Tuple{
		t.X - a.X,
		t.Y - a.Y,
		t.Z - a.Z,
		t.W - a.W,
	}
}
