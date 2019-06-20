package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	W int
	H int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.W, i.H)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (self Image) At(x, y int) color.Color {
	return color.RGBA{R: uint8(x), G: uint8(y), B: 255, A: 255}
}

func main() {
	m := Image{200, 200}
	pic.ShowImage(m)
}
