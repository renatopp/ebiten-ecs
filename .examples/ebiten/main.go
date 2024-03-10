package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const WINDOW_WIDTH = 800
const WINDOW_HEIGHT = 600
const SCREEN_WIDTH = 2000
const SCREEN_HEIGHT = 1000

type Game struct {
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, 0, SCREEN_WIDTH, SCREEN_HEIGHT, color.RGBA{0x80, 0xa0, 0xc0, 0xff}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetCursorMode(ebiten.CursorModeCaptured)
	ebiten.RunGame(NewGame())
}
