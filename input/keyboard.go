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
	var cube *world.Cube = nil
	worldRef := world.GetInstance()
	for _, c := range worldRef.Cubes {
		if c.IsSelected {
			cube = c
			break
		}
	}

	if cube == nil {
		return
	}

	if rl.IsKeyDown(rl.KeyRight) {
		cube.PositionX += .25
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		cube.PositionX -= .25
	}
	if rl.IsKeyDown(rl.KeyUp) {
		cube.PositionY += .25
	}
	if rl.IsKeyDown(rl.KeyDown) {
		cube.PositionY -= .25
	}

	if rl.IsKeyDown(rl.KeyEqual) && (rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift)) {
		// Plus sign pressed
		cube.Width += .25
		cube.Height += .25
		cube.Length += .25
	}
	if rl.IsKeyDown(rl.KeyMinus) {
		cube.Width -= .25
		cube.Height -= .25
		cube.Length -= .25
	}

}
