package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHighIntToUInt8(t *testing.T) {
	t.Parallel()
	assert.Equal(t, uint8(255), intToUInt8(255))
	assert.Equal(t, uint8(255), intToUInt8(256))
	assert.Equal(t, uint8(255), intToUInt8(1000))
}

func TestLowIntToUInt8(t *testing.T) {
	t.Parallel()
	assert.Equal(t, uint8(0), intToUInt8(0))
	assert.Equal(t, uint8(0), intToUInt8(-1))
	assert.Equal(t, uint8(0), intToUInt8(-1000))
}

func TestIntToUInt8(t *testing.T) {
	t.Parallel()
	assert.Equal(t, uint8(200), intToUInt8(200))
	assert.Equal(t, uint8(10), intToUInt8(10))
}
