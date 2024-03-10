package inputs

import "github.com/hajimehoshi/ebiten/v2"

type mouseWheel struct {
	axis     int // 0=x, 1=y
	previous float64
	current  float64
}

func (m *mouseWheel) Name() string {
	return "MouseWheel"
}

func (m *mouseWheel) IsDown() bool {
	return m.current != 0
}

func (m *mouseWheel) IsUp() bool {
	return m.current == 0
}

func (m *mouseWheel) IsPressed() bool {
	return m.current != 0 && m.previous == 0
}

func (m *mouseWheel) IsReleased() bool {
	return m.current == 0 && m.previous != 0
}

func (m *mouseWheel) GetBool() bool {
	return m.current != 0
}

func (m *mouseWheel) GetFloat() float64 {
	return float64(m.current)
}

func (m *mouseWheel) update() {
	m.previous = m.current
	x, y := ebiten.Wheel()
	if m.axis == 0 {
		m.current = x
	} else {
		m.current = y
	}
}

var MouseWheelX = &mouseWheel{axis: 0}
var MouseWheelY = &mouseWheel{axis: 1}
