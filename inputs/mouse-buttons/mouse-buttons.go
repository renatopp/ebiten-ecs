package mouseButtons

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/renatopp/skald/inputs"
)

var Left = inputs.NewMouseButton(ebiten.MouseButton0)
var Middle = inputs.NewMouseButton(ebiten.MouseButton1)
var Right = inputs.NewMouseButton(ebiten.MouseButton2)
var Back = inputs.NewMouseButton(ebiten.MouseButton3)
var Forward = inputs.NewMouseButton(ebiten.MouseButton4)

var Button0 = Left
var Button1 = Right
var Button2 = Middle
var Button3 = Back
var Button4 = Forward

var Any = inputs.Any{Left, Middle, Right, Back, Forward}
var All = inputs.All{Left, Middle, Right, Back, Forward}
