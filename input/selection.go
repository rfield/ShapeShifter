package input

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"rjfield.com/graphics/camera"
	"rjfield.com/graphics/world"
)

func SetSelectedObjects() {
	cameraRef := camera.GetInstance()
	worldRef := world.GetInstance()

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		ray := rl.GetScreenToWorldRay(rl.GetMousePosition(), *cameraRef)

		for _, shape := range worldRef.Shapes {
			collision := shape.GetBoundingBoxRayCollision(ray)
			if collision.Hit {
				shape.SetColor(rl.Green)
				shape.SetSelected(true)
			} else {
				shape.SetColor(rl.Maroon)
				shape.SetSelected(false)
			}
		}
	}
}
