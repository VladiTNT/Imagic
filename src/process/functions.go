package process

import "math"

func clamp(y float64) uint8 {
	if y > 255 {
		return 255
	}
	if y < 0 {
		return 0
	}
	return uint8(math.Round(y))
}

func bitShiftCorrection(r, g, b, a uint32) (uint32, uint32, uint32, uint32) {
	return r >> 8, g >> 8, b >> 8, a >> 8
}
