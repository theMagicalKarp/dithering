package pkg

import (
	"math"
)

func quantize(value, factor int) (int, int) {
	quant := int(math.Round(float64(factor)*float64(value)/255.0)) * (255 / factor)

	return quant, value - quant
}

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

			grayScale[x+1][y] += quantErr * 7 / 16
			grayScale[x-1][y+1] += quantErr * 3 / 16
			grayScale[x][y+1] += quantErr * 5 / 16
			grayScale[x+1][y+1] += quantErr * 1 / 16
		}
	}
}
