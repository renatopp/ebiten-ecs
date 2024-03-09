package main

import sk "github.com/renatopp/skald"

type DPaddle struct {
	Width    float64
	Height   float64
	Velocity float64
}

var Paddle = sk.NewComponent[DPaddle]()

type DBall struct {
	Radius   float64
	Velocity float64
}

var Ball = sk.NewComponent[DBall]()

func main() {
	game := sk.NewGame()

	// paddle := sk.NewEntity(Transform, Sprite, Paddle, Collision)
	// ball := sk.NewEntity(Transform, Sprite, Collision)

	game.Play()
}
