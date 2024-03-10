package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/vector"
	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/core"
)

var CIRCLE = 0
var BOX = 1

type DBody struct {
	LinearVelocity  sk.Vec2
	AngularVelocity float64

	Mass        float64
	Density     float64
	Restitution float64
	Area        float64

	Static bool

	Shape  int
	Radius float64
	Width  float64
	Height float64
}

// var CiBody = skald.NewComponent()

func main() {
	game := sk.NewGame()

	box := sk.NewEntityWithOptions(sk.EntityOptions{
		Components: sk.Components{core.Transform, core.Sprite},
		OnSpawned: func(e *sk.EntityInstance) {
			t := core.Transform.Get(e)
			t.Position = sk.Vec2{X: 400, Y: 300}

			s := core.Sprite.Get(e)
			s.Texture = sk.NewEmptyTexture(30, 30)

			vector.DrawFilledRect(s.Texture.Image, 0, 0, 30, 30, color.Transparent, true)
		},
	})

	game.AddSystem(core.SpriteRenderer)
	game.World.Spawn(box)

	game.Play()
}
