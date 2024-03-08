package components

import (
	sk "github.com/renatopp/skald"
)

type DCamera struct {
	Rect sk.Rect
}

var Camera = sk.NewComponent[DCamera]()
