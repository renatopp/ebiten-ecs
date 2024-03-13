package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/vector"
	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/core"
	"github.com/renatopp/skald/feats/physics"
	"github.com/renatopp/skald/inputs"
	"github.com/renatopp/skald/mathf"
	"github.com/renatopp/skald/random"
)

func main() {
	game := sk.NewGame()
	// handlet := game.Assets.Load(".examples/_assets/tree.png")
	// handle := game.Assets.Load(".examples/_assets/rabbitv3.png")
	//
	bodies := sk.NewEntityWithOptions(sk.EntityOptions{
		Components: sk.Components{core.Transform, core.Sprite, physics.RigidBody},
		OnSpawned: func(e *sk.EntityInstance) {
			ppu := game.Screen.PixelsPerUnit

			r := physics.RigidBody.Get(e)
			r.AsRandom()

			bounds := r.GetBounds()
			s := core.Sprite.Get(e)
			// s.Texture = handlet.AsTexture()
			s.Texture = sk.NewEmptyTexture(int(bounds.Size.X*ppu), int(bounds.Size.Y*ppu))

			red := uint8(mathf.Lerp(0, 255, r.Mass/5))
			c := color.RGBA{red, 0, 0, 255}

			// Draw the body
			if r.Shape == physics.CIRCLE {
				vector.StrokeCircle(s.Texture.Image, float32(ppu*r.Radius), float32(ppu*r.Radius), float32(ppu*r.Radius)-1, 1, c, false)
				vector.StrokeLine(s.Texture.Image, float32(ppu*r.Radius), float32(ppu*r.Radius), float32(ppu*r.Radius), float32(ppu*0), 1, c, false)
			} else {
				vector.StrokeRect(s.Texture.Image, 1, 1, float32(ppu*r.Width)-2, float32(ppu*r.Height)-2, 1, c, false)
				vector.StrokeLine(s.Texture.Image, float32(ppu*r.Width/2), float32(ppu*r.Height/2), float32(ppu*r.Width/2), float32(ppu*0), 1, c, false)
			}
		},
	})

	game.AddSystem(core.SpriteRenderer)
	game.AddSystem(core.ScreenFlight)
	game.AddSystem(core.ScreenZoom)
	game.AddSystem(physics.PhysicsUpdater)
	game.AddSystem(MoveBall)
	game.World.SpawnMulti(20, bodies, func(e *sk.EntityInstance) {
		t := core.Transform.Get(e)
		t.MoveTo(random.Range(-2.5, 2.5), random.Range(-2.5, 2.5))
	})

	game.Play()
}

// ----------------------------------------------------------------------------
// PHYSICS
// ----------------------------------------------------------------------------

var MoveBall = sk.NewSystem(func(g *sk.Game) error {
	for i, r := range physics.TransformAndBodyQuery.Query() {
		// r.Transform.Position = r.Transform.Position.Add(r.physics.RigidBody.LinearVelocity)
		if i == 0 {
			d := sk.Vec2{}
			if inputs.KeyArrowUp.IsDown() {
				d.Y += -1
			}
			if inputs.KeyArrowDown.IsDown() {
				d.Y += 1
			}
			if inputs.KeyArrowLeft.IsDown() {
				d.X += -1
			}
			if inputs.KeyArrowRight.IsDown() {
				d.X += 1
			}

			d = d.Normalized().MulS(g.Timer.DeltaTime * 2)
			r.Transform.MoveBy(d.X, d.Y)

			if inputs.Key2.IsDown() {
				r.Transform.RotateBy(g.Timer.DeltaTime * 180)
			}

			if inputs.Key1.IsDown() {
				r.Transform.RotateBy(-g.Timer.DeltaTime * 180)
			}
		}

	}
	return nil
}, physics.TransformAndBodyQuery)
