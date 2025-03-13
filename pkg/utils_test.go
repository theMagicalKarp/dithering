package pkg_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theMagicalKarp/dithering/pkg"
)

func TestHighIntToUInt8(t *testing.T) {
	t.Parallel()
	assert.Equal(t, uint8(255), pkg.IntToUInt8(255))
	assert.Equal(t, uint8(255), pkg.IntToUInt8(256))
	assert.Equal(t, uint8(255), pkg.IntToUInt8(1000))
}

func TestLowIntToUInt8(t *testing.T) {
	t.Parallel()
	assert.Equal(t, uint8(0), pkg.IntToUInt8(0))
	assert.Equal(t, uint8(0), pkg.IntToUInt8(-1))
	assert.Equal(t, uint8(0), pkg.IntToUInt8(-1000))
}

func TestIntToUInt8(t *testing.T) {
	t.Parallel()
	assert.Equal(t, uint8(200), pkg.IntToUInt8(200))
	assert.Equal(t, uint8(10), pkg.IntToUInt8(10))
}
