package main

import (
	"jtracer"
)

/*
# ======================================================
# refraction-bend.yml
#
# This file describes the scene illustrated at the start
# of the "Transparency and Refraction" section, in
# chapter 11 ("Reflection and Refraction") of
# "The Ray Tracer Challenge"
#
# by Jamis Buck <jamis@jamisbuck.org>
# ======================================================

# ======================================================
# the camera
# ======================================================

- add: camera
  width: 200
  height: 200
  field-of-view: 0.5
  from: [-4.5, 0.85, -4]
  to: [0, 0.85, 0]
  up: [0, 1, 0]


# ======================================================
# light sources
# ======================================================

- add: light
  at: [-4.9, 4.9, 1]
  intensity: [1, 1, 1]

# ======================================================
# define some constants to avoid duplication
# ======================================================

- define: wall
  value:
    pattern:
      type: checkers
      colors:
        - [ 0.00, 0.00, 0.00 ]
        - [ 0.75, 0.75, 0.75 ]
      transform:
        - [ scale, 0.5, 0.5, 0.5 ]
    specular: 0

# ======================================================
# describe the scene
# ======================================================

# floor
- add: plane
  transform:
    - [ rotate-y, 0.31415 ]
  material:
    pattern:
      type: checkers
      colors:
        - [0.00, 0.00, 0.00]
        - [0.75, 0.75, 0.75]
    ambient: 0.5
    diffuse: 0.4
    specular: 0.8
    reflective: 0.1

# ceiling
- add: plane
  transform:
    - [ translate, 0, 5, 0 ]
  material:
    pattern:
      type: checkers
      colors:
        - [0.85, 0.85, 0.85]
        - [1.00, 1.00, 1.00]
      transform:
        - [ scale, 0.2, 0.2, 0.2 ]
    ambient: 0.5
    specular: 0

# west wall
- add: plane
  transform:
    - [ rotate-y, 1.5708 ] # orient texture
    - [ rotate-z, 1.5708 ] # rotate to vertical
    - [ translate, -5, 0, 0 ]
  material: wall

# east wall
- add: plane
  transform:
    - [ rotate-y, 1.5708 ] # orient texture
    - [ rotate-z, 1.5708 ] # rotate to vertical
    - [ translate, 5, 0, 0 ]
  material: wall

# north wall
- add: plane
  transform:
    - [ rotate-x, 1.5708 ] # rotate to vertical
    - [ translate, 0, 0, 5 ]
  material: wall

# south wall
- add: plane
  transform:
    - [ rotate-x, 1.5708 ] # rotate to vertical
    - [ translate, 0, 0, -5 ]
  material: wall

# background ball
- add: sphere
  transform:
    - [ translate, 4, 1, 4 ]
  material:
    color: [ 0.8, 0.1, 0.3 ]
    specular: 0

# background ball
- add: sphere
  transform:
    - [ scale, 0.4, 0.4, 0.4 ]
    - [ translate, 4.6, 0.4, 2.9 ]
  material:
    color: [ 0.1, 0.8, 0.2 ]
    shininess: 200

# background ball
- add: sphere
  transform:
    - [ scale, 0.6, 0.6, 0.6 ]
    - [ translate, 2.6, 0.6, 4.4 ]
  material:
    color: [ 0.2, 0.1, 0.8 ]
    shininess: 10
    specular: 0.4

# glass ball
- add: sphere
  transform:
    - [ scale, 1, 1, 1 ]
    - [ translate, 0.25, 1, 0 ]
  material:
    color: [ 0.8, 0.8, 0.9 ]
    ambient: 0
    diffuse: 0.2
    specular: 0.9
    shininess: 300
    transparency: 0.8
    refractive-index: 1.57
*/

func main() {

	wallMaterial := jtracer.Material{
		HasPattern: true,
		Pattern: jtracer.CheckersPattern{
			A:         jtracer.Black,
			B:         jtracer.Color{0.75, 0.75, 0.75},
			Transform: jtracer.Scaling(0.5, 0.5, 0.5),
		},
		Specular:        0,
		Ambient:         0.1,
		Diffuse:         0.9,
		Shininess:       200.0,
		RefractiveIndex: 1.0,
	}

	floor := jtracer.NewPlane()
	floor.Shape.Transform = jtracer.RotationY(0.31415)
	floor.Material.Pattern = jtracer.NewCheckersPattern(
		jtracer.Black,
		jtracer.Color{0.75, 0.75, 0.75},
	)
	floor.Material.HasPattern = true
	floor.Material.Ambient = 0.5
	floor.Material.Diffuse = 0.4
	floor.Material.Specular = 0.8
	floor.Material.Reflectivity = 0.1

	ceiling := jtracer.NewPlane()
	ceiling.Shape.Transform = jtracer.RotationY(0.31415)
	ceiling.Material.Pattern = jtracer.CheckersPattern{
		A:         jtracer.Color{0.85, 0.85, 0.85},
		B:         jtracer.White,
		Transform: jtracer.Scaling(0.2, 0.2, 0.2),
	}
	ceiling.Material.Ambient = 0.5
	ceiling.Material.Specular = 0

	westWall := jtracer.NewPlane()
	westWall.Shape.Transform = jtracer.NewTranslation(-5, 0, 0).
		Multiply(jtracer.RotationZ(1.5708)).
		Multiply(jtracer.RotationY(1.5708))
	westWall.Material = wallMaterial

	eastWall := jtracer.NewPlane()
	eastWall.Shape.Transform = jtracer.NewTranslation(5, 0, 0).
		Multiply(jtracer.RotationZ(1.5708)).
		Multiply(jtracer.RotationY(1.5708))
	eastWall.Material = wallMaterial

	northWall := jtracer.NewPlane()
	northWall.Shape.Transform = jtracer.NewTranslation(0, 0, 5).
		Multiply(jtracer.RotationX(1.5708))
	northWall.Material = wallMaterial

	southWall := jtracer.NewPlane()
	southWall.Shape.Transform = jtracer.NewTranslation(0, 0, -5).
		Multiply(jtracer.RotationX(1.5708))
	southWall.Material = wallMaterial

	s0 := jtracer.NewSphere()
	s0.Transform = jtracer.NewTranslation(4, 1, 4)
	s0.Material.Color = jtracer.Color{0.8, 0.1, 0.3}
	s0.Material.Specular = 0

	s1 := jtracer.NewSphere()
	s1.Transform = jtracer.NewTranslation(4.6, 0.4, 2.9).
		Multiply(jtracer.Scaling(0.4, 0.4, 0.4))
	s1.Material.Color = jtracer.Color{0.1, 0.8, 0.2}
	s1.Material.Shininess = 200

	s2 := jtracer.NewSphere()
	s2.Transform = jtracer.NewTranslation(2.6, 0.6, 4.4).
		Multiply(jtracer.Scaling(0.6, 0.6, 0.6))
	s2.Material.Color = jtracer.Color{0.2, 0.1, 0.8}
	s2.Material.Shininess = 10
	s2.Material.Specular = 0.4

	s3 := jtracer.NewSphere()
	s3.Transform = jtracer.NewTranslation(0.25, 1, 0).
		Multiply(jtracer.Scaling(1, 1, 1))
	s3.Material.Color = jtracer.Color{0.8, 0.8, 0.9}
	s3.Material.Ambient = 0
	s3.Material.Diffuse = 0.2
	s3.Material.Shininess = 300
	s3.Material.Specular = 0.9
	s3.Material.Transparency = 0.8
	s3.Material.RefractiveIndex = 1.57

	world := jtracer.World{
		Objects: []jtracer.Shaper{
			floor, ceiling, westWall, eastWall, northWall, southWall, s1, s2, s3, s0,
		},
		Light: jtracer.NewPointLight(
			*jtracer.NewPoint(-4.9, 4.9, -1),
			jtracer.Color{Red: 1, Green: 1, Blue: 1},
		),
	}

	camera := jtracer.NewCamera(200, 200, 0.5)
	camera.Transform = jtracer.ViewTransform(
		jtracer.NewPoint(-4.5, 0.85, -4),
		jtracer.NewPoint(0, 0.85, 0),
		jtracer.NewVector(0, 1, 0),
	)
	canvas := camera.Render(world)
	err := canvas.SavePNG("out.png")
	if err != nil {
		panic(err)
	}
}
