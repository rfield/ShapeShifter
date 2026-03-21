package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"rjfield.com/graphics/camera"
	"rjfield.com/graphics/world"
)

func main() {
	// Initialization
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - 3d cube in Go")

	// Define the cameraRef to look into our 3D world
	cameraRef := camera.GetInstance()
	// camera := rl.Camera3D{}
	// camera.Position = rl.NewVector3(10.0, 10.0, 10.0) // Camera position
	// camera.Target = rl.NewVector3(0.0, 0.0, 0.0)      // Camera looking at point
	// camera.Up = rl.NewVector3(0.0, 1.0, 0.0)          // Camera up vector (rotation towards target)
	// camera.Fovy = 45.0                                // Camera field-of-view Y
	// camera.Projection = rl.CameraPerspective          // Camera projection type

	// rl.SetCameraMode(camera, rl.CameraFree) // Enables built-in mouse/keyboard control

	// Define the c1 parameters
	worldRef := world.GetInstance()
	c1 := world.NewCube(0.0, 0.0, 0.0, 2.0, 2.0, 2.0, rl.Maroon)
	worldRef.AddCube(c1)
	c2 := world.NewCube(3.0, 0.0, 0.0, 1.0, 1.0, 1.0, rl.Maroon)
	worldRef.AddCube(c2)
	// cubePosition := rl.NewVector3(0.0, 0.0, 0.0)
	// cubeWidth := float32(2.0)
	// cubeHeight := float32(2.0)
	// cubeLength := float32(2.0)
	// // cubeColor := raylib.Red
	// wireCubeColor := rl.Maroon

	rl.SetTargetFPS(60) // Set our game to run at 60 frames-per-second

	// yAxisX := float32(screenWidth / 2)

	// Main game loop
	for !rl.WindowShouldClose() { // Detect window close button or ESC key
		// Update
		// You can update camera movement here if needed. Raylib provides UpdateCamera()
		// rl.UpdateCamera(&camera, rl.CameraFirstPerson) // Update camera (simply update camera position and rotation)

		processKeyboard()

		// min := rl.NewVector3(cubePosition.X-cubeWidth/2, cubePosition.Y-cubeHeight/2, cubePosition.Z-cubeLength/2)
		// max := rl.NewVector3(cubePosition.X+cubeWidth/2, cubePosition.Y+cubeHeight/2, cubePosition.Z+cubeLength/2)
		// cubeBBox := rl.NewBoundingBox(min, max)
		// ray := rl.GetScreenToWorldRay(rl.GetMousePosition(), *cameraRef)
		// if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		// 	collision := rl.GetRayCollisionBox(ray, cubeBBox)
		// 	if collision.Hit {
		// 		wireCubeColor = rl.Green
		// 	} else {
		// 		wireCubeColor = rl.Maroon
		// 	}
		// }

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
			// collision := c1.GetBoundingBoxRayCollision(ray)
			// if collision.Hit {
			// 	c1.Color = rl.Green
			// } else {
			// 	c1.Color = rl.Maroon
			// }
		}

		// Draw
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite) // Clear the background with a color

		rl.BeginMode3D(*cameraRef) // Start 3D mode

		// Draw the cube
		// raylib.DrawCube(cubePosition, cubeWidth, cubeHeight, cubeLength, cubeColor)
		// Optional: Draw the cube wires to see the edges clearly
		// rl.DrawCubeWires(cubePosition, cubeWidth, cubeHeight, cubeLength, wireCubeColor)
		for _, cube := range worldRef.Cubes {
			cube.Draw()
		}

		// Optional: Draw a grid to visualize the 3D space
		rl.DrawGrid(10, 1.0)

		rl.DrawLine3D(
			rl.NewVector3(0, 0, 0),
			rl.NewVector3(0, 5, 0),
			rl.Green,
		)

		rl.EndMode3D() // End 3D mode

		rl.DrawFPS(10, 10) // Draw FPS counter in the corner

		rl.EndDrawing()
	}

	// De-initialization
	rl.CloseWindow() // Close window and OpenGL context
}

func processKeyboard() {
	// cubePosition *rl.Vector3, //camera *rl.Camera3D,
	// cubewidth *float32, cubeHeight *float32, cubeLength *float32) {

	cameraRef := camera.GetInstance()

	if rl.IsKeyDown(rl.KeyK) {
		cameraRef.Position.X += .25
	}
	if rl.IsKeyDown(rl.KeyJ) {
		cameraRef.Position.X -= .25
	}

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
