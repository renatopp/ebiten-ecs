package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/vector"
	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/core"
	"github.com/renatopp/skald/utils"
)

func main() {
	game := sk.NewGame()
	handlet := game.Assets.Load(".examples/_assets/tree.png")
	handle := game.Assets.Load(".examples/_assets/rabbitv3.png")

	bodies := sk.NewEntityWithOptions(sk.EntityOptions{
		Components: sk.Components{core.Transform, core.Sprite, RigidBody},
		OnSpawned: func(e *sk.EntityInstance) {
			r := RigidBody.Get(e)
			r.AsRandom()

			bounds := r.GetBounds()
			s := core.Sprite.Get(e)
			s.Texture = handlet.AsTexture()
			s.Texture = sk.NewEmptyTexture(int(bounds.Size.X), int(bounds.Size.Y))

			red := uint8(utils.Lerp(0, 255, r.Mass/100))
			c := color.RGBA{red, 0, 0, 255}

			// Draw the body
			if r.Shape == CIRCLE {
				vector.StrokeCircle(s.Texture.Image, float32(r.Radius), float32(r.Radius), float32(r.Radius)-1, 1, c, false)
				vector.StrokeLine(s.Texture.Image, float32(r.Radius), float32(r.Radius), float32(r.Radius), float32(0), 1, c, false)
			} else {
				vector.StrokeRect(s.Texture.Image, 1, 1, float32(r.Width)-2, float32(r.Height)-2, 1, c, false)
				vector.StrokeLine(s.Texture.Image, float32(r.Width/2), float32(r.Height/2), float32(r.Width/2), float32(0), 1, c, false)
			}
		},
	})

	rabbits := sk.NewEntityWithOptions(sk.EntityOptions{
		Components: sk.Components{core.Transform, core.Sprite},
		OnSpawned: func(e *sk.EntityInstance) {
			s := core.Sprite.Get(e)
			s.Texture = handle.AsTexture()
		},
	})

	game.AddSystem(core.SpriteRenderer)
	game.AddSystem(core.ScreenFlight)
	game.AddSystem(core.ScreenZoom)
	game.World.SpawnMulti(10, bodies, func(e *sk.EntityInstance) {
		t := core.Transform.Get(e)
		t.Position = sk.Vec2{X: rand.Float64()*10 - 5, Y: rand.Float64()*10 - 5}
	})
	game.World.SpawnMulti(10, rabbits, func(e *sk.EntityInstance) {
		t := core.Transform.Get(e)
		t.Position = sk.Vec2{X: rand.Float64()*10 - 5, Y: rand.Float64()*10 - 5}
	})

	game.Play()
}

// ----------------------------------------------------------------------------
// DRAW
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// RIGID BODY
// ----------------------------------------------------------------------------
var CIRCLE = 0
var BOX = 1

var MIN_BODY_SIZE float64 = 0.01 * 0.01
var MAX_BODY_SIZE float64 = 64 * 64
var MIN_DENSITY float64 = 0.01
var MAX_DENSITY float64 = 24

type DRigidBody struct {
	LinearVelocity  sk.Vec2
	AngularVelocity float64

	Mass        float64
	Density     float64
	Restitution float64
	Area        float64

	Static bool

	Shape  int // 0: Circle, 1: Box
	Radius float64
	Width  float64
	Height float64
}

func (r *DRigidBody) GetBounds() sk.Rect {
	if r.Shape == CIRCLE {
		return sk.Rect{
			Position: sk.Vec2{X: -r.Radius, Y: -r.Radius},
			Size:     sk.Vec2{X: r.Radius * 2, Y: r.Radius * 2},
		}
	} else {
		return sk.Rect{
			Position: sk.Vec2{X: -r.Width / 2, Y: -r.Height / 2},
			Size:     sk.Vec2{X: r.Width, Y: r.Height},
		}
	}
}

func (r *DRigidBody) AsCircle(radius, density, restitution float64, static bool) {
	area := 2 * radius * math.Pi

	if area > MAX_BODY_SIZE {
		area = MAX_BODY_SIZE
		println("warn: body size too big")
	}

	if area < MIN_BODY_SIZE {
		area = MIN_BODY_SIZE
		println("warn: body size too small")
	}

	if density < MIN_DENSITY {
		density = MIN_DENSITY
		println("warn: density too small")
	}

	if density > MAX_DENSITY {
		density = MAX_DENSITY
		println("warn: density too big")
	}

	r.Shape = CIRCLE
	r.Radius = radius
	r.Area = area
	r.Density = density
	r.Restitution = utils.Clamp(restitution, 0, 1)
	r.Static = static

	r.Mass = area * density
}

func (r *DRigidBody) AsBox(width, height, density, restitution float64, static bool) {
	area := width * height

	if area > MAX_BODY_SIZE {
		area = MAX_BODY_SIZE
		println("warn: body size too big")
	}

	if area < MIN_BODY_SIZE {
		area = MIN_BODY_SIZE
		println("warn: body size too small")
	}

	if density < MIN_DENSITY {
		density = MIN_DENSITY
		println("warn: density too small")
	}

	if density > MAX_DENSITY {
		density = MAX_DENSITY
		println("warn: density too big")
	}

	r.Shape = BOX
	r.Width = width
	r.Height = height
	r.Area = area
	r.Density = density
	r.Restitution = utils.Clamp(restitution, 0, 1)
	r.Static = static

	r.Mass = area * density
}

func (r *DRigidBody) AsRandom() {
	if rand.Intn(2) == 0 {
		r.AsCircle(rand.Float64()*23+5, rand.Float64()*20+1, rand.Float64(), false)
	} else {
		r.AsBox(rand.Float64()*23+5, rand.Float64()*23+5, rand.Float64()*20+1, rand.Float64(), false)
	}
}

var RigidBody = sk.NewComponentWithOptions(sk.ComponentOptions[DRigidBody]{
	OnCreated: func(r *DRigidBody) {
		r.Mass = 1
		r.Density = 1
		r.Restitution = 0.5
		r.Shape = CIRCLE
		r.Radius = 1
	},
})
