package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

import (
	"jtracer"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	scene, err := jtracer.LoadSceneFile("test.yaml")
	if err != nil {
		panic(err)
	}

	canvas := scene.Camera.Render(jtracer.World{
		Objects: scene.Objects,
		Light:   scene.Light,
	})
	err = canvas.SavePNG("out.png")
	if err != nil {
		panic(err)
	}
}
