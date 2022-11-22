package jtracer

import "math"

// Color defines a color
type Color struct {
	Red, Green, Blue float64
}

var Red = Color{1, 0, 0}
var White = Color{1, 1, 1}
var Black = Color{0, 0, 0}

// Add adds a tuple to this tuple
func (c *Color) Add(a *Color) *Color {
	return &Color{
		a.Red + c.Red,
		a.Green + c.Green,
		a.Blue + c.Blue,
	}
}

func (c *Color) Subtract(a *Color) *Color {
	return &Color{
		c.Red - a.Red,
		c.Green - a.Green,
		c.Blue - a.Blue,
	}
}

// MultiplyByScalar multiplies a tuple from this tuple
func (c *Color) MultiplyByScalar(a float64) *Color {
	return &Color{
		c.Red * a,
		c.Green * a,
		c.Blue * a,
	}
}

// Multiply multiplies a tuple from this tuple
func (c *Color) Multiply(a *Color) *Color {
	return &Color{
		c.Red * a.Red,
		c.Green * a.Green,
		c.Blue * a.Blue,
	}
}

func (c *Color) Equals(a *Color) bool {
	return floatEquals(c.Red, a.Red) &&
		floatEquals(c.Green, a.Green) &&
		floatEquals(c.Blue, a.Blue)
}

func clamp(scaledValue int) int {
	if scaledValue < 0 {
		return 0
	}
	if scaledValue > 255 {
		return 255
	}
	return scaledValue
}

// Normalize scales a color to a range of 0 to 255
func (c *Color) Normalize() (int, int, int) {
	return clamp(int(math.Round(c.Red * 255))),
		clamp(int(math.Round(c.Green * 255))),
		clamp(int(math.Round(c.Blue * 255)))
}
