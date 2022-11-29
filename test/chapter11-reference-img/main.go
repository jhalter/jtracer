package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

import (
	"jtracer"
)

/*
# ======================================================
# reflect-refract.yml
#
# This file describes the scene illustrated at the start
# of chapter 11, "Reflection and Refraction", in "The
# Ray Tracer Challenge"
#
# by Jamis Buck <jamis@jamisbuck.org>
# ======================================================

# ======================================================
# the camera
# ======================================================

- add: camera
  width: 400
  height: 200
  field-of-view: 1.152
  from: [-2.6, 1.5, -3.9]
  to: [-0.6, 1, -0.8]
  up: [0, 1, 0]

# ======================================================
# light sources
# ======================================================

- add: light
  at: [-4.9, 4.9, -1]
  intensity: [1, 1, 1]

# ======================================================
# define constants to avoid duplication
# ======================================================

- define: wall-material
  value:
    pattern:
      type: stripes
      colors:
        - [0.45, 0.45, 0.45]
        - [0.55, 0.55, 0.55]
      transform:
        - [ scale, 0.25, 0.25, 0.25 ]
        - [ rotate-y, 1.5708 ]
    ambient: 0
    diffuse: 0.4
    specular: 0
    reflective: 0.3

# ======================================================
# describe the elements of the scene
# ======================================================

# the checkered floor
- add: plane
  transform:
    - [ rotate-y, 0.31415 ]
  material:
    pattern:
      type: checkers
      colors:
        - [0.35, 0.35, 0.35]
        - [0.65, 0.65, 0.65]
    specular: 0
    reflective: 0.4

# the ceiling
- add: plane
  transform:
    - [ translate, 0, 5, 0 ]
  material:
    color: [0.8, 0.8, 0.8]
    ambient: 0.3
    specular: 0

# west wall
- add: plane
  transform:
    - [ rotate-y, 1.5708 ] # orient texture
    - [ rotate-z, 1.5708 ] # rotate to vertical
    - [ translate, -5, 0, 0 ]
  material: wall-material

# east wall
- add: plane
  transform:
    - [ rotate-y, 1.5708 ] # orient texture
    - [ rotate-z, 1.5708 ] # rotate to vertical
    - [ translate, 5, 0, 0 ]
  material: wall-material

# north wall
- add: plane
  transform:
    - [ rotate-x, 1.5708 ] # rotate to vertical
    - [ translate, 0, 0, 5 ]
  material: wall-material

# south wall
- add: plane
  transform:
    - [ rotate-x, 1.5708 ] # rotate to vertical
    - [ translate, 0, 0, -5 ]
  material: wall-material

# ----------------------
# background balls
# ----------------------

- add: sphere
  transform:
    - [ scale, 0.4, 0.4, 0.4 ]
    - [ translate, 4.6, 0.4, 1 ]
  material:
    color: [0.8, 0.5, 0.3]
    shininess: 50

- add: sphere
  transform:
    - [ scale, 0.3, 0.3, 0.3 ]
    - [ translate, 4.7, 0.3, 0.4 ]
  material:
    color: [0.9, 0.4, 0.5]
    shininess: 50

- add: sphere
  transform:
    - [ scale, 0.5, 0.5, 0.5 ]
    - [ translate, -1, 0.5, 4.5 ]
  material:
    color: [0.4, 0.9, 0.6]
    shininess: 50

- add: sphere
  transform:
    - [ scale, 0.3, 0.3, 0.3 ]
    - [ translate, -1.7, 0.3, 4.7 ]
  material:
    color: [0.4, 0.6, 0.9]
    shininess: 50

# ----------------------
# foreground balls
# ----------------------

# red sphere
- add: sphere
  transform:
    - [ translate, -0.6, 1, 0.6 ]
  material:
    color: [1, 0.3, 0.2]
    specular: 0.4
    shininess: 5

# blue glass sphere
- add: sphere
  transform:
    - [ scale, 0.7, 0.7, 0.7 ]
    - [ translate, 0.6, 0.7, -0.6 ]
  material:
    color: [0, 0, 0.2]
    ambient: 0
    diffuse: 0.4
    specular: 0.9
    shininess: 300
    reflective: 0.9
    transparency: 0.9
    refractive-index: 1.5

# green glass sphere
- add: sphere
  transform:
    - [ scale, 0.5, 0.5, 0.5 ]
    - [ translate, -0.7, 0.5, -0.8 ]
  material:
    color: [0, 0.2, 0]
    ambient: 0
    diffuse: 0.4
    specular: 0.9
    shininess: 300
    reflective: 0.9
    transparency: 0.9
    refractive-index: 1.5
*/

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	wallPattern := jtracer.NewStripePattern(jtracer.Color{0.45, 0.45, 0.45}, jtracer.Color{0.55, 0.55, 0.55})
	wallPattern.SetTransform(jtracer.RotationY(1.5708).Multiply(jtracer.Scaling(0.25, 0.25, 0.25)))
	wallMaterial := jtracer.Material{
		HasPattern:      true,
		Pattern:         &wallPattern,
		Diffuse:         0.4,
		Reflectivity:    0.3,
		RefractiveIndex: 1.0,
	}

	// # the checkered floor
	floor := jtracer.NewPlane()
	floor.SetTransform(jtracer.RotationY(0.31415))
	floorPattern := jtracer.NewCheckersPattern(
		jtracer.Color{Red: 0.35, Green: 0.35, Blue: 0.35},
		jtracer.Color{Red: 0.65, Green: 0.65, Blue: 0.65},
	)
	floorPattern.SetTransform(jtracer.IdentityMatrix)
	floor.Material = jtracer.NewMaterial()
	floor.Material.Color = jtracer.Color{Red: 0.5, Green: 0.5, Blue: 0.5}
	floor.Material.Specular = 0
	floor.Material.HasPattern = true
	floor.Material.Pattern = &floorPattern
	floor.Material.Reflectivity = 0.4

	// # west wall
	westWall := jtracer.NewPlane()
	westWall.SetTransform(jtracer.NewTranslation(-5, 0, 0).
		Multiply(jtracer.RotationZ(1.5708)).
		Multiply(jtracer.RotationY(1.5708)))
	westWall.Material = wallMaterial

	// # east wall
	eastWall := jtracer.NewPlane()
	eastWall.SetTransform(jtracer.NewTranslation(5, 0, 0).
		Multiply(jtracer.RotationZ(1.5708)).
		Multiply(jtracer.RotationY(1.5708)))
	eastWall.Material = wallMaterial

	northWall := jtracer.NewPlane()
	northWall.SetTransform(jtracer.NewTranslation(0, 0, 5).
		Multiply(jtracer.RotationX(1.5708)))
	northWall.Material = wallMaterial

	southWall := jtracer.NewPlane()
	southWall.SetTransform(jtracer.NewTranslation(0, 0, -5).Multiply(jtracer.RotationX(1.5708)))
	southWall.Material = wallMaterial

	// background balls
	s1 := jtracer.NewSphere()
	s1.SetTransform(jtracer.NewTranslation(4.6, 0.4, 1).Multiply(jtracer.Scaling(0.4, 0.4, 0.4)))
	s1.Material.Color = jtracer.Color{0.8, 0.5, 0.3}
	s1.Material.Shininess = 50

	s2 := jtracer.NewSphere()
	s2.SetTransform(jtracer.NewTranslation(4.7, 0.3, 0.3).
		Multiply(jtracer.Scaling(0.3, 0.3, 0.3)),
	)
	s2.Material.Color = jtracer.Color{0.9, 0.4, 0.5}
	s2.Material.Shininess = 50

	s3 := jtracer.NewSphere()
	s3.SetTransform(jtracer.NewTranslation(-1, 0.5, 4.5).
		Multiply(jtracer.Scaling(0.5, 0.5, 0.5)))
	s3.Material.Color = jtracer.Color{0.4, 0.9, 0.6}
	s3.Material.Shininess = 50

	s4 := jtracer.NewSphere()
	s4.SetTransform(
		jtracer.NewTranslation(-1.7, 0.3, 4.7).
			Multiply(jtracer.Scaling(0.3, 0.3, 0.3)),
	)
	s4.Material.Color = jtracer.Color{0.4, 0.6, 0.9}
	s4.Material.Shininess = 50

	// foreground balls
	fg1 := jtracer.NewSphere()
	fg1.SetTransform(jtracer.NewTranslation(-0.6, 1, 0.6))
	fg1.Material.Color = jtracer.Color{1, 0.3, 0.2}
	fg1.Material.Shininess = 5
	fg1.Material.Specular = 0.4

	fg2 := jtracer.NewSphere()
	fg2.SetTransform(
		jtracer.NewTranslation(0.6, 0.7, -0.6).
			Multiply(jtracer.Scaling(0.7, 0.7, 0.7)),
	)
	fg2.Material.Color = jtracer.Color{0, 0, 0.2}
	fg2.Material.Ambient = 0
	fg2.Material.Diffuse = 0.4
	fg2.Material.Shininess = 300
	fg2.Material.Specular = 0.9
	fg2.Material.Reflectivity = 0.9
	fg2.Material.Transparency = 0.9
	fg2.Material.RefractiveIndex = 1.5

	fg3 := jtracer.NewSphere()
	fg3.SetTransform(
		jtracer.NewTranslation(-0.7, 0.5, -0.8).
			Multiply(jtracer.Scaling(0.5, 0.5, 0.5)),
	)
	fg3.Material.Color = jtracer.Color{0, 0.2, 0.0}
	fg3.Material.Ambient = 0
	fg3.Material.Diffuse = 0.4
	fg3.Material.Shininess = 300
	fg3.Material.Specular = 0.9
	fg3.Material.Reflectivity = 0.9
	fg3.Material.Transparency = 0.9
	fg3.Material.RefractiveIndex = 1.5

	// -----
	world := jtracer.World{
		Objects: []jtracer.Shaper{floor, eastWall, westWall, northWall, southWall, s1, s2, s3, s4, fg1, fg2, fg3},
		Light:   jtracer.NewPointLight(*jtracer.NewPoint(-4.9, 4.9, -1), jtracer.Color{Red: 1, Green: 1, Blue: 1}),
	}

	camera := jtracer.NewCamera(2000, 1000, 1.152)
	camera.Transform = jtracer.ViewTransform(
		jtracer.NewPoint(-2.6, 1.5, -3.9),
		jtracer.NewPoint(-0.6, 1, -0.8),
		jtracer.NewVector(0, 1, 0),
	)

	canvas := camera.Render(world)
	err := canvas.SavePNG("out.png")
	if err != nil {
		panic(err)
	}

}
