package sk

import "math"

type Vec2 struct {
	X, Y float64
}

func (v Vec2) Equal(v2 Vec2) bool {
	return v.X == v2.X && v.Y == v2.Y
}

func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v.X + v2.X, v.Y + v2.Y}
}

func (v Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{v.X - v2.X, v.Y - v2.Y}
}

func (v Vec2) Mul(v2 Vec2) Vec2 {
	return Vec2{v.X * v2.X, v.Y * v2.Y}
}

func (v Vec2) Div(v2 Vec2) Vec2 {
	return Vec2{v.X / v2.X, v.Y / v2.Y}
}

func (v Vec2) AddS(s float64) Vec2 {
	return Vec2{v.X + s, v.Y + s}
}

func (v Vec2) SubS(s float64) Vec2 {
	return Vec2{v.X - s, v.Y - s}
}

func (v Vec2) MulS(s float64) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

func (v Vec2) DivS(s float64) Vec2 {
	return Vec2{v.X / s, v.Y / s}
}

func (v Vec2) Dot(v2 Vec2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vec2) Cross(v2 Vec2) float64 {
	return v.X*v2.Y - v.Y*v2.X
}

func (v Vec2) Distance(v2 Vec2) float64 {
	return math.Sqrt((v2.X-v.X)*(v2.X-v.X) + (v2.Y-v.Y)*(v2.Y-v.Y))
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2) LengthSqr() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2) Normalized() Vec2 {
	l := v.Length()
	if l == 0 {
		return Vec2{}
	}

	return v.DivS(l)
}

func (v Vec2) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

func (v Vec2) Rotate(angle float64) Vec2 {
	c := math.Cos(angle)
	s := math.Sin(angle)
	return Vec2{v.X*c - v.Y*s, v.X*s + v.Y*c}
}

func (v Vec2) Lerp(v2 Vec2, t float64) Vec2 {
	return v.Add(v2.Sub(v).MulS(t))
}

func (v Vec2) Clamp(min, max Vec2) Vec2 {
	return Vec2{
		math.Min(math.Max(v.X, min.X), max.X),
		math.Min(math.Max(v.Y, min.Y), max.Y),
	}
}

func (v Vec2) ClampS(min, max float64) Vec2 {
	return Vec2{
		math.Min(math.Max(v.X, min), max),
		math.Min(math.Max(v.Y, min), max),
	}
}

func (v Vec2) Abs() Vec2 {
	return Vec2{math.Abs(v.X), math.Abs(v.Y)}
}

func (v Vec2) Min(v2 Vec2) Vec2 {
	return Vec2{math.Min(v.X, v2.X), math.Min(v.Y, v2.Y)}
}

func (v Vec2) Max(v2 Vec2) Vec2 {
	return Vec2{math.Max(v.X, v2.X), math.Max(v.Y, v2.Y)}
}

func (v Vec2) Floor() Vec2 {
	return Vec2{math.Floor(v.X), math.Floor(v.Y)}
}

func (v Vec2) Ceil() Vec2 {
	return Vec2{math.Ceil(v.X), math.Ceil(v.Y)}
}

func (v Vec2) Round() Vec2 {
	return Vec2{math.Round(v.X), math.Round(v.Y)}
}
