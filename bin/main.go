package main

import (
	"flag"
	"jtracer"
	"log"
	"math"
	"net/http"
)
import _ "net/http/pprof"

func main() {
	height := flag.Float64("height", 600.0, "Height of output image")
	width := flag.Float64("width", 600, "Height of output image")
	outputFile := flag.String("out", "out.png", "Filename of output image")

	flag.Parse()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	
	world := jtracer.DefaultWorld()
	camera := jtracer.NewCamera(*height, *width, math.Pi/3)
	camera.Transform = jtracer.ViewTransform(
		jtracer.NewPoint(0, 1.5, -5),
		jtracer.NewPoint(0, 1, 0),
		jtracer.NewVector(0, 1, 0),
	)

	canvas := camera.Render(world)
	err := canvas.SavePNG(*outputFile)
	if err != nil {
		panic(err)
	}

}
