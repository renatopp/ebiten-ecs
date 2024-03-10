package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	sk "github.com/renatopp/skald"
)

type DCamera struct {
	Size    sk.Vec2
	Surface *ebiten.Image
	Zoom    float64
}

var Camera = sk.NewComponentWithOptions(sk.ComponentOptions[DCamera]{
	Dependencies: sk.Components{Transform},
	OnCreated: func(d *DCamera) {
		d.Size = sk.Vec2{X: sk.SCREEN_WIDTH, Y: sk.SCREEN_HEIGHT}
		d.Surface = ebiten.NewImage(int(d.Size.X), int(d.Size.Y))
		d.Zoom = 1
	},
})
