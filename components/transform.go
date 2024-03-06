package components

import (
	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/mathz"
)

type DTransform struct {
	Position mathz.Vec2
	Scale    mathz.Vec2
	Rotation float64
}

var Transform = sk.NewComponent[DTransform]()
