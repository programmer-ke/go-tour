package main

import (
	"fmt"
	"image"
	"image/color"
)

type Image struct {
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{
		uint8(x), uint8(y), 255, 255,
	}
}

func main() {
	m := Image{}
	fmt.Println(m.Bounds())
	fmt.Println(m.ColorModel())
	fmt.Println(m.At(0, 0))
}
