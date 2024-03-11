package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/vector"
	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/core"
)

func main() {
	game := sk.NewGame()
	handle := game.Assets.Load(".examples/_assets/rabbitv3.png")

	shapes := sk.NewEntity(core.Transform, core.Sprite)
	game.AddSystem(core.SpriteRenderer)
	game.World.SpawnMulti(10, shapes, func(e *sk.EntityInstance) {
		t := core.Transform.Get(e)
		t.Position = sk.Vec2{X: rand.Float64()*10 - 5, Y: rand.Float64()*10 - 5}

		s := core.Sprite.Get(e)
		s.Texture = handle.AsTexture()
		s.Texture = sk.NewEmptyTexture(25, 32)
		// s.Texture.Image.Fill(color.RGBA{123, 12, 32, 255})
		vector.StrokeRect(s.Texture.Image, 1, 1, 10, 10, 1, color.Black, false)
	})

	game.Play()
}
