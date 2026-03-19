package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Initialization
	screenWidth := int32(800)
	screenHeight := int32(450)

	raylib.InitWindow(screenWidth, screenHeight, "raylib [core] example - 3d cube in Go")

	// Define the camera to look into our 3D world
	camera := raylib.Camera3D{}
	camera.Position = raylib.NewVector3(10.0, 10.0, 10.0) // Camera position
	camera.Target = raylib.NewVector3(0.0, 0.0, 0.0)      // Camera looking at point
	camera.Up = raylib.NewVector3(0.0, 1.0, 0.0)          // Camera up vector (rotation towards target)
	camera.Fovy = 45.0                                    // Camera field-of-view Y
	camera.Projection = raylib.CameraPerspective          // Camera projection type

	// Define the cube parameters
	cubePosition := raylib.NewVector3(0.0, 0.0, 0.0)
	cubeWidth := float32(2.0)
	cubeHeight := float32(2.0)
	cubeLength := float32(2.0)
	cubeColor := raylib.Red

	raylib.SetTargetFPS(60) // Set our game to run at 60 frames-per-second

	// yAxisX := float32(screenWidth / 2)

	// Main game loop
	for !raylib.WindowShouldClose() { // Detect window close button or ESC key
		// Update
		// You can update camera movement here if needed. Raylib provides UpdateCamera()

		if raylib.IsKeyDown(raylib.KeyRight) || raylib.IsKeyDown(raylib.KeyD) {
			cubePosition.X += .25
		}
		if raylib.IsKeyDown(raylib.KeyLeft) || raylib.IsKeyDown(raylib.KeyA) {
			cubePosition.X -= .25
		}
		if raylib.IsKeyDown(raylib.KeyUp) || raylib.IsKeyDown(raylib.KeyW) {
			cubePosition.Y += .25
		}
		if raylib.IsKeyDown(raylib.KeyDown) || raylib.IsKeyDown(raylib.KeyS) {
			cubePosition.Y -= .25
		}

		// Draw
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RayWhite) // Clear the background with a color

		raylib.BeginMode3D(camera) // Start 3D mode

		// Draw the cube
		raylib.DrawCube(cubePosition, cubeWidth, cubeHeight, cubeLength, cubeColor)
		// Optional: Draw the cube wires to see the edges clearly
		raylib.DrawCubeWires(cubePosition, cubeWidth, cubeHeight, cubeLength, raylib.Maroon)

		// Optional: Draw a grid to visualize the 3D space
		raylib.DrawGrid(10, 1.0)

		// Draws a thick blue line vertically down the center
		// raylib.DrawLineEx(
		// 	raylib.Vector2{X: yAxisX, Y: 0},
		// 	raylib.Vector2{X: yAxisX, Y: float32(screenHeight)},
		// 	4.0, // Thickness
		// 	raylib.Blue,
		// )
		raylib.DrawLine3D(
			raylib.NewVector3(0, 0, 0),
			raylib.NewVector3(0, 5, 0),
			raylib.Green,
		)

		raylib.EndMode3D() // End 3D mode

		raylib.DrawFPS(10, 10) // Draw FPS counter in the corner

		raylib.EndDrawing()
	}

	// De-initialization
	raylib.CloseWindow() // Close window and OpenGL context
}
