package core

import (
	sk "github.com/renatopp/skald"
)

type DSprite struct {
	Texture *sk.Texture
	Layer   uint
	ZIndex  int
}

var Sprite = sk.NewComponent[DSprite]()
