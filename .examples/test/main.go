package main

import sk "github.com/renatopp/skald"

func main() {
	game := sk.NewGame()

	as := AssetServer.Get(game)
	as.Load("sample")
}
