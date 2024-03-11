package inputs

import "github.com/hajimehoshi/ebiten/v2"

type mouseCursor struct {
	lastX, lastY, x, y int
}

func (m *mouseCursor) Position() (x, y int) {
	return m.x, m.y
}

func (m *mouseCursor) HasMoved() bool {
	return m.x != m.lastX || m.y != m.lastY
}

func (m *mouseCursor) Moved() (x, y int) {
	return m.x - m.lastX, m.y - m.lastY
}

func (m *mouseCursor) update() {
	m.lastX, m.lastY = m.x, m.y
	m.x, m.y = ebiten.CursorPosition()
	if m.x == 0 && m.y == 0 {
		m.lastX, m.lastY = m.x, m.y
	}
}

var MouseCursor = &mouseCursor{}
