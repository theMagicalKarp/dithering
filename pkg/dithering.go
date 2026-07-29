// Package pkg implements Floyd–Steinberg dithering over 8-bit grayscale
// buffers, plus the helpers needed to move images in and out of that
// representation.
package pkg

import (
	"math"
)

// maxGray is the largest value an 8-bit grayscale sample can hold.
const maxGray = 255

// Floyd–Steinberg pushes the quantization error of each pixel onto the four
// neighbours that have not been visited yet, in these proportions.
const (
	diffusionRight     = 7
	diffusionDownLeft  = 3
	diffusionDown      = 5
	diffusionDownRight = 1
	diffusionTotal     = 16
)

// quantize snaps value to the nearest of the factor+1 evenly spaced levels and
// returns that level along with the error it introduced.
func quantize(value, factor int) (int, int) {
	quant := int(math.Round(float64(factor)*float64(value)/maxGray)) * (maxGray / factor)

	return quant, value - quant
}

// ApplyDither quantizes grayScale in place to factor+1 levels, diffusing the
// error of every pixel into its neighbours so the result keeps the apparent
// tonal range of the original.
func ApplyDither(grayScale [][]int, factor int) {
	var quantErr int

	width := len(grayScale)
	height := len(grayScale[0])

	for y := range height {
		for x := range width {
			grayScale[x][y], quantErr = quantize(grayScale[x][y], factor)

			if x+1 >= width || y+1 >= height || x == 0 {
				continue
			}

			grayScale[x+1][y] += quantErr * diffusionRight / diffusionTotal
			grayScale[x-1][y+1] += quantErr * diffusionDownLeft / diffusionTotal
			grayScale[x][y+1] += quantErr * diffusionDown / diffusionTotal
			grayScale[x+1][y+1] += quantErr * diffusionDownRight / diffusionTotal
		}
	}
}
