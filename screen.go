package sk

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/renatopp/skald/utils"
)

type Screen struct {
	Surface       *ebiten.Image
	PixelsPerUnit float64

	position Vec2 // In units
	size     Vec2 // In units
	rotation float64
	color    color.Color

	zoom          float64 // internal value, from 0 to steps
	zoomSteps     float64 // how many steps for a mouse wheel to go from min to max zoom
	minZoom       float64
	maxZoom       float64
	effectiveZoom float64

	viewportWidth  int
	viewportHeight int
	viewportSize   Vec2
}

func NewScreen() *Screen {
	s := &Screen{
		Surface:       ebiten.NewImage(1, 1),
		PixelsPerUnit: PIXELS_PER_UNIT,
		position:      Vec2{X: 0, Y: 0},
		size:          Vec2{X: SCREEN_WIDTH, Y: SCREEN_HEIGHT},
		zoom:          6,
		zoomSteps:     20,
		effectiveZoom: 0,
		rotation:      0,
		color:         color.White,
		minZoom:       0.1,
		maxZoom:       10,
	}

	s.SetColor(color.RGBA{110, 164, 191, 255})
	s.resize()

	return s
}

func (s *Screen) GetPosition() Vec2 {
	return s.position
}

func (s *Screen) SetPosition(x, y float64) {
	s.position = Vec2{X: x, Y: y}
}

func (s *Screen) Move(x, y float64) {
	s.position.X += x
	s.position.Y += y
}

func (s *Screen) GetSize() Vec2 {
	return s.size
}

func (s *Screen) SetSize(w, h float64) {
	s.size = Vec2{X: w, Y: h}
	s.resize()
}

func (s *Screen) GetZoom() float64 {
	return s.zoom
}

func (s *Screen) SetZoom(zoom float64) {
	s.zoom = zoom
	s.zoom = utils.Clamp(zoom, 0, s.zoomSteps)
	s.resize()
}

func (s *Screen) Zoom(zoom float64) {
	s.zoom = utils.Clamp(s.zoom+zoom, 0, s.zoomSteps)
	s.resize()
}

func (s *Screen) GetZoomLimits() (float64, float64) {
	return s.minZoom, s.maxZoom
}

func (s *Screen) SetZoomLimits(min, max float64) {
	s.minZoom = min
	s.maxZoom = max
	s.SetZoom(s.zoom)
}

func (s *Screen) GetRotation() float64 {
	return s.rotation
}

func (s *Screen) SetRotation(rotation float64) {
	s.rotation = rotation
}

func (s *Screen) Rotate(rotation float64) {
	s.rotation += rotation

	t := rotation * utils.Deg2Rad
	x2 := s.position.X*math.Cos(t) - s.position.Y*math.Sin(t)
	y2 := s.position.X*math.Sin(t) + s.position.Y*math.Cos(t)
	s.Move(s.position.X-x2, s.position.Y-y2)

}

func (s *Screen) GetColor() color.Color {
	return s.color
}

func (s *Screen) SetColor(color color.Color) {
	s.color = color
}

func (s *Screen) Clear() {
	s.Surface.Fill(s.color)
}

func (s *Screen) ScreenPositionToWorld(x, y int) (float64, float64) {
	return s.position.X, s.position.Y
}

func (s *Screen) WorldPositionToScreen(x, y float64) (int, int) {
	return int((x-s.position.X)*s.PixelsPerUnit + 0.5), int((y-s.position.Y)*s.PixelsPerUnit + 0.5)
}

func (s *Screen) resize() {
	s.effectiveZoom = utils.Lerp(s.minZoom, s.maxZoom, utils.EaseInQuad(s.zoom/s.zoomSteps))
	newW := s.size.X * 1.0 / s.effectiveZoom
	newH := s.size.Y * 1.0 / s.effectiveZoom
	if newW <= 16384 && newH <= 16384 {
		s.viewportWidth = int(newW * s.PixelsPerUnit)
		s.viewportHeight = int(newH * s.PixelsPerUnit)
		s.viewportSize = Vec2{X: newW, Y: newH}
		s.Surface.Dispose()
		s.Surface = ebiten.NewImage(s.viewportWidth, s.viewportHeight)
	}
}
