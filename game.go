package sk

import (
	"fmt"
	"log"

	"github.com/renatopp/skald/assets"
	"github.com/renatopp/skald/render"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	World    *World
	Assets   *assets.AssetServer
	Renderer *render.Renderer
	Total    int

	systems []*systemEntry
}

func NewGame() *Game {
	return &Game{
		World:    NewWorld(),
		Assets:   assets.NewAssetServer(),
		Renderer: render.NewRenderer(),

		systems: make([]*systemEntry, 0),
	}
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

func (g *Game) Play() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&runtime{Game: g}); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) draw(screen *ebiten.Image) {
	g.Renderer.Draw(screen)

	msg := fmt.Sprintf("TPS: %0.2f;\nFPS: %0.2f\nTotal: %d", ebiten.ActualTPS(), ebiten.ActualFPS(), g.Total)
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) tick() error {
	g.Renderer.Clear()

	buffer := make([]*systemEntry, 0)

	ebiten.SetVsyncEnabled(false)
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
