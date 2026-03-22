package input

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"rjfield.com/graphics/camera"
	"rjfield.com/graphics/world"
)

func ProcessKeyboard() {

	// Handle camera movement
	cameraRef := camera.GetInstance()
	if rl.IsKeyDown(rl.KeyK) {
		cameraRef.Position.X += .25
	}
	if rl.IsKeyDown(rl.KeyJ) {
		cameraRef.Position.X -= .25
	}

	// Handle cube manipulation
	var shape world.Shape = nil
	worldRef := world.GetInstance()
	for _, s := range worldRef.Shapes {
		if s.GetSelected() {
			shape = s
			break
		}
	}

	if shape == nil {
		return
	}

	if rl.IsKeyDown(rl.KeyRight) {
		shape.Move(rl.NewVector3(.25, 0, 0))
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		shape.Move(rl.NewVector3(-.25, 0, 0))
	}
	if rl.IsKeyDown(rl.KeyUp) {
		shape.Move(rl.NewVector3(0, .25, 0))
	}
	if rl.IsKeyDown(rl.KeyDown) {
		shape.Move(rl.NewVector3(0, -.25, 0))
	}

	if rl.IsKeyDown(rl.KeyEqual) && (rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift)) {
		// Plus sign pressed
		shape.Grow(1.1)
	}
	if rl.IsKeyDown(rl.KeyMinus) {
		shape.Grow(1.0 / 1.1)
	}

}
