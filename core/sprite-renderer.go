package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	sk "github.com/renatopp/skald"
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

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Reset()
		op.GeoM.Translate(r.Transform.Position.X, r.Transform.Position.Y)
		op.GeoM.Scale(r.Transform.Scale.X, r.Transform.Scale.Y)
		op.GeoM.Rotate(r.Transform.Rotation)
		g.Renderer.Queue(r.Sprite.Layer, r.Sprite.ZIndex, r.Sprite.Texture.Image, op)
	}

	return nil
}, SpriteRendererQuery)
