package inputs

import "github.com/hajimehoshi/ebiten/v2"

type MouseButton struct {
	code ebiten.MouseButton

	prevValue bool
	curValue  bool
}

func NewMouseButton(code ebiten.MouseButton) *MouseButton {
	return &MouseButton{
		code: code,
	}
}

func (k *MouseButton) Code() int {
	return int(k.code)
}

func (k *MouseButton) Name() string {
	return string(k.code)
}

func (k *MouseButton) Device() Device {
	return DeviceMouse
}

func (k *MouseButton) IsDown() bool {
	return k.curValue
}

func (k *MouseButton) IsUp() bool {
	return !k.curValue
}

func (k *MouseButton) IsPressed() bool {
	return k.curValue && !k.prevValue
}

func (k *MouseButton) IsReleased() bool {
	return !k.curValue && k.prevValue
}

func (k *MouseButton) GetBool() bool {
	return k.curValue
}

func (k *MouseButton) GetFloat() float64 {
	if k.curValue {
		return 1
	}
	return 0
}

func (k *MouseButton) update() {
	k.prevValue = k.curValue
	k.curValue = ebiten.IsMouseButtonPressed(k.code)
}
