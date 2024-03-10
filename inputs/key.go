package inputs

import "github.com/hajimehoshi/ebiten/v2"

type Key struct {
	code ebiten.Key

	prevValue bool
	curValue  bool
}

func NewKey(code ebiten.Key) *Key {
	return &Key{
		code: code,
	}
}

func (k *Key) Code() int {
	return int(k.code)
}

func (k *Key) Name() string {
	return ebiten.KeyName(k.code)
}

func (k *Key) Device() Device {
	return DeviceKeyboard
}

func (k *Key) IsDown() bool {
	return k.curValue
}

func (k *Key) IsUp() bool {
	return !k.curValue
}

func (k *Key) IsPressed() bool {
	return k.curValue && !k.prevValue
}

func (k *Key) IsReleased() bool {
	return !k.curValue && k.prevValue
}

func (k *Key) GetBool() bool {
	return k.curValue
}

func (k *Key) GetFloat() float64 {
	if k.curValue {
		return 1
	}
	return 0
}

func (k *Key) update() {
	k.prevValue = k.curValue
	k.curValue = ebiten.IsKeyPressed(k.code)
}
