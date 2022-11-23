package main

import "jtracer"

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

	// for each pixel in the row
	for y := 0.0; y < canvasPixels-1.0; y++ {
		worldY := half - pixelSize*y

		// compute the world x coordinate (left = -half, right = half)
		for x := 0.0; x < canvasPixels-1.0; x++ {
			worldX := -half + pixelSize*x
			position := jtracer.NewPoint(worldX, worldY, wallZ)

			xs := shape.Intersects(jtracer.Ray{
				Origin:    *rayOrigin,
				Direction: *position.Subtract(rayOrigin).Normalize(),
			})

			hit := xs.Hit()
			if hit != nil {
				canvas.WritePixel(int(x), int(y), &color)
			}
		}
	}

	err := canvas.SavePNG("chapter5.png")
	if err != nil {
		panic(err)
	}
}
