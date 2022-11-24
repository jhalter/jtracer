package jtracer

import "math"

func NewTranslation(x float64, y float64, z float64) Matrix {
	return Matrix{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	}
}

func Scaling(x float64, y float64, z float64) Matrix {
	return Matrix{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	}
}

func RotationX(rad float64) Matrix {
	return Matrix{
		{1, 0, 0, 0},
		{0, math.Cos(rad), -math.Sin(rad), 0},
		{0, math.Sin(rad), math.Cos(rad), 0},
		{0, 0, 0, 1},
	}
}

func RotationY(rad float64) Matrix {
	return Matrix{
		{math.Cos(rad), 0, math.Sin(rad), 0},
		{0, 1, 0, 0},
		{-math.Sin(rad), 0, math.Cos(rad), 0},
		{0, 0, 0, 1},
	}
}

func RotationZ(rad float64) Matrix {
	return Matrix{
		{math.Cos(rad), -math.Sin(rad), 0, 0},
		{math.Sin(rad), math.Cos(rad), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func Shearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	return Matrix{
		{1, xy, xz, 0},
		{yx, 1, yz, 0},
		{zx, zy, 1, 0},
		{0, 0, 0, 1},
	}
}

func ViewTransform(from, to, up *Tuple) Matrix {
	forward := to.Subtract(from).Normalize()
	upn := up.Normalize()
	left := forward.Cross(upn)
	trueUp := left.Cross(forward)

	orientation := Matrix{
		{left.X, left.Y, left.Z, 0},
		{trueUp.X, trueUp.Y, trueUp.Z, 0},
		{-forward.X, -forward.Y, -forward.Z, 0},
		{0, 0, 0, 1},
	}

	return orientation.Multiply(
		NewTranslation(
			-from.X,
			-from.Y,
			-from.Z,
		),
	)
}
