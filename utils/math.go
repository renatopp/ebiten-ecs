package utils

import "math"

func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func ClampInt(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func Lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

func EaseInQuad(t float64) float64 {
	return t * t
}

const Deg2Rad float64 = math.Pi / 180
const Rad2Deg float64 = 180 / math.Pi
