package sk

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Touch struct {
	id        ebiten.TouchID
	prevValue bool
	curValue  bool
}

func (k *Touch) Code() int {
	return 0
}

func (k *Touch) Name() string {
	return string(0)
}

func (k *Touch) Device() InputDevice {
	return InputDeviceMouse
}

func (k *Touch) IsDown() bool {
	return k.curValue
}

func (k *Touch) IsUp() bool {
	return !k.curValue
}

func (k *Touch) IsPressed() bool {
	return k.curValue && !k.prevValue
}

func (k *Touch) IsReleased() bool {
	return !k.curValue && k.prevValue
}

func (k *Touch) GetBool() bool {
	return k.curValue
}

func (k *Touch) GetFloat() float32 {
	if k.curValue {
		return 1
	}
	return 0
}

func (k *Touch) update() {
	k.prevValue = k.curValue

	// if inpututil.Touch(k.id) {
	// 	k.curValue = true
	// } else if inpututil.IsTouchJustReleased(k.id) {
	// 	k.curValue = false
	// }
}
