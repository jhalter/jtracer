package jtracer

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
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

const ppmHeader = `
P3
%d %d
255
`

func (c *Canvas) ToPPM() string {
	var ppm strings.Builder
	ppm.WriteString(fmt.Sprintf(ppmHeader, c.Width, c.Height))

	for y := range c.Data {
		var line strings.Builder
		leadingSpace := ""
		for x := range c.Data[y] {
			color := c.Data[y][x]
			red, green, blue := color.Normalize()
			line.WriteString(fmt.Sprintf("%s%d %d %d", leadingSpace, red, green, blue))
			leadingSpace = " "
			if line.Len() > 70 {
				leadingSpace = "\n"
				ppm.WriteString(line.String())
				line.Reset()
			}
		}
		ppm.WriteString(line.String())
		ppm.WriteString("\n")
	}
	ppm.WriteString("\n")
	return ppm.String()
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
