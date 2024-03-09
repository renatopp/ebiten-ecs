package inputs

type Device string

var DeviceMouse Device = "mouse"
var DeviceKeyboard Device = "keyboard"
var DeviceGamepad Device = "gamepad"
var DeviceTouch Device = "touch"

type ITrigger interface {
	Name() string
	IsDown() bool
	IsUp() bool
	IsPressed() bool
	IsReleased() bool
	GetBool() bool
	GetFloat() float32

	update()
}

type ICursor interface {
}

// -----------------------------------------------------------------------------
// INPUT SYSTEM
// -----------------------------------------------------------------------------
type System struct {
}

func NewSystem() *System {
	return &System{}
}

func (i *System) Update() {
	KeyAll.update()
	MouseButtonAll.update()
}

// -----------------------------------------------------------------------------
// INPUT ANY
// -----------------------------------------------------------------------------
type Any []ITrigger

func (any Any) Name() string {
	return "Any"
}

func (any Any) IsDown() bool {
	for _, trigger := range any {
		if trigger.IsDown() {
			return true
		}
	}
	return false
}

func (any Any) IsUp() bool {
	for _, trigger := range any {
		if trigger.IsUp() {
			return true
		}
	}
	return false
}

func (any Any) IsPressed() bool {
	for _, trigger := range any {
		if trigger.IsPressed() {
			return true
		}
	}
	return false
}

func (any Any) IsReleased() bool {
	for _, trigger := range any {
		if trigger.IsReleased() {
			return true
		}
	}
	return false
}

func (any Any) GetBool() bool {
	for _, trigger := range any {
		if trigger.GetBool() {
			return true
		}
	}
	return false

}

func (any Any) GetFloat() float32 {
	for _, trigger := range any {
		f := trigger.GetFloat()
		if f != 0 {
			return f
		}
	}
	return 0
}

func (any Any) update() {
	for _, trigger := range any {
		trigger.update()
	}
}

// -----------------------------------------------------------------------------
// INPUT ALL
// -----------------------------------------------------------------------------
type All []ITrigger

func (all All) Name() string {
	return "All"
}

func (all All) IsDown() bool {
	for _, trigger := range all {
		if !trigger.IsDown() {
			return false
		}
	}
	return true
}

func (all All) IsUp() bool {
	for _, trigger := range all {
		if !trigger.IsUp() {
			return false
		}
	}
	return true
}

func (all All) IsPressed() bool {
	pressed := false

	for _, trigger := range all {
		if !trigger.IsDown() {
			return false
		}

		if trigger.IsPressed() {
			pressed = true
		}
	}

	return pressed
}

func (all All) IsReleased() bool {
	released := false

	for _, trigger := range all {
		if !trigger.IsUp() {
			return false
		}

		if trigger.IsReleased() {
			released = true
		}
	}

	return released
}

func (all All) GetBool() bool {
	for _, trigger := range all {
		if !trigger.GetBool() {
			return false
		}
	}
	return true
}

func (all All) GetFloat() float32 {
	for _, trigger := range all {
		f := trigger.GetFloat()
		if f == 0 {
			return 0
		}
	}
	return 1
}

func (all All) update() {
	for _, trigger := range all {
		trigger.update()
	}
}
