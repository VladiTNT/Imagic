package process

import (
	"image/color"
	"math"
)

func clamp(y float64) uint8 {
	if y > 255 {
		return 255
	}
	if y < 0 {
		return 0
	}
	return uint8(math.Round(y))
}

// Stupid function because color.Color interface doesn't automatically correct color values.
func bitShiftCorrection(r, g, b, a uint32) (uint32, uint32, uint32, uint32) {
	return r >> 8, g >> 8, b >> 8, a >> 8
}

func averageColor(colors []color.Color) color.Color {
	var sumR, sumG, sumB, sumA uint64

	for _, c := range colors {
		r, g, b, a := c.RGBA()
		sumR += uint64(r)
		sumG += uint64(g)
		sumB += uint64(b)
		sumA += uint64(a)
	}

	n := uint64(len(colors))
	return color.RGBA64{
		R: uint16(sumR / n),
		G: uint16(sumG / n),
		B: uint16(sumB / n),
		A: uint16(sumA / n),
	}
}

func averageGray(c color.Color) color.Gray16 {
	r, g, b, a := c.RGBA()
	s := (r + g + b + a) / 4
	return color.Gray16{Y: uint16(s)}
}
