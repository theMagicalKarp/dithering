package pkg

// IntToUInt8 clamps value into the 0-255 range of an 8-bit grayscale sample.
// Error diffusion routinely pushes samples past both ends of that range, so
// callers rely on the clamping rather than treating it as an error.
func IntToUInt8(value int) uint8 {
	if value < 0 {
		return 0
	}

	if value > maxGray {
		return maxGray
	}

	return uint8(value)
}
