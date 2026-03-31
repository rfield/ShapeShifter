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
		selecting := false // Only allow one object to be selected at a time
		for _, shape := range worldRef.Shapes {
			collision := shape.GetBoundingBoxRayCollision(ray)
			if collision.Hit && !selecting {
				shape.SetColor(rl.Green)
				shape.SetSelected(true)
				selecting = true
			} else {
				shape.SetColor(rl.Maroon)
				shape.SetSelected(false)
			}
		}
	}
}

func DragSelectedObjects() {
	cameraRef := camera.GetInstance()
	worldRef := world.GetInstance()

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		ray := rl.GetScreenToWorldRay(rl.GetMousePosition(), *cameraRef)

		for _, shape := range worldRef.Shapes {

			collision := shape.GetBoundingBoxRayCollision(ray)
			if collision.Hit {
				distance := collision.Distance
				newPos := rl.Vector3Add(ray.Position, rl.Vector3Scale(ray.Direction, distance))
				shape.SetPosition(newPos)
			}
		}
	}
}
