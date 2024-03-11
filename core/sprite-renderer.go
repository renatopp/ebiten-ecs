package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/utils"
)

type SpriteRendererResult struct {
	Transform *DTransform
	Sprite    *DSprite
}

var SpriteRendererQuery = sk.NewQuery[SpriteRendererResult]()

var SpriteRenderer = sk.NewSystem(func(g *sk.Game) error {
	for _, r := range SpriteRendererQuery.Query() {
		if !r.Transform.Enabled {
			continue
		}

		bounds := r.Sprite.Texture.Image.Bounds()

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Reset()
		op.GeoM.Translate(-float64(bounds.Max.X)/2, -float64(bounds.Max.Y)/2)
		op.GeoM.Rotate(r.Transform.Rotation * utils.Deg2Rad)
		op.GeoM.Translate(r.Transform.Position.X*g.Screen.PixelsPerUnit, r.Transform.Position.Y*g.Screen.PixelsPerUnit)
		op.GeoM.Scale(r.Transform.Scale.X, r.Transform.Scale.Y)
		g.Renderer.Queue(r.Sprite.Layer, r.Sprite.ZIndex, r.Sprite.Texture.Image, op)
	}

	return nil
}, SpriteRendererQuery)
