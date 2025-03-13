package pkg

func IntToUInt8(value int) uint8 {
	if value < 0 {
		return 0
	}

	if value > 255 {
		return 255
	}

	return uint8(value)
}
