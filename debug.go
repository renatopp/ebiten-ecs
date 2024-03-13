package sk

import "github.com/hajimehoshi/ebiten/v2"

type Debug struct {
	Enabled bool

	beforeDraws []func(*ebiten.Image)
	afterDraws  []func(*ebiten.Image)
}

func NewDebug() *Debug {
	return &Debug{
		Enabled: true,
	}
}

func (d *Debug) BeforeDraw(fn func(*ebiten.Image)) {
	if !d.Enabled {
		return
	}

	d.beforeDraws = append(d.beforeDraws, fn)
}

func (d *Debug) AfterDraw(fn func(*ebiten.Image)) {
	if !d.Enabled {
		return
	}

	d.afterDraws = append(d.afterDraws, fn)
}

func (d *Debug) execBeforeDraw(screen *ebiten.Image) {
	for _, fn := range d.beforeDraws {
		fn(screen)
	}
}

func (d *Debug) execAfterDraw(screen *ebiten.Image) {
	for _, fn := range d.afterDraws {
		fn(screen)
	}
}
