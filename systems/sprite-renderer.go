package systems

import (
	sk "github.com/renatopp/skald"
	c "github.com/renatopp/skald/components"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteRendererResult struct {
	Transform *c.DTransform
	Sprite    *c.DSprite
}

var SpriteRendererQuery = sk.NewQuery[SpriteRendererResult]()

var SpriteRenderer = sk.NewSystem(func(g *sk.Game) error {
	for _, r := range SpriteRendererQuery.Query() {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Reset()
		op.GeoM.Translate(r.Transform.Position.X, 600-r.Transform.Position.Y)
		op.GeoM.Scale(r.Transform.Scale.X, r.Transform.Scale.Y)
		op.GeoM.Rotate(r.Transform.Rotation)
		g.Renderer.Queue(r.Sprite.Layer, r.Sprite.ZIndex, r.Sprite.Texture.Image, op)
	}

	return nil
}, SpriteRendererQuery)
