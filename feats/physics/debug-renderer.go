package physics

import sk "github.com/renatopp/skald"

var DebugRenderer = sk.NewSystem(func(g *sk.Game) error {
	// for _, r := range PhysicsQuery.Query() {
	// 	r.RigidBody.image.Clear()

	// 	// if r.RigidBody.Shape == CIRCLE {
	// 	// 	vector.StrokeCircle(r.RigidBody.image, 250, 250, 250, 1, color.RGBA{255, 0, 0, 255}, false)
	// 	// } else {
	// 	// 	vector.StrokeRect(r.RigidBody.image, 1, 1, 498, 498, 1, color.RGBA{255, 0, 0, 255}, false)
	// 	// }

	// 	// bounds := r.RigidBody.image.Bounds()
	// 	// halfX, halfY := float64(bounds.Max.X)/2, float64(bounds.Max.Y)/2

	// 	// ops := &ebiten.DrawImageOptions{}
	// 	// ops.GeoM.Translate(-halfX, -halfY)
	// 	// ops.GeoM.Scale(r.Transform.scaleX, r.Transform.scaleY)
	// 	// ops.GeoM.Rotate(r.Transform.rotation * mathf.Deg2Rad)
	// 	// ops.GeoM.Translate(r.Transform.posX*g.Screen.PixelsPerUnit, r.Transform.posY*g.Screen.PixelsPerUnit)
	// 	// g.Renderer.Queue(r.Sprite.Layer, r.Sprite.ZIndex, r.Sprite.Texture.Image, ops)
	// }

	return nil
}, TransformAndBodyQuery)
