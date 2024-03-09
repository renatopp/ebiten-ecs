package main

import sk "github.com/renatopp/skald"

func main() {
	game := sk.NewGame()

	// paddle := sk.NewEntity(Transform, Sprite, Paddle, Collision)
	// ball := sk.NewEntity(Transform, Sprite, Collision)

	game.Play()
}
