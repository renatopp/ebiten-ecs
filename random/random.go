package random

import "math/rand"

func Range(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func RangeInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func Float() float64 {
	return rand.Float64()
}

func Int() int {
	return rand.Int()
}

func Bool() bool {
	return rand.Intn(2) == 0
}

var Coin = Bool

func Dice(sides int) int {
	return rand.Intn(sides) + 1
}

func NDices(n, sides int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += Dice(sides)
	}
	return sum
}

// -1 to 1
func Polar() float64 {
	return rand.Float64()*2 - 1
}

// -1 or 1
func PolarInt() int {
	if Coin() {
		return 1
	}

	return -1
}

func Normal(mean, std float64) float64 {
	return rand.NormFloat64()*std + mean
}

func NormalInt(mean, std int) int {
	return int(rand.NormFloat64()*float64(std) + float64(mean))
}
