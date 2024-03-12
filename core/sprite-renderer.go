package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/mathf"
)

type SpriteRendererResult struct {
	Transform *DTransform
	Sprite    *DSprite
}

var SpriteRendererQuery = sk.NewQuery[SpriteRendererResult]()

var SpriteRenderer = sk.NewSystem(func(g *sk.Game) error {
	for _, r := range SpriteRendererQuery.Query() {
		if !r.Transform.enabled {
			continue
		}

		bounds := r.Sprite.Texture.Image.Bounds()
		halfX, halfY := float64(bounds.Max.X)/2, float64(bounds.Max.Y)/2

		var ops *ebiten.DrawImageOptions
		if r.Transform.version == r.Sprite.transformVersion {
			ops = r.Sprite.ops
		} else {
			ops = &ebiten.DrawImageOptions{}
			ops.GeoM.Translate(-halfX, -halfY)
			ops.GeoM.Scale(r.Transform.scaleX, r.Transform.scaleY)
			ops.GeoM.Rotate(r.Transform.rotation * mathf.Deg2Rad)
			ops.GeoM.Translate(r.Transform.posX*g.Screen.PixelsPerUnit, r.Transform.posY*g.Screen.PixelsPerUnit)
			r.Sprite.transformVersion = r.Transform.version
			r.Sprite.ops = ops
		}
		g.Renderer.Queue(r.Sprite.Layer, r.Sprite.ZIndex, r.Sprite.Texture.Image, ops)
	}

	return nil
}, SpriteRendererQuery)
