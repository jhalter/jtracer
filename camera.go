package jtracer

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"math"
	"sync"
)

type Camera struct {
	Hsize      float64
	Vsize      float64
	Fov        float64
	Transform  Matrix
	HalfWidth  float64
	HalfHeight float64
	PixelSize  float64
}

func NewCamera(hsize, vsize, fov float64) Camera {
	c := Camera{
		Hsize:     hsize,
		Vsize:     vsize,
		Fov:       fov,
		Transform: IdentityMatrix,
	}

	halfView := math.Tan(c.Fov / 2)
	aspect := c.Hsize / c.Vsize

	if aspect >= 1.0 {
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspect
	} else {
		c.HalfWidth = halfView * aspect
		c.HalfHeight = halfView
	}

	c.PixelSize = (c.HalfWidth * 2) / c.Hsize

	return c
}

func (c Camera) RayForPixel(px, py float64) Ray {
	// the offset from the edge of the canvas to the pixel's center
	xOffset := (px + 0.5) * c.PixelSize
	yOffset := (py + 0.5) * c.PixelSize

	// the untransformed coordinates of the pixel in world space.
	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset

	// using the camera matrix, transform the canvas point and the origin,
	// and then compute the ray's direction vector.
	// (remember that the canvas is at z=-1)

	pixel := c.Transform.Inverse().MultiplyByTuple(*NewPoint(worldX, worldY, -1))
	origin := c.Transform.Inverse().MultiplyByTuple(*NewPoint(0, 0, 0))
	direction := pixel.Subtract(&origin).Normalize()

	return Ray{origin, *direction}
}

const RendererCount = 8
const MaxReflections = 5

func (c Camera) Render(w World) Canvas {
	image := NewCanvas(int(c.Hsize), int(c.Vsize))

	var wg sync.WaitGroup
	wg.Add(RendererCount)

	waitCh := make(chan struct{})

	yComplete := 0
	yDone := make(chan int, 1024)
	go func() {
		for {
			<-yDone
			yComplete++
			spew.Dump(yComplete)
			if yComplete%100 == 0 {
				fmt.Printf("%v", float64(yComplete)/c.Vsize)
			}
		}
	}()

	go func() {
		for i := 0; i < RendererCount; i++ {
			chunkSize := c.Vsize / RendererCount
			yStart := math.Floor(float64(i) * chunkSize)
			yEnd := math.Ceil(yStart + chunkSize - 1)

			go func() {
				defer wg.Done()
				fmt.Printf("%v to %v\n", yStart, yEnd)
				for y := yStart; y <= yEnd; y++ {
					yDone <- 1
					for x := 0.0; x < c.Hsize; x++ {
						r := c.RayForPixel(x, y)
						color := w.ColorAt(r, MaxReflections)
						image.WritePixel(int(x), int(y), &color)
						//messages <- Pixel{int(x), int(y), color}
					}
				}
				//fmt.Printf("RENDERER DONE\n")
			}()
		}

		wg.Wait()
		close(waitCh)
	}()

	// Block until the wait group is done
	<-waitCh

	return *image
}
