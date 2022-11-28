package main

import (
	"jtracer"
)

func main() {
	// start the ray at z = -5
	rayOrigin := jtracer.NewPoint(0, 0, -5)

	// put the wall at z = 10
	wallZ := 10.0
	wallSize := 7.0
	canvasPixels := 600.0

	pixelSize := wallSize / canvasPixels
	half := wallSize / 2

	canvas := jtracer.NewCanvas(int(canvasPixels), int(canvasPixels))
	color := jtracer.Red
	shape := jtracer.NewSphere()
	shape.Material = jtracer.NewMaterial()
	shape.Material.Color = jtracer.Color{Red: 1, Green: 0.2, Blue: 1}

	lightPosition := jtracer.NewPoint(-10, 10, -10)
	lightColor := jtracer.Color{Red: 1, Green: 1, Blue: 1}
	light := jtracer.NewPointLight(*lightPosition, lightColor)

	// for each pixel in the row
	for y := 0.0; y < canvasPixels-1.0; y++ {
		worldY := half - pixelSize*y

		// compute the world x coordinate (left = -half, right = half)
		for x := 0.0; x < canvasPixels-1.0; x++ {
			worldX := -half + pixelSize*x

			position := jtracer.NewPoint(worldX, worldY, wallZ)

			r := jtracer.Ray{Origin: *rayOrigin, Direction: *position.Subtract(rayOrigin).Normalize()}
			xs := shape.LocalIntersect(r)

			hit := xs.Hit()
			if hit != nil {
				point := r.Position(hit.T)
				normal := shape.LocalNormalAt(*point)
				eye := r.Direction.Negate()

				color = shape.Material.Lighting(nil, light, *point, *eye, normal, false)

				canvas.WritePixel(int(x), int(y), &color)
			}
		}
	}

	err := canvas.SavePNG("chapter6.png")
	if err != nil {
		panic(err)
	}
}
