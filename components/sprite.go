package components

import (
	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/assets"
)

type DSprite struct {
	Texture *assets.Texture
	Layer   uint
	ZIndex  int
}

var Sprite = sk.NewComponent[DSprite]()
