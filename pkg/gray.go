package pkg

import (
	"image"
	"image/color"
)

func toGray(c color.Color) int {
	r, g, b, _ := c.RGBA()

	return int((19595*r + 38470*g + 7471*b + 1<<15) >> 24)
}

type ReadableImage interface {
	At(x, y int) color.Color
	Bounds() image.Rectangle
}

type WritableImage interface {
	ReadableImage
	Set(x, y int, c color.Color)
}

func NewGreyScale(bounds image.Rectangle) [][]int {
	grayscale := make([][]int, bounds.Max.X)
	for i := 0; i < bounds.Max.X; i++ {
		grayscale[i] = make([]int, bounds.Max.Y)
	}

	return grayscale
}

func ReadGreyScale(src ReadableImage, dst [][]int) {
	bounds := src.Bounds()
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			dst[x][y] = toGray(src.At(x, y))
		}
	}
}

func WriteGrayScale(src [][]int, dst WritableImage) {
	bounds := dst.Bounds()
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			dst.Set(x, y, color.Gray{intToUInt8(src[x][y])})
		}
	}
}
