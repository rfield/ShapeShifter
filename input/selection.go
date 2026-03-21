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

		for _, cube := range worldRef.Cubes {
			collision := cube.GetBoundingBoxRayCollision(ray)
			if collision.Hit {
				cube.Color = rl.Green
				cube.IsSelected = true
			} else {
				cube.Color = rl.Maroon
				cube.IsSelected = false
			}
		}
	}
}
