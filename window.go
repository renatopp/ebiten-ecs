package sk

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Window struct {
	title          string
	cursorCaptured bool
	cursorVisible  bool
}

func NewWindow() *Window {
	win := &Window{
		title: "Skald Game",
	}

	win.SetTitle(win.title)
	win.SetResizable(true)
	win.SetVsync(false)
	win.SetSize(WINDOW_WIDTH, WINDOW_HEIGHT)

	return win
}

func (win *Window) GetMonitorSize() (int, int) {
	return ebiten.ScreenSizeInFullscreen()
}

func (win *Window) GetSize() (int, int) {
	return ebiten.WindowSize()
}

func (win *Window) SetSize(w, h int) {
	ebiten.SetWindowSize(w, h)
}

func (win *Window) GetSizeLimits() (minW, minH, maxW, maxH int) {
	return ebiten.WindowSizeLimits()
}

func (win *Window) SetSizeLimits(minW, minH, maxW, maxH int) {
	ebiten.SetWindowSizeLimits(minW, minH, maxW, maxH)
}

func (win *Window) GetPosition() (int, int) {
	return ebiten.WindowPosition()
}

func (win *Window) SetPosition(x, y int) {
	ebiten.SetWindowPosition(x, y)
}

func (win *Window) GetTitle() string {
	return win.title
}

func (win *Window) SetTitle(title string) {
	win.title = title
	ebiten.SetWindowTitle(title)
}

func (win *Window) IsResizable() bool {
	return ebiten.WindowResizingMode() == ebiten.WindowResizingModeEnabled
}

func (win *Window) SetResizable(resizable bool) {
	if resizable {
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	} else {
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	}
}

func (win *Window) IsFullscreen() bool {
	return ebiten.IsFullscreen()
}

func (win *Window) SetFullscreen(fullscreen bool) {
	ebiten.SetFullscreen(fullscreen)
}

func (win *Window) IsFocused() bool {
	return ebiten.IsFocused()
}

func (win *Window) IsVsync() bool {
	return ebiten.IsVsyncEnabled()
}

func (win *Window) SetVsync(vsync bool) {
	ebiten.SetVsyncEnabled(vsync)
}

func (win *Window) IsCursorVisible() bool {
	return ebiten.CursorMode() == ebiten.CursorModeVisible
}

func (win *Window) SetCursorVisible(visible bool) {
	if win.cursorCaptured {
		return
	}

	win.cursorVisible = visible
	if visible {
		ebiten.SetCursorMode(ebiten.CursorModeVisible)
	} else {
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
	}
}

func (win *Window) IsCursorCaptured() bool {
	return ebiten.CursorMode() == ebiten.CursorModeCaptured
}

func (win *Window) SetCursorCaptured(captured bool) {
	win.cursorCaptured = captured

	if captured {
		ebiten.SetCursorMode(ebiten.CursorModeCaptured)
	} else {
		win.SetCursorVisible(win.cursorVisible)
	}
}

func (win *Window) GetCursorShape() ebiten.CursorShapeType {
	return ebiten.CursorShape()
}

func (win *Window) SetCursorShape(shape ebiten.CursorShapeType) {
	ebiten.SetCursorShape(shape)
}

func (win *Window) Close() {
	os.Exit(0)
}

func (win *Window) IsMaximized() bool {
	return ebiten.IsWindowMaximized()
}

func (win *Window) Maximize() {
	ebiten.MaximizeWindow()
}

func (win *Window) IsMinimized() bool {
	return ebiten.IsWindowMinimized()
}

func (win *Window) Minimize() {
	ebiten.MinimizeWindow()
}

func (win *Window) Restore() {
	ebiten.RestoreWindow()
}

func (win *Window) GetFPS() float64 {
	return ebiten.ActualFPS()
}

func (win *Window) GetTPS() float64 {
	return ebiten.ActualTPS()
}

func (win *Window) SetTPS(tps int) {
	ebiten.SetTPS(tps)
}

func (win *Window) IsMousePassthrough() bool {
	return ebiten.IsWindowMousePassthrough()
}

func (win *Window) SetMousePassthrough(passthrough bool) {
	ebiten.SetWindowMousePassthrough(passthrough)
}

func (win *Window) IsDecorated() bool {
	return ebiten.IsWindowDecorated()
}

func (win *Window) SetDecorated(decorated bool) {
	ebiten.SetWindowDecorated(decorated)
}

func (win *Window) IsFloating() bool {
	return ebiten.IsWindowFloating()
}

func (win *Window) SetFloating(floating bool) {
	ebiten.SetWindowFloating(floating)
}

func (win *Window) IsRunnableOnUnfocused() bool {
	return ebiten.IsRunnableOnUnfocused()
}

func (win *Window) SetRunnableOnUnfocused(runnable bool) {
	ebiten.SetRunnableOnUnfocused(runnable)
}
