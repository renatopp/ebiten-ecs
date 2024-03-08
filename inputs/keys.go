package inputs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type _keys struct {
	A              ITrigger
	B              ITrigger
	C              ITrigger
	D              ITrigger
	E              ITrigger
	F              ITrigger
	G              ITrigger
	H              ITrigger
	I              ITrigger
	J              ITrigger
	K              ITrigger
	L              ITrigger
	M              ITrigger
	N              ITrigger
	O              ITrigger
	P              ITrigger
	Q              ITrigger
	R              ITrigger
	S              ITrigger
	T              ITrigger
	U              ITrigger
	V              ITrigger
	W              ITrigger
	X              ITrigger
	Y              ITrigger
	Z              ITrigger
	AltLeft        ITrigger
	AltRight       ITrigger
	ArrowDown      ITrigger
	ArrowLeft      ITrigger
	ArrowRight     ITrigger
	ArrowUp        ITrigger
	Backquote      ITrigger
	Backslash      ITrigger
	Backspace      ITrigger
	BracketLeft    ITrigger
	BracketRight   ITrigger
	CapsLock       ITrigger
	Comma          ITrigger
	ContextMenu    ITrigger
	ControlLeft    ITrigger
	ControlRight   ITrigger
	Control        ITrigger
	Delete         ITrigger
	Num0           ITrigger
	Num1           ITrigger
	Num2           ITrigger
	Num3           ITrigger
	Num4           ITrigger
	Num5           ITrigger
	Num6           ITrigger
	Num7           ITrigger
	Num8           ITrigger
	Num9           ITrigger
	End            ITrigger
	Enter          ITrigger
	Equal          ITrigger
	Escape         ITrigger
	F1             ITrigger
	F2             ITrigger
	F3             ITrigger
	F4             ITrigger
	F5             ITrigger
	F6             ITrigger
	F7             ITrigger
	F8             ITrigger
	F9             ITrigger
	F10            ITrigger
	F11            ITrigger
	F12            ITrigger
	Home           ITrigger
	Insert         ITrigger
	MetaLeft       ITrigger
	MetaRight      ITrigger
	Minus          ITrigger
	NumLock        ITrigger
	Numpad0        ITrigger
	Numpad1        ITrigger
	Numpad2        ITrigger
	Numpad3        ITrigger
	Numpad4        ITrigger
	Numpad5        ITrigger
	Numpad6        ITrigger
	Numpad7        ITrigger
	Numpad8        ITrigger
	Numpad9        ITrigger
	NumpadAdd      ITrigger
	NumpadDecimal  ITrigger
	NumpadDivide   ITrigger
	NumpadEnter    ITrigger
	NumpadEqual    ITrigger
	NumpadMultiply ITrigger
	NumpadSubtract ITrigger
	PageDown       ITrigger
	PageUp         ITrigger
	Pause          ITrigger
	Period         ITrigger
	PrintScreen    ITrigger
	Quote          ITrigger
	ScrollLock     ITrigger
	Semicolon      ITrigger
	ShiftLeft      ITrigger
	ShiftRight     ITrigger
	Slash          ITrigger
	Space          ITrigger
	Tab            ITrigger
	CtrlLeft       ITrigger
	CtrlRight      ITrigger
	Ctrl           ITrigger
	Shift          ITrigger
	Alt            ITrigger
	Meta           ITrigger
	Any            ITrigger
	All            ITrigger
}

func newKeys() *_keys {
	keys := &_keys{}
	keys.A = NewKey(ebiten.KeyA)
	keys.B = NewKey(ebiten.KeyB)
	keys.C = NewKey(ebiten.KeyC)
	keys.D = NewKey(ebiten.KeyD)
	keys.E = NewKey(ebiten.KeyE)
	keys.F = NewKey(ebiten.KeyF)
	keys.G = NewKey(ebiten.KeyG)
	keys.H = NewKey(ebiten.KeyH)
	keys.I = NewKey(ebiten.KeyI)
	keys.J = NewKey(ebiten.KeyJ)
	keys.K = NewKey(ebiten.KeyK)
	keys.L = NewKey(ebiten.KeyL)
	keys.M = NewKey(ebiten.KeyM)
	keys.N = NewKey(ebiten.KeyN)
	keys.O = NewKey(ebiten.KeyO)
	keys.P = NewKey(ebiten.KeyP)
	keys.Q = NewKey(ebiten.KeyQ)
	keys.R = NewKey(ebiten.KeyR)
	keys.S = NewKey(ebiten.KeyS)
	keys.T = NewKey(ebiten.KeyT)
	keys.U = NewKey(ebiten.KeyU)
	keys.V = NewKey(ebiten.KeyV)
	keys.W = NewKey(ebiten.KeyW)
	keys.X = NewKey(ebiten.KeyX)
	keys.Y = NewKey(ebiten.KeyY)
	keys.Z = NewKey(ebiten.KeyZ)
	keys.AltLeft = NewKey(ebiten.KeyAltLeft)
	keys.AltRight = NewKey(ebiten.KeyAltRight)
	keys.ArrowDown = NewKey(ebiten.KeyArrowDown)
	keys.ArrowLeft = NewKey(ebiten.KeyArrowLeft)
	keys.ArrowRight = NewKey(ebiten.KeyArrowRight)
	keys.ArrowUp = NewKey(ebiten.KeyArrowUp)
	keys.Backquote = NewKey(ebiten.KeyBackquote)
	keys.Backslash = NewKey(ebiten.KeyBackslash)
	keys.Backspace = NewKey(ebiten.KeyBackspace)
	keys.BracketLeft = NewKey(ebiten.KeyBracketLeft)
	keys.BracketRight = NewKey(ebiten.KeyBracketRight)
	keys.CapsLock = NewKey(ebiten.KeyCapsLock)
	keys.Comma = NewKey(ebiten.KeyComma)
	keys.ContextMenu = NewKey(ebiten.KeyContextMenu)
	keys.ControlLeft = NewKey(ebiten.KeyControlLeft)
	keys.ControlRight = NewKey(ebiten.KeyControlRight)
	keys.Delete = NewKey(ebiten.KeyDelete)
	keys.Num0 = NewKey(ebiten.KeyDigit0)
	keys.Num1 = NewKey(ebiten.KeyDigit1)
	keys.Num2 = NewKey(ebiten.KeyDigit2)
	keys.Num3 = NewKey(ebiten.KeyDigit3)
	keys.Num4 = NewKey(ebiten.KeyDigit4)
	keys.Num5 = NewKey(ebiten.KeyDigit5)
	keys.Num6 = NewKey(ebiten.KeyDigit6)
	keys.Num7 = NewKey(ebiten.KeyDigit7)
	keys.Num8 = NewKey(ebiten.KeyDigit8)
	keys.Num9 = NewKey(ebiten.KeyDigit9)
	keys.End = NewKey(ebiten.KeyEnd)
	keys.Enter = NewKey(ebiten.KeyEnter)
	keys.Equal = NewKey(ebiten.KeyEqual)
	keys.Escape = NewKey(ebiten.KeyEscape)
	keys.F1 = NewKey(ebiten.KeyF1)
	keys.F2 = NewKey(ebiten.KeyF2)
	keys.F3 = NewKey(ebiten.KeyF3)
	keys.F4 = NewKey(ebiten.KeyF4)
	keys.F5 = NewKey(ebiten.KeyF5)
	keys.F6 = NewKey(ebiten.KeyF6)
	keys.F7 = NewKey(ebiten.KeyF7)
	keys.F8 = NewKey(ebiten.KeyF8)
	keys.F9 = NewKey(ebiten.KeyF9)
	keys.F10 = NewKey(ebiten.KeyF10)
	keys.F11 = NewKey(ebiten.KeyF11)
	keys.F12 = NewKey(ebiten.KeyF12)
	keys.Home = NewKey(ebiten.KeyHome)
	keys.Insert = NewKey(ebiten.KeyInsert)
	keys.MetaLeft = NewKey(ebiten.KeyMetaLeft)
	keys.MetaRight = NewKey(ebiten.KeyMetaRight)
	keys.Minus = NewKey(ebiten.KeyMinus)
	keys.NumLock = NewKey(ebiten.KeyNumLock)
	keys.Numpad0 = NewKey(ebiten.KeyNumpad0)
	keys.Numpad1 = NewKey(ebiten.KeyNumpad1)
	keys.Numpad2 = NewKey(ebiten.KeyNumpad2)
	keys.Numpad3 = NewKey(ebiten.KeyNumpad3)
	keys.Numpad4 = NewKey(ebiten.KeyNumpad4)
	keys.Numpad5 = NewKey(ebiten.KeyNumpad5)
	keys.Numpad6 = NewKey(ebiten.KeyNumpad6)
	keys.Numpad7 = NewKey(ebiten.KeyNumpad7)
	keys.Numpad8 = NewKey(ebiten.KeyNumpad8)
	keys.Numpad9 = NewKey(ebiten.KeyNumpad9)
	keys.NumpadAdd = NewKey(ebiten.KeyNumpadAdd)
	keys.NumpadDecimal = NewKey(ebiten.KeyNumpadDecimal)
	keys.NumpadDivide = NewKey(ebiten.KeyNumpadDivide)
	keys.NumpadEnter = NewKey(ebiten.KeyNumpadEnter)
	keys.NumpadEqual = NewKey(ebiten.KeyNumpadEqual)
	keys.NumpadMultiply = NewKey(ebiten.KeyNumpadMultiply)
	keys.NumpadSubtract = NewKey(ebiten.KeyNumpadSubtract)
	keys.PageDown = NewKey(ebiten.KeyPageDown)
	keys.PageUp = NewKey(ebiten.KeyPageUp)
	keys.Pause = NewKey(ebiten.KeyPause)
	keys.Period = NewKey(ebiten.KeyPeriod)
	keys.PrintScreen = NewKey(ebiten.KeyPrintScreen)
	keys.Quote = NewKey(ebiten.KeyQuote)
	keys.ScrollLock = NewKey(ebiten.KeyScrollLock)
	keys.Semicolon = NewKey(ebiten.KeySemicolon)
	keys.ShiftLeft = NewKey(ebiten.KeyShiftLeft)
	keys.ShiftRight = NewKey(ebiten.KeyShiftRight)
	keys.Slash = NewKey(ebiten.KeySlash)
	keys.Space = NewKey(ebiten.KeySpace)
	keys.Tab = NewKey(ebiten.KeyTab)
	keys.CtrlLeft = keys.ControlLeft
	keys.CtrlRight = keys.ControlRight
	keys.Control = Any{keys.CtrlLeft, keys.CtrlRight}
	keys.Ctrl = keys.Control
	keys.Shift = Any{keys.ShiftLeft, keys.ShiftRight}
	keys.Alt = Any{keys.AltLeft, keys.AltRight}
	keys.Meta = Any{keys.MetaLeft, keys.MetaRight}
	keys.Any = Any{keys.A, keys.B, keys.C, keys.D, keys.E, keys.F, keys.G, keys.H, keys.I, keys.J, keys.K, keys.L, keys.M, keys.N, keys.O, keys.P, keys.Q, keys.R, keys.S, keys.T, keys.U, keys.V, keys.W, keys.X, keys.Y, keys.Z, keys.AltLeft, keys.AltRight, keys.ArrowDown, keys.ArrowLeft, keys.ArrowRight, keys.ArrowUp, keys.Backquote, keys.Backslash, keys.Backspace, keys.BracketLeft, keys.BracketRight, keys.CapsLock, keys.Comma, keys.ContextMenu, keys.ControlLeft, keys.ControlRight, keys.Delete, keys.Num0, keys.Num1, keys.Num2, keys.Num3, keys.Num4, keys.Num5, keys.Num6, keys.Num7, keys.Num8, keys.Num9, keys.End, keys.Enter, keys.Equal, keys.Escape, keys.F1, keys.F2, keys.F3, keys.F4, keys.F5, keys.F6, keys.F7, keys.F8, keys.F9, keys.F10, keys.F11, keys.F12, keys.Home, keys.Insert, keys.MetaLeft, keys.MetaRight, keys.Minus, keys.NumLock, keys.Numpad0, keys.Numpad1, keys.Numpad2, keys.Numpad3, keys.Numpad4, keys.Numpad5, keys.Numpad6, keys.Numpad7, keys.Numpad8, keys.Numpad9, keys.NumpadAdd, keys.NumpadDecimal, keys.NumpadDivide, keys.NumpadEnter, keys.NumpadEqual, keys.NumpadMultiply, keys.NumpadSubtract, keys.PageDown, keys.PageUp, keys.Pause, keys.Period, keys.PrintScreen, keys.Quote, keys.ScrollLock, keys.Semicolon, keys.ShiftLeft, keys.ShiftRight, keys.Slash, keys.Space, keys.Tab}
	keys.All = All{keys.A, keys.B, keys.C, keys.D, keys.E, keys.F, keys.G, keys.H, keys.I, keys.J, keys.K, keys.L, keys.M, keys.N, keys.O, keys.P, keys.Q, keys.R, keys.S, keys.T, keys.U, keys.V, keys.W, keys.X, keys.Y, keys.Z, keys.AltLeft, keys.AltRight, keys.ArrowDown, keys.ArrowLeft, keys.ArrowRight, keys.ArrowUp, keys.Backquote, keys.Backslash, keys.Backspace, keys.BracketLeft, keys.BracketRight, keys.CapsLock, keys.Comma, keys.ContextMenu, keys.ControlLeft, keys.ControlRight, keys.Delete, keys.Num0, keys.Num1, keys.Num2, keys.Num3, keys.Num4, keys.Num5, keys.Num6, keys.Num7, keys.Num8, keys.Num9, keys.End, keys.Enter, keys.Equal, keys.Escape, keys.F1, keys.F2, keys.F3, keys.F4, keys.F5, keys.F6, keys.F7, keys.F8, keys.F9, keys.F10, keys.F11, keys.F12, keys.Home, keys.Insert, keys.MetaLeft, keys.MetaRight, keys.Minus, keys.NumLock, keys.Numpad0, keys.Numpad1, keys.Numpad2, keys.Numpad3, keys.Numpad4, keys.Numpad5, keys.Numpad6, keys.Numpad7, keys.Numpad8, keys.Numpad9, keys.NumpadAdd, keys.NumpadDecimal, keys.NumpadDivide, keys.NumpadEnter, keys.NumpadEqual, keys.NumpadMultiply, keys.NumpadSubtract, keys.PageDown, keys.PageUp, keys.Pause, keys.Period, keys.PrintScreen, keys.Quote, keys.ScrollLock, keys.Semicolon, keys.ShiftLeft, keys.ShiftRight, keys.Slash, keys.Space, keys.Tab}
	return keys
}

var Keys = newKeys()
