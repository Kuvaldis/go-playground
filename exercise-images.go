package main

import (
	"image"
	"image/color"
)

type Image struct {
	width  int
	height int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: i.width, Y: i.height}}
}

func (i Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{R: v, G: v, B: 255, A: 255}
}

func main8() {
	//m := Image{width: 200, height: 100}
	//pic.ShowImage(m)
}
