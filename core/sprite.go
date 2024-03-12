package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	sk "github.com/renatopp/skald"
)

type DSprite struct {
	Texture *sk.Texture
	Layer   uint
	ZIndex  int

	ops              *ebiten.DrawImageOptions
	transformVersion uint64
}

var Sprite = sk.NewComponent[DSprite]()
