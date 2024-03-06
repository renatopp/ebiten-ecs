package components

import (
	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/mathz"
)

type DCamera struct {
	Rect mathz.Rect
}

var Camera = sk.NewComponent[DCamera]()
