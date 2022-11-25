package main

import (
	"jtracer"
	"math"
)

func main() {
	floor := jtracer.NewPlane()
	floor.Transform = jtracer.Scaling(10, 0.01, 10)
	floor.Material = jtracer.NewMaterial()
	floor.Material.Color = jtracer.Color{Red: 0.5, Green: 0.5, Blue: 0.5}
	floor.Material.Specular = 0
	floor.Material.HasPattern = true
	floor.Material.Pattern = jtracer.CheckersPattern{A: jtracer.White, B: jtracer.Black, Transform: jtracer.Scaling(0.1, 0.1, 0.1)}
	floor.Material.Reflectivity = 0.9

	middle := jtracer.NewSphere()
	middle.Transform = jtracer.NewTranslation(-0.5, 1, 0.5)
	middle.Material = jtracer.NewMaterial()
	middle.Material.Color = jtracer.Color{Red: 0.1, Green: 1, Blue: 0.5}
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3
	middle.Material.HasPattern = true
	middle.Material.Pattern = jtracer.StripePattern{A: jtracer.White, B: jtracer.Black, Transform: jtracer.Scaling(0.05, 0.05, 0.05).Multiply(jtracer.RotationY(45))}

	right := jtracer.NewSphere()
	right.Transform = jtracer.NewTranslation(1.5, 0.5, -0.5).Multiply(jtracer.Scaling(0.5, 0.5, 0.5))
	right.Material = jtracer.NewMaterial()
	right.Material.Color = jtracer.Color{Red: 0.5, Green: 1, Blue: 0.1}
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.3
	//right.Material.Reflectivity = 0.8

	left := jtracer.NewSphere()
	left.Transform = jtracer.NewTranslation(-1.5, 0.33, -0.75).Multiply(jtracer.Scaling(0.33, 0.33, 0.33))
	left.Material = jtracer.NewMaterial()
	left.Material.Color = jtracer.Color{Red: 1, Green: 0.8, Blue: 0.1}
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3

	world := jtracer.NewWorld()
	world.Objects = []jtracer.Shaper{floor, middle, right, left}
	world.Light = jtracer.NewPointLight(*jtracer.NewPoint(-10, 10, -10), jtracer.Color{Red: 1, Green: 1, Blue: 1})

	camera := jtracer.NewCamera(600, 300, math.Pi/3)
	camera.Transform = jtracer.ViewTransform(
		jtracer.NewPoint(0, 1.5, -5),
		jtracer.NewPoint(0, 1, 0),
		jtracer.NewVector(0, 1, 0),
	)

	canvas := camera.Render(world)
	err := canvas.SavePNG("chapter11.png")
	if err != nil {
		panic(err)
	}
}
