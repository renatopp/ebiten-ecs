package sk

type InputDevice string

var InputDeviceMouse InputDevice = "mouse"
var InputDeviceKeyboard InputDevice = "keyboard"
var InputDeviceGamepad InputDevice = "gamepad"
var InputDeviceTouch InputDevice = "touch"

type InputKey interface {
	Code() int
	Name() string
	Device() InputDevice

	IsDown() bool
	IsUp() bool
	IsPressed() bool
	IsReleased() bool
	GetBool() bool
	GetFloat() float32

	update()
}
