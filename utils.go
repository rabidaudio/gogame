package main

import "image/color"

func HexColor(val uint32) color.RGBA {
	return color.RGBA{
		R: uint8(val >> 16),
		G: uint8(val >> 8),
		B: uint8(val >> 0),
		A: 0xFF,
	}
}

func SUnit(v int) float64 {
	return float64(v)
}
