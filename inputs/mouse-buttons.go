package inputs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type _mouseButtons struct {
	Left    ITrigger
	Middle  ITrigger
	Right   ITrigger
	Back    ITrigger
	Forward ITrigger

	Button0 ITrigger
	Button1 ITrigger
	Button2 ITrigger
	Button3 ITrigger
	Button4 ITrigger

	Any ITrigger
	All ITrigger
}

func newMouseButtons() *_mouseButtons {
	buttons := &_mouseButtons{}
	buttons.Left = NewMouseButton(ebiten.MouseButton0)
	buttons.Middle = NewMouseButton(ebiten.MouseButton1)
	buttons.Right = NewMouseButton(ebiten.MouseButton2)
	buttons.Back = NewMouseButton(ebiten.MouseButton3)
	buttons.Forward = NewMouseButton(ebiten.MouseButton4)

	buttons.Button0 = buttons.Left
	buttons.Button1 = buttons.Right
	buttons.Button2 = buttons.Middle
	buttons.Button3 = buttons.Back
	buttons.Button4 = buttons.Forward

	buttons.Any = Any{buttons.Left, buttons.Middle, buttons.Right, buttons.Back, buttons.Forward}
	buttons.All = All{buttons.Left, buttons.Middle, buttons.Right, buttons.Back, buttons.Forward}
	return buttons
}

var MouseButtons = newMouseButtons()
