package sk

var nextTimerId ID = 0

type Timer struct {
	id        ID
	deltaTime float64
	scale     float64
	paused    bool

	attached map[ID]*Timer
}

func NewTimer() *Timer {
	t := &Timer{
		id:        nextTimerId,
		deltaTime: 0,
		scale:     1,
		paused:    false,
		attached:  make(map[ID]*Timer),
	}
	nextTimerId++

	return t
}

// Delta time in seconds
func (t *Timer) Update(dt float64) {
	if t.paused {
		t.deltaTime = 0
		return
	}

	t.deltaTime = dt * t.scale

	for _, other := range t.attached {
		other.Update(dt)
	}
}

func (t *Timer) GetDeltaTime() float64 {
	return t.deltaTime
}

func (t *Timer) IsPaused() bool {
	return t.paused
}

func (t *Timer) Pause() {
	t.paused = true
}

func (t *Timer) Resume() {
	t.paused = false
}

func (t *Timer) GetScale() float64 {
	return t.scale
}

func (t *Timer) SetScale(scale float64) {
	t.scale = scale
}

func (t *Timer) Attach(other *Timer) {
	t.attached[other.id] = other
}

func (t *Timer) Detach(other *Timer) {
	delete(t.attached, other.id)
}
