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
