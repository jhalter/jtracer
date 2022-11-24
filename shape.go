package jtracer

type Shape interface {
	Intersects(r Ray) Intersections
	NormalAt(worldPoint Tuple) Tuple
	GetMaterial() Material
}
