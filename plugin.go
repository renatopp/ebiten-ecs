package sk

type IPlugin interface {
	Register(g *Game)
	Unregister(g *Game)
}

type Plugin struct {
	// Systems []ISystem
}
