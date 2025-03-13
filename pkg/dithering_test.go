package pkg_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/dithering/pkg"
)

func TestApplyBasicDither(t *testing.T) {
	t.Parallel()

	img := [][]int{
		{
			0, 200, 0, 100, 250,
		},
		{
			201, 45, 194, 140, 62,
		},
		{
			202, 237, 172, 113, 93,
		},
		{
			211, 90, 39, 73, 21,
		},
		{
			59, 30, 70, 153, 125,
		},
	}

	pkg.ApplyDither(img, 1)

	expected := [][]int{
		{
			0, 255, 0, 0, 255,
		},
		{
			255, 0, 255, 0, 0,
		},
		{
			255, 255, 255, 0, 255,
		},
		{
			255, 0, 0, 0, 0,
		},
		{
			0, 0, 0, 255, 255,
		},
	}

	assert.Equal(t, expected, img)
}

func TestApplyBasicDitherWithFactor(t *testing.T) {
	t.Parallel()

	img := [][]int{
		{
			0, 200, 0, 100, 250,
		},
		{
			201, 45, 194, 140, 62,
		},
		{
			202, 237, 172, 113, 93,
		},
		{
			211, 90, 39, 73, 21,
		},
		{
			59, 30, 70, 153, 125,
		},
	}

	pkg.ApplyDither(img, 3)

	expected := [][]int{
		{
			0, 170, 0, 85, 255,
		},
		{
			170, 85, 170, 170, 85,
		},
		{
			255, 255, 170, 85, 85,
		},
		{
			170, 85, 0, 85, 0,
		},
		{
			85, 0, 85, 170, 85,
		},
	}

	assert.Equal(t, expected, img)
}
