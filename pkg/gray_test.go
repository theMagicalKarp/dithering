package pkg_test

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/theMagicalKarp/dithering/pkg"
)

// testdataImage decodes the PNG fixture of the given name from testdata.
func testdataImage(t *testing.T, name string) image.Image {
	t.Helper()

	file, err := os.DirFS("testdata").Open(name)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, file.Close())
	})

	img, err := png.Decode(file)
	require.NoError(t, err)

	return img
}

// grayScaleOf reads img into the [x][y] sample layout the package uses. It is
// only a pixel read: the fixtures it is used on are already grayscale PNGs, so
// the color model conversion leaves their samples untouched.
func grayScaleOf(t *testing.T, img image.Image) [][]int {
	t.Helper()

	bounds := img.Bounds()

	samples := make([][]int, bounds.Max.X)
	for x := range bounds.Max.X {
		samples[x] = make([]int, bounds.Max.Y)

		for y := range bounds.Max.Y {
			gray, ok := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			require.True(t, ok)

			samples[x][y] = int(gray.Y)
		}
	}

	return samples
}

func TestNewGreyScale(t *testing.T) {
	t.Parallel()

	bounds := image.Rect(0, 0, 5, 5)

	expected := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	assert.Equal(t, expected, pkg.NewGreyScale(bounds))
}

func TestReadGreyScale(t *testing.T) {
	t.Parallel()

	src := testdataImage(t, "gopher.png")
	expected := grayScaleOf(t, testdataImage(t, "gopher_gray.png"))

	dst := pkg.NewGreyScale(src.Bounds())
	pkg.ReadGreyScale(src, dst)

	assert.Equal(t, expected, dst)
}

func TestWriteGrayScale(t *testing.T) {
	t.Parallel()

	expected := testdataImage(t, "gopher_gray.png")
	bounds := expected.Bounds()

	dst := image.NewNRGBA(bounds)
	pkg.WriteGrayScale(grayScaleOf(t, expected), dst)

	for x := range bounds.Max.X {
		for y := range bounds.Max.Y {
			gray, ok := color.GrayModel.Convert(expected.At(x, y)).(color.Gray)
			require.True(t, ok)

			want := color.NRGBA{R: gray.Y, G: gray.Y, B: gray.Y, A: math.MaxUint8}
			require.Equal(t, want, dst.NRGBAAt(x, y), "pixel (%d, %d)", x, y)
		}
	}
}
