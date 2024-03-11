package core

import (
	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/inputs"
)

var SCREEN_FLIGHT_MOVE_SPEED float64 = 10
var SCREEN_FLIGHT_ROTATE_SPEED float64 = 90

var ScreenFlight = sk.NewSystem(func(g *sk.Game) error {
	dt := g.Timer.GetDeltaTime()
	if inputs.KeyW.IsDown() {
		g.Screen.Move(0, -SCREEN_FLIGHT_MOVE_SPEED*dt)
	} else if inputs.KeyS.IsDown() {
		g.Screen.Move(0, SCREEN_FLIGHT_MOVE_SPEED*dt)
	}

	if inputs.KeyA.IsDown() {
		g.Screen.Move(-SCREEN_FLIGHT_MOVE_SPEED*dt, 0)
	} else if inputs.KeyD.IsDown() {
		g.Screen.Move(SCREEN_FLIGHT_MOVE_SPEED*dt, 0)
	}

	if inputs.KeyQ.IsDown() {
		g.Screen.Rotate(-SCREEN_FLIGHT_ROTATE_SPEED * dt)
	} else if inputs.KeyE.IsDown() {
		g.Screen.Rotate(SCREEN_FLIGHT_ROTATE_SPEED * dt)
	}

	return nil
})
