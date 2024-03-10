package core

import (
	sk "github.com/renatopp/skald"
)

type DTransform struct {
	Enabled  bool
	Position sk.Vec2
	Scale    sk.Vec2
	Rotation float64
}

var Transform = sk.NewComponentWithOptions(sk.ComponentOptions[DTransform]{
	OnCreated: func(c *DTransform) {
		c.Enabled = true
		c.Position = sk.Vec2{X: 0, Y: 0}
		c.Scale = sk.Vec2{X: 1, Y: 1}
		c.Rotation = 0
	},
})
