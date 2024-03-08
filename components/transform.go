package components

import (
	sk "github.com/renatopp/skald"
	sk1 "github.com/renatopp/skald/sk"
)

type DTransform struct {
	Position sk1.Vec2
	Scale    sk1.Vec2
	Rotation float64
}

var Transform = sk.NewComponent[DTransform]()
