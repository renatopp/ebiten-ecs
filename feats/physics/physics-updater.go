package physics

import (
	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/core"
	"github.com/renatopp/skald/mathf"
)

var TransformAndBodyQuery = sk.NewQuery[struct {
	Transform *core.DTransform
	RigidBody *DRigidBody
}]()

var PhysicsUpdater = sk.NewSystem(func(g *sk.Game) error {
	objs := TransformAndBodyQuery.Query()

	for i := 0; i < len(objs)-1; i++ {
		for j := i + 1; j < len(objs); j++ {
			a := objs[i]
			b := objs[j]
			ok, normal, depth := collideCircleCircle(a.Transform, b.Transform, a.RigidBody, b.RigidBody)

			if ok {
				hdepth := depth / 2

				a.Transform.MoveBy(normal.X*hdepth, normal.Y*hdepth)
				b.Transform.MoveBy(-normal.X*hdepth, -normal.Y*hdepth)
			}
		}
	}
	return nil
}, TransformAndBodyQuery)

func collideCircleCircle(at, bt *core.DTransform, ab, bb *DRigidBody) (ok bool, normal *sk.Vec2, depth float64) {
	ax, ay := at.GetPosition()
	bx, by := bt.GetPosition()

	t := ab.Radius + bb.Radius
	d := mathf.Distance(ax, ay, bx, by)
	if d < t {
		nx, ny := mathf.Normalize(ax-bx, ay-by)
		depth = t - d
		return true, &sk.Vec2{X: nx, Y: ny}, depth
	}

	return false, nil, 0
}
