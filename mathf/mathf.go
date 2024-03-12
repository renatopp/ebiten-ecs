package mathf

import "math"

const Pi = math.Pi
const Pi2 = 2 * math.Pi
const PiHalf = math.Pi / 2
const E = math.E
const Phi = math.Phi
const Sqrt2 = math.Sqrt2
const SqrtE = math.SqrtE
const SqrtPi = math.SqrtPi
const SqrtPhi = math.SqrtPhi
const Ln2 = math.Ln2
const Ln10 = math.Ln10
const Log2E = math.Log2E
const Log10E = math.Log10E

const Deg2Rad float64 = math.Pi / 180
const Rad2Deg float64 = 180 / math.Pi

var Abs = math.Abs
var Acos = math.Acos
var Acosh = math.Acosh
var Asin = math.Asin
var Asinh = math.Asinh
var Atan = math.Atan
var Atan2 = math.Atan2
var Atanh = math.Atanh
var Cbrt = math.Cbrt
var Ceil = math.Ceil
var Copysign = math.Copysign
var Cos = math.Cos
var Cosh = math.Cosh
var Dim = math.Dim
var Erf = math.Erf
var Erfc = math.Erfc
var Erfcinv = math.Erfcinv
var Erfinv = math.Erfinv
var Exp = math.Exp
var Exp2 = math.Exp2
var Expm1 = math.Expm1
var FMA = math.FMA
var Float32bits = math.Float32bits
var Float32frombits = math.Float32frombits
var Float64bits = math.Float64bits
var Float64frombits = math.Float64frombits
var Floor = math.Floor
var Frexp = math.Frexp
var Gamma = math.Gamma
var Hypot = math.Hypot
var Ilogb = math.Ilogb
var Inf = math.Inf
var IsInf = math.IsInf
var IsNaN = math.IsNaN
var J0 = math.J0
var J1 = math.J1
var Jn = math.Jn
var Ldexp = math.Ldexp
var Lgamma = math.Lgamma
var Log = math.Log
var Log10 = math.Log10
var Log1p = math.Log1p
var Log2 = math.Log2
var Logb = math.Logb
var Max = math.Max
var Min = math.Min
var Mod = math.Mod
var Modf = math.Modf
var NaN = math.NaN
var Nextafter = math.Nextafter
var Nextafter32 = math.Nextafter32
var Pow = math.Pow
var Pow10 = math.Pow10
var Remainder = math.Remainder
var Round = math.Round
var RoundToEven = math.RoundToEven
var Signbit = math.Signbit
var Sin = math.Sin
var Sincos = math.Sincos
var Sinh = math.Sinh
var Sqrt = math.Sqrt
var Tan = math.Tan
var Tanh = math.Tanh
var Trunc = math.Trunc
var Y0 = math.Y0
var Y1 = math.Y1
var Yn = math.Yn

func Clamp(value, min, max float64) float64 {
	return math.Min(math.Max(value, min), max)
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

func Dot(x1, y1, x2, y2 float64) float64 {
	return x1*x2 + y1*y2
}

func Cross(x1, y1, x2, y2 float64) float64 {
	return x1*y2 - y1*x2
}

func Distance(x1, y1, x2, y2 float64) float64 {
	return Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

func Length(x, y float64) float64 {
	return Sqrt(x*x + y*y)
}

func LengthSqr(x, y float64) float64 {
	return x*x + y*y
}

func Normalize(x, y float64) (float64, float64) {
	l := Length(x, y)
	if l == 0 {
		return 0, 0
	}
	return x / l, y / l
}

func Angle(x, y float64) float64 {
	return Atan2(y, x)
}

func Rotate(x, y, angle float64) (float64, float64) {
	c := Cos(angle)
	s := Sin(angle)
	return x*c - y*s, x*s + y*c
}

func Lerp2(x1, y1, x2, y2, t float64) (float64, float64) {
	return x1 + (x2-x1)*t, y1 + (y2-y1)*t
}

func Clamp2(x, y, minX, minY, maxX, maxY float64) (float64, float64) {
	return Clamp(x, minX, maxX), Clamp(y, minY, maxY)
}
