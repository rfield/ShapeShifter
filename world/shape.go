package world

import rl "github.com/gen2brain/raylib-go/raylib"

type Shape interface {
	Draw()
	Move(delta rl.Vector3)
	Grow(factor float32)
	SetColor(color rl.Color)
	SetSelected(isSelected bool)
	GetSelected() bool
	GetBoundingBox() (min, max rl.Vector3)
	GetBoundingBoxRayCollision(ray rl.Ray) (collision rl.RayCollision)
}
