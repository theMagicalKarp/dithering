package pkg

import (
	"image"
	"image/color"
)

// Rec. 601 luma weights scaled by 1<<16, matching the stdlib's own
// color.GrayModel conversion. lumaRound is the half-step that makes the shift
// round rather than truncate, and lumaShift undoes both that scaling and the
// low byte of RGBA's 16-bit channels in one go.
const (
	lumaRed   = 19595
	lumaGreen = 38470
	lumaBlue  = 7471
	lumaRound = 1 << 15
	lumaShift = 24
)

// ReadableImage is the subset of image.Image this package needs to sample an
// image.
type ReadableImage interface {
	At(x, y int) color.Color
	Bounds() image.Rectangle
}

// WritableImage is a ReadableImage whose pixels can also be assigned, such as
// an *image.NRGBA.
type WritableImage interface {
	ReadableImage
	Set(x, y int, c color.Color)
}

// NewGreyScale allocates a zeroed grayscale buffer covering bounds, indexed as
// buffer[x][y].
func NewGreyScale(bounds image.Rectangle) [][]int {
	grayscale := make([][]int, bounds.Max.X)
	for i := range bounds.Max.X {
		grayscale[i] = make([]int, bounds.Max.Y)
	}

	return grayscale
}

// ReadGreyScale fills dst with the luminance of every pixel in src. dst must
// already cover src's bounds, as returned by NewGreyScale.
func ReadGreyScale(src ReadableImage, dst [][]int) {
	bounds := src.Bounds()
	for x := range bounds.Max.X {
		for y := range bounds.Max.Y {
			dst[x][y] = toGray(src.At(x, y))
		}
	}
}

// WriteGrayScale paints the samples in src onto dst as gray pixels, clamping
// any that error diffusion pushed outside the 0-255 range.
func WriteGrayScale(src [][]int, dst WritableImage) {
	bounds := dst.Bounds()
	for x := range bounds.Max.X {
		for y := range bounds.Max.Y {
			dst.Set(x, y, color.Gray{Y: IntToUInt8(src[x][y])})
		}
	}
}

// toGray reduces a color to its perceived brightness as an 8-bit sample.
func toGray(c color.Color) int {
	r, g, b, _ := c.RGBA()

	return int((lumaRed*r + lumaGreen*g + lumaBlue*b + lumaRound) >> lumaShift)
}
