package jtracer

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Canvas struct {
	Width  int
	Height int
	Data   [][]Color
}

func NewCanvas(width int, height int) *Canvas {
	data := make([][]Color, height)
	for i := range data {
		data[i] = make([]Color, width)
	}
	return &Canvas{width, height, data}
}

func (c *Canvas) WritePixel(x, y int, color *Color) {
	c.Data[y][x] = *color
}

func (c *Canvas) PixelAt(x int, y int) *Color {
	return &c.Data[y][x]
}

func (c *Canvas) ToImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
	for y := range c.Data {
		for x := range c.Data[y] {
			canvasColor := c.Data[y][x]
			red, green, blue := canvasColor.Normalize()
			img.Set(x, y, color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: 255})
		}
	}
	return img
}

func (c *Canvas) SavePNG(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	err = png.Encode(f, c.ToImage())
	if err != nil {
		return err
	}
	return nil
}
