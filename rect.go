package sk

import "math"

type Rect struct {
	Position Vec2
	Size     Vec2
}

func (r Rect) Left() float64 {
	return r.Position.X
}

func (r Rect) Right() float64 {
	return r.Position.X + r.Size.X
}

func (r Rect) Top() float64 {
	return r.Position.Y
}

func (r Rect) Bottom() float64 {
	return r.Position.Y + r.Size.Y
}

func (r Rect) Min() Vec2 {
	return r.Position
}

func (r Rect) Max() Vec2 {
	return r.Position.Add(r.Size)
}

func (r Rect) Center() Vec2 {
	return r.Position.Add(r.Size.DivS(2))
}

func (r Rect) Contains(v Vec2) bool {
	return v.X >= r.Left() && v.X <= r.Right() &&
		v.Y >= r.Top() && v.Y <= r.Bottom()
}

func (r Rect) Intersects(r2 Rect) bool {
	return r.Left() <= r2.Right() && r.Right() >= r2.Left() &&
		r.Top() <= r2.Bottom() && r.Bottom() >= r2.Top()
}

func (r Rect) Intersect(r2 Rect) Rect {
	if !r.Intersects(r2) {
		return Rect{}
	}

	return Rect{
		Position: Vec2{
			X: math.Max(r.Left(), r2.Left()),
			Y: math.Max(r.Top(), r2.Top()),
		},
		Size: Vec2{
			X: math.Min(r.Right(), r2.Right()) - math.Max(r.Left(), r2.Left()),
			Y: math.Min(r.Bottom(), r2.Bottom()) - math.Max(r.Top(), r2.Top()),
		},
	}
}

func (r Rect) Union(r2 Rect) Rect {
	return Rect{
		Position: Vec2{
			X: math.Min(r.Left(), r2.Left()),
			Y: math.Min(r.Top(), r2.Top()),
		},
		Size: Vec2{
			X: math.Max(r.Right(), r2.Right()) - math.Min(r.Left(), r2.Left()),
			Y: math.Max(r.Bottom(), r2.Bottom()) - math.Min(r.Top(), r2.Top()),
		},
	}
}

func (r Rect) Expand(v Vec2) Rect {
	return Rect{
		Position: r.Position.Sub(v),
		Size:     r.Size.Add(v.MulS(2)),
	}
}

func (r Rect) ExpandS(s float64) Rect {
	return Rect{
		Position: r.Position.SubS(s),
		Size:     r.Size.AddS(s * 2),
	}
}

func (r Rect) Shrink(v Vec2) Rect {
	return Rect{
		Position: r.Position.Add(v),
		Size:     r.Size.Sub(v.MulS(2)),
	}
}

func (r Rect) ShrinkS(s float64) Rect {
	return Rect{
		Position: r.Position.AddS(s),
		Size:     r.Size.SubS(s * 2),
	}
}
