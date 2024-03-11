package core

import (
	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/inputs"
)

var ScreenZoom = sk.NewSystem(func(g *sk.Game) error {
	if inputs.MouseWheelY.IsDown() {
		wheel := inputs.MouseWheelY.GetFloat()
		g.Screen.Zoom(wheel)
	}

	return nil
})
