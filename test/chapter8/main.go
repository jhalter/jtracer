package main

import (
	"jtracer"
	"math"
)

func main() {
	floor := jtracer.NewSphere()
	floor.Transform = jtracer.Scaling(10, 0.01, 10)
	floor.Material = jtracer.NewMaterial()
	floor.Material.Color = jtracer.Color{Red: 1, Green: 0.9, Blue: 0.9}
	floor.Material.Specular = 0

	leftWall := jtracer.NewSphere()
	tf1 := jtracer.NewTranslation(0, 0, 5)
	tf2 := jtracer.RotationY(-math.Pi / 4)
	tf3 := jtracer.RotationX(math.Pi / 2)
	tf4 := jtracer.Scaling(10, 0.01, 10)
	leftWall.Transform = tf1.Multiply(tf2).Multiply(tf3).Multiply(tf4)
	// leftWall.Transform = tf4.Multiply(tf3).Multiply(tf2).Multiply(tf1)

	leftWall.Material = floor.Material

	rightWall := jtracer.NewSphere()
	tf1 = jtracer.NewTranslation(0, 0, 5)
	tf2 = jtracer.RotationY(math.Pi / 4)
	tf3 = jtracer.RotationX(math.Pi / 2)
	tf4 = jtracer.Scaling(10, 0.01, 10)
	rightWall.Transform = tf1.Multiply(tf2).Multiply(tf3).Multiply(tf4)
	rightWall.Material = floor.Material

	middle := jtracer.NewSphere()
	middle.Transform = jtracer.NewTranslation(-0.5, 1, 0.5)
	middle.Material = jtracer.NewMaterial()
	middle.Material.Color = jtracer.Color{Red: 0.1, Green: 1, Blue: 0.5}
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3

	right := jtracer.NewSphere()
	right.Transform = jtracer.NewTranslation(1.5, 0.5, -0.5).Multiply(jtracer.Scaling(0.5, 0.5, 0.5))
	right.Material = jtracer.NewMaterial()
	right.Material.Color = jtracer.Color{Red: 0.5, Green: 1, Blue: 0.1}
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.3

	left := jtracer.NewSphere()
	left.Transform = jtracer.NewTranslation(-1.5, 0.33, -0.75).Multiply(jtracer.Scaling(0.33, 0.33, 0.33))
	left.Material = jtracer.NewMaterial()
	left.Material.Color = jtracer.Color{Red: 1, Green: 0.8, Blue: 0.1}
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3

	world := jtracer.NewWorld()
	world.Objects = []jtracer.Shaper{floor, leftWall, rightWall, middle, right, left}
	world.Light = jtracer.NewPointLight(*jtracer.NewPoint(-10, 10, -10), jtracer.Color{Red: 1, Green: 1, Blue: 1})

	camera := jtracer.NewCamera(600, 300, math.Pi/3)
	camera.Transform = jtracer.ViewTransform(
		jtracer.NewPoint(0, 1.5, -5),
		jtracer.NewPoint(0, 1, 0),
		jtracer.NewVector(0, 1, 0),
	)

	canvas := camera.Render(world)
	canvas.SavePNG("chapter8.png")
}
