package sk

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/renatopp/skald/inputs"
)

type Game struct {
	World    *World
	Assets   *AssetServer
	Renderer *Renderer
	Inputs   *inputs.System
	Window   *Window
	Screen   *Screen

	systems  []*systemEntry
	services map[ID]interface{}

	totalEntities int // TODO: remove me
}

func NewGame() *Game {
	ebiten.SetVsyncEnabled(false)

	g := &Game{
		World:    NewWorld(),
		Assets:   NewAssetServer(),
		Renderer: NewRenderer(),
		Inputs:   inputs.NewSystem(),
		Window:   NewWindow(),
		Screen:   NewScreen(),

		systems:  make([]*systemEntry, 0),
		services: make(map[ID]interface{}),
	}

	return g
}

func (g *Game) AddSystem(system *SystemDefinition, options ...*SystemOptions) {
	for _, entry := range g.systems {
		if entry.definition == system {
			return
		}
	}

	opt := NewSystemOptions()
	if len(options) > 0 {
		opt = options[0]
	}

	for _, query := range system.queries {
		g.World.AddQuery(query)
	}

	g.systems = append(g.systems, &systemEntry{
		definition: system,
		options:    opt,
	})
}

func (g *Game) RemoveSystem(system *SystemDefinition) {
	for i, entry := range g.systems {
		if entry.definition == system {
			for _, query := range system.queries {
				g.World.RemoveQuery(query)
			}

			g.systems = append(g.systems[:i], g.systems[i+1:]...)
			break
		}
	}
}

func (g *Game) AddService(service IService) {
	if _, ok := g.services[service.Id()]; ok {
		return
	}

	g.services[service.Id()] = service.New()
}

func (g *Game) RemoveService(service IService) {
	delete(g.services, service.Id())
}

func (g *Game) Play() {
	if err := ebiten.RunGameWithOptions(&runtime{Game: g}, &ebiten.RunGameOptions{
		InitUnfocused: false,
	}); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-g.Screen.position.X*g.Screen.PixelsPerUnit, -g.Screen.position.Y*g.Screen.PixelsPerUnit)

	screen.Fill(g.Screen.color)
	g.Renderer.Draw(screen, op)

	// TODO: Remove me
	msg := fmt.Sprintf("TPS: %0.2f;\nFPS: %0.2f\nTotal: %d", ebiten.ActualTPS(), ebiten.ActualFPS(), g.totalEntities)
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) tick() error {
	g.Renderer.Clear()
	g.Inputs.Update()

	buffer := make([]*systemEntry, 0)

	for _, entry := range g.systems {
		if !entry.options.once {
			buffer = append(buffer, entry)
		}

		if err := entry.definition.fn(g); err != nil {
			return err
		}
	}

	g.systems = buffer
	return nil
}
