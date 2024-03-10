package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	sk "github.com/renatopp/skald"
	c "github.com/renatopp/skald/components"

	s "github.com/renatopp/skald/systems"
)

var FollowMouse = sk.NewSystem(func(g *sk.Game) error {
	for _, box := range g.World.Entities() {
		t := c.Transform.Get(box)
		x, y := ebiten.CursorPosition()
		t.Position = sk.Vec2{X: float64(x), Y: float64(600 - y)}
	}

	return nil
})

func main() {
	game := sk.NewGame()

	box := sk.NewEntityWithOptions(sk.EntityOptions{
		Components: sk.Components{c.Transform, c.Sprite},
		OnSpawned: func(e *sk.EntityInstance) {
			t := c.Transform.Get(e)
			t.Position = sk.Vec2{X: 100, Y: 100}

			s := c.Sprite.Get(e)
			s.Texture = sk.NewEmptyTexture(50, 30)

			vector.DrawFilledRect(s.Texture.Image, -25, -15, 50, 30, color.White, true)
		},
	})

	game.AddSystem(s.SpriteRenderer)f
	game.AddSystem(FollowMouse)
	game.World.Spawn(box)

	game.Play()
}
