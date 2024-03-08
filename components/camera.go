package components

import (
	sk "github.com/renatopp/skald"
	sk1 "github.com/renatopp/skald/sk"
)

type DCamera struct {
	Rect sk1.Rect
}

var Camera = sk.NewComponent[DCamera]()
