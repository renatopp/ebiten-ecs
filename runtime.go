package sk

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type runtime struct {
	Game *Game
}

func (g *runtime) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *runtime) Draw(screen *ebiten.Image) {
	g.Game.draw(screen)
}

func (g *runtime) Update() error {
	return g.Game.tick()
}
