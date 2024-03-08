package inputs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var MouseButtonLeft = NewMouseButton(ebiten.MouseButton0)
var MouseButtonMiddle = NewMouseButton(ebiten.MouseButton1)
var MouseButtonRight = NewMouseButton(ebiten.MouseButton2)
var MouseButtonBack = NewMouseButton(ebiten.MouseButton3)
var MouseButtonForward = NewMouseButton(ebiten.MouseButton4)

var MouseButton0 = MouseButtonLeft
var MouseButton1 = MouseButtonRight
var MouseButton2 = MouseButtonMiddle
var MouseButton3 = MouseButtonBack
var MouseButton4 = MouseButtonForward

var MouseButtonAny = Any{MouseButtonLeft, MouseButtonMiddle, MouseButtonRight, MouseButtonBack, MouseButtonForward}
var MouseButtonAll = All{MouseButtonLeft, MouseButtonMiddle, MouseButtonRight, MouseButtonBack, MouseButtonForward}
