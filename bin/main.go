package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"jtracer"
	"log"
	"net/http"
	"os"
)
import _ "net/http/pprof"

func main() {
	height := flag.Float64("height", -1, "Height of output image")
	width := flag.Float64("width", -1, "Height of output image")
	outputFile := flag.String("out", "out.png", "Filename of output image")

	flag.Parse()

	inputFileName := os.Args[len(os.Args)-1]
	spew.Dump(inputFileName)

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	if *height != -1.0 {
		// TODO
	}

	if *width != -1.0 {
		// TODO
	}

	scene, err := jtracer.LoadSceneFile(inputFileName)
	if err != nil {
		panic(err)
	}

	canvas := scene.Camera.Render(jtracer.World{
		Objects: scene.Objects,
		Light:   scene.Light,
	})
	err = canvas.SavePNG(*outputFile)
	if err != nil {
		panic(err)
	}

}
