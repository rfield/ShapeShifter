package world

import rl "github.com/gen2brain/raylib-go/raylib"

type Shape interface {
	Draw()
	GetBoundingBox() (min, max rl.Vector3)
	GetBoundingBoxRayCollision(ray rl.Ray) (collision rl.RayCollision)
}
