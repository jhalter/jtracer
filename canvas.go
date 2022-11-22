package jtracer

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
