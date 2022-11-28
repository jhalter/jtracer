package main

import (
	"github.com/davecgh/go-spew/spew"
	"jtracer"
	"math"
)

func main() {
	floor := jtracer.NewPlane()
	floor.Transform = jtracer.Scaling(10, 0.01, 10).Multiply(jtracer.RotationX(45))
	floor.Material = jtracer.NewMaterial()
	floor.Material.Color = jtracer.Color{Red: 0.5, Green: 0.5, Blue: 0.5}
	floor.Material.Specular = 0
	floor.Material.HasPattern = true
	floor.Material.Pattern = jtracer.CheckersPattern{A: jtracer.White, B: jtracer.Black, Transform: jtracer.Scaling(0.1, 0.1, 0.1)}
	floor.Material.Reflectivity = 0.4

	middle := jtracer.NewSphere()
	middle.Transform = jtracer.NewTranslation(-0.5, 1, 0.5)
	middle.Material = jtracer.NewMaterial()
	middle.Material.Color = jtracer.Color{Red: 0.1, Green: 1, Blue: 0.5}
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3
	middle.Material.Reflectivity = 0.7

	wall := jtracer.NewPlane()
	wall.Transform = jtracer.NewTranslation(0, 0, 5).Multiply(jtracer.RotationX(math.Pi / 2))

	spew.Dump(jtracer.NewTranslation(0, 0, 5).Multiply(jtracer.RotationX(math.Pi / 2)))
	spew.Dump(jtracer.IdentityMatrix.Multiply(jtracer.RotationX(math.Pi / 2)).Multiply(jtracer.NewTranslation(0, 0, 5)))

	//wall.Transform = jtracer.NewTranslation(0, 0, 3).Multiply(jtracer.RotationX(math.Pi / 2))

	wall.Material.Color = jtracer.Color{
		Red:   0.3,
		Green: 0.7,
		Blue:  0.4,
	}

	//right := jtracer.NewSphere()
	//right.Transform = jtracer.NewTranslation(1.5, 0.5, -0.5).Multiply(jtracer.Scaling(0.5, 0.5, 0.5))
	//right.Material = jtracer.NewMaterial()
	//right.Material.Color = jtracer.Color{Red: 0.5, Green: 1, Blue: 0.1}
	//right.Material.Diffuse = 0.7
	//right.Material.Specular = 0.3
	//right.Material.Reflectivity = 0
	//right.Material.Transparency = 2.9
	//right.Material.RefractiveIndex = 1.0

	right := jtracer.Sphere{
		Shape: jtracer.Shape{
			ID:        2,
			Transform: jtracer.NewTranslation(1.5, 0.5, -0.5).Multiply(jtracer.Scaling(0.5, 0.5, 0.5)),
			Material: jtracer.Material{
				Color:           jtracer.Color{0.5, 0.9, 0.4},
				Ambient:         0.1,
				Diffuse:         0.7,
				Specular:        0.3,
				Shininess:       0.0,
				RefractiveIndex: 1.3,
				Transparency:    0.0,
			},
		},
	}

	left := jtracer.NewSphere()
	left.Transform = jtracer.NewTranslation(-1.5, 0.33, -0.75).Multiply(jtracer.Scaling(0.33, 0.33, 0.33))
	left.Material = jtracer.NewMaterial()
	left.Material.Color = jtracer.Color{Red: 1, Green: 0.8, Blue: 0.1}
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3
	left.Material.Reflectivity = 0.1

	world := jtracer.NewWorld()
	world.Objects = []jtracer.Shaper{floor, wall, middle, right, left}
	world.Light = jtracer.NewPointLight(*jtracer.NewPoint(-10, 10, -10), jtracer.Color{Red: 1, Green: 1, Blue: 1})

	camera := jtracer.NewCamera(300, 150, math.Pi/3)
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
