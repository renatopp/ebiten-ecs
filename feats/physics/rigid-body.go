package physics

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/mathf"
	"github.com/renatopp/skald/random"
)

var CIRCLE = 0
var BOX = 1

var MIN_BODY_SIZE float64 = 0.01 * 0.01
var MAX_BODY_SIZE float64 = 64 * 64
var MIN_DENSITY float64 = 0.01
var MAX_DENSITY float64 = 24

type transformer struct {
	X   float64
	Y   float64
	Sin float64
	Cos float64
}

// func newTransformer(x, y, angle float64) *transformer {
// 	return &transformer{
// 		X:   x,
// 		Y:   y,
// 		Sin: math.Sin(angle),
// 		Cos: math.Cos(angle),
// 	}
// }

type DRigidBody struct {
	LinearVelocity  sk.Vec2
	AngularVelocity float64

	Mass        float64
	Density     float64
	Restitution float64
	Area        float64

	Static bool

	// Temp
	Shape  int // 0: Circle, 1: Box
	Radius float64
	Width  float64
	Height float64

	image               *ebiten.Image
	vertices            []*sk.Vec2
	transformedVertices []*sk.Vec2
	transformers        []*transformer
	triangleIndices     []int
	transformVersion    uint64
}

func (r *DRigidBody) GetBounds() sk.Rect {
	if r.Shape == CIRCLE {
		return sk.Rect{
			Position: sk.Vec2{X: -r.Radius, Y: -r.Radius},
			Size:     sk.Vec2{X: r.Radius * 2, Y: r.Radius * 2},
		}
	} else {
		return sk.Rect{
			Position: sk.Vec2{X: -r.Width / 2, Y: -r.Height / 2},
			Size:     sk.Vec2{X: r.Width, Y: r.Height},
		}
	}
}

// Temp
func (r *DRigidBody) AsCircle(radius, density, restitution float64, static bool) {
	area := 2 * radius * math.Pi

	if area > MAX_BODY_SIZE {
		area = MAX_BODY_SIZE
		println("warn: body size too big")
	}

	if area < MIN_BODY_SIZE {
		area = MIN_BODY_SIZE
		println("warn: body size too small")
	}

	if density < MIN_DENSITY {
		density = MIN_DENSITY
		println("warn: density too small")
	}

	if density > MAX_DENSITY {
		density = MAX_DENSITY
		println("warn: density too big")
	}

	r.Shape = CIRCLE
	r.Radius = radius
	r.Area = area
	r.Density = density
	r.Restitution = mathf.Clamp(restitution, 0, 1)
	r.Static = static
	r.Mass = area * density

	r.vertices = nil
	r.transformedVertices = nil
	r.transformers = nil
	r.triangleIndices = nil
	r.transformVersion = 9999999999999999999
}

// Temp
func (r *DRigidBody) AsBox(width, height, density, restitution float64, static bool) {
	area := width * height

	if area > MAX_BODY_SIZE {
		area = MAX_BODY_SIZE
		println("warn: body size too big")
	}

	if area < MIN_BODY_SIZE {
		area = MIN_BODY_SIZE
		println("warn: body size too small")
	}

	if density < MIN_DENSITY {
		density = MIN_DENSITY
		println("warn: density too small")
	}

	if density > MAX_DENSITY {
		density = MAX_DENSITY
		println("warn: density too big")
	}

	r.Shape = BOX
	r.Width = width
	r.Height = height
	r.Area = area
	r.Density = density
	r.Restitution = mathf.Clamp(restitution, 0, 1)
	r.Static = static
	r.Mass = area * density

	r.vertices = []*sk.Vec2{
		{X: -width / 2, Y: -height / 2},
		{X: width / 2, Y: -height / 2},
		{X: width / 2, Y: height / 2},
		{X: -width / 2, Y: height / 2},
	}
	r.transformedVertices = make([]*sk.Vec2, len(r.vertices))
	// r.transformers = make([]*transformer, len(r.vertices))
	r.triangleIndices = []int{0, 1, 2, 0, 2, 3}
	r.transformVersion = 9999999999999999999
}

// Temp
func (r *DRigidBody) AsRandom() {
	r.AsCircle(random.Range(.1, .5), random.Range(.5, 1), random.Float(), false)
	// r.AsBox(random.Range(.1, .5), random.Range(.1, .5), random.Range(.5, 1), random.Float(), false)

	// if rand.Intn(2) == 0 {
	// r.AsCircle(random.Range(.1, .5), random.Range(.5, 1), random.Float(), false)
	// } else {
	// 	r.AsBox(random.Range(.1, .5), random.Range(.1, .5), random.Range(.5, 1), random.Float(), false)
	// }
}

var RigidBody = sk.NewComponentWithOptions(sk.ComponentOptions[DRigidBody]{
	OnCreated: func(r *DRigidBody) {
		r.Mass = 1
		r.Density = 1
		r.Restitution = 0.5
		r.Shape = CIRCLE
		r.Radius = 1
		r.image = ebiten.NewImage(500, 500)
	},
})
