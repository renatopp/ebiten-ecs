package components

import (
	sk "github.com/renatopp/skald"
)

type DTransform struct {
	Enabled  bool
	Position sk.Vec2
	Scale    sk.Vec2
	Rotation float64
}

var Transform = sk.NewComponent[DTransform]()
