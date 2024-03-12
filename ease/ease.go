package ease

import "math"

const c1 = 1.70158
const c2 = c1 * 1.525
const c3 = c1 + 1
const c4 = (2 * math.Pi) / 3
const c5 = (2 * math.Pi) / 4.5

func Linear(t float64) float64 {
	return t
}

func InSine(t float64) float64 {
	return 1 - math.Cos(t*(math.Pi/2))
}

func OutSine(t float64) float64 {
	return math.Sin(t * (math.Pi / 2))
}

func InOutSine(t float64) float64 {
	return -(math.Cos(math.Pi*t) - 1) / 2
}

func InQuad(t float64) float64 {
	return t * t
}

func OutQuad(t float64) float64 {
	return 1 - (1-t)*(1-t)
}

func InOutQuad(t float64) float64 {
	if t < 0.5 {
		return 2 * t * t
	}
	return 1 - math.Pow(-2*t+2, 2)/2
}

func InCubic(t float64) float64 {
	return t * t * t
}

func OutCubic(t float64) float64 {
	return 1 - math.Pow(1-t, 3)
}

func InOutCubic(t float64) float64 {
	if t < 0.5 {
		return 4 * t * t * t
	}
	return 1 - math.Pow(-2*t+2, 3)/2
}

func InQuart(t float64) float64 {
	return t * t * t * t
}

func OutQuart(t float64) float64 {
	return 1 - math.Pow(1-t, 4)
}

func InOutQuart(t float64) float64 {
	if t < 0.5 {
		return 8 * t * t * t * t
	}
	return 1 - math.Pow(-2*t+2, 4)/2
}

func InQuint(t float64) float64 {
	return t * t * t * t * t
}

func OutQuint(t float64) float64 {
	return 1 - math.Pow(1-t, 5)
}

func InOutQuint(t float64) float64 {
	if t < 0.5 {
		return 16 * t * t * t * t * t
	}
	return 1 - math.Pow(-2*t+2, 5)/2
}

func InExpo(t float64) float64 {
	if t == 0 {
		return 0
	}
	return math.Pow(2, 10*(t-1))
}

func OutExpo(t float64) float64 {
	if t == 1 {
		return 1
	}
	return 1 - math.Pow(2, -10*t)
}

func InOutExpo(t float64) float64 {
	if t == 0 {
		return 0
	}
	if t == 1 {
		return 1
	}
	if t < 0.5 {
		return math.Pow(2, 20*t-10) / 2
	}
	return (2 - math.Pow(2, -20*t+10)) / 2
}

func InCirc(t float64) float64 {
	return 1 - math.Sqrt(1-t*t)
}

func OutCirc(t float64) float64 {
	return math.Sqrt(1 - math.Pow(t-1, 2))
}

func InOutCirc(t float64) float64 {
	if t < 0.5 {
		return (1 - math.Sqrt(1-4*t*t)) / 2
	}
	return (math.Sqrt(-((2*t - 3) * (2*t - 1))) + 1) / 2
}

func InBack(t float64) float64 {
	return c3*t*t*t - c1*t*t
}

func OutBack(t float64) float64 {
	return 1 + c3*math.Pow(t-1, 3) + c1*math.Pow(t-1, 2)
}

func InOutBack(t float64) float64 {
	if t < 0.5 {
		return (math.Pow(2*t, 2) * ((c2+1)*2*t - c2)) / 2
	}
	return (math.Pow(2*t-2, 2)*((c2+1)*(t*2-2)+c2) + 2) / 2
}

func InElastic(t float64) float64 {
	if t == 0 {
		return 0
	}
	if t == 1 {
		return 1
	}
	return -math.Pow(2, 10*(t-1)) * math.Sin((t*10-10.75)*c4)
}

func OutElastic(t float64) float64 {
	if t == 0 {
		return 0
	}
	if t == 1 {
		return 1
	}
	return math.Pow(2, -10*t)*math.Sin((t*10-0.75)*c4) + 1
}

func InOutElastic(t float64) float64 {
	if t == 0 {
		return 0
	}
	if t == 1 {
		return 1
	}
	if t < 0.5 {
		return -(math.Pow(2, 20*t-10) * math.Sin((20*t-11.125)*c5)) / 2
	}
	return (math.Pow(2, -20*t+10) * math.Sin((20*t-11.125)*c5)) / 2
}

func InBounce(t float64) float64 {
	return 1 - OutBounce(1-t)
}

func OutBounce(t float64) float64 {
	if t < 1/2.75 {
		return 7.5625 * t * t
	} else if t < 2/2.75 {
		t -= 1.5 / 2.75
		return 7.5625*t*t + 0.75
	} else if t < 2.5/2.75 {
		t -= 2.25 / 2.75
		return 7.5625*t*t + 0.9375
	} else {
		t -= 2.625 / 2.75
		return 7.5625*t*t + 0.984375
	}
}

func InOutBounce(t float64) float64 {
	if t < 0.5 {
		return InBounce(t*2) / 2
	}
	return OutBounce(t*2-1)/2 + 0.5
}
