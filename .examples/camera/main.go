package main

import (
	"image/color"
	"math/rand"

	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/core"
)

func main() {
	game := sk.NewGame()
	game.Screen.SetColor(color.RGBA{66, 189, 66, 255})

	handle := game.Assets.Load(".examples/_assets/tree.png")

	rabbits := sk.NewEntityWithOptions(sk.EntityOptions{
		Components: sk.Components{core.Transform, core.Sprite},
		OnSpawned: func(e *sk.EntityInstance) {
			s := core.Sprite.Get(e)
			s.Texture = handle.AsTexture()
		},
	})

	game.AddSystem(core.SpriteRenderer)
	game.AddSystem(core.ScreenZoom)
	game.AddSystem(core.ScreenFlight)
	game.World.SpawnMulti(5000, rabbits, func(e *sk.EntityInstance) {
		t := core.Transform.Get(e)
		t.MoveTo(rand.Float64()*50-25, rand.Float64()*50-25)
	})
	game.Play()
}
