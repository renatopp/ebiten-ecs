package core

import (
	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/inputs"
)

var ScreenZoom = sk.NewSystem(func(g *sk.Game) error {
	if inputs.MouseWheelY.IsDown() {
		g.Screen.Zoom(inputs.MouseWheelX.GetFloat())
	}

	return nil
})
