package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"rjfield.com/graphics/camera"
	"rjfield.com/graphics/input"
	"rjfield.com/graphics/world"
)

func main() {
	// Initialization
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "ShapeShifter Demo - 3D objects in Go using Raylib")

	// Define the cameraRef to look into our 3D world
	cameraRef := camera.GetInstance()

	// Add some (cube) objects to the world
	worldRef := world.GetInstance()
	c1 := world.NewCube(0.0, 0.0, 0.0, 2.0, 2.0, 2.0, rl.Maroon)
	c1.IsSolid = true
	worldRef.AddShape(c1)
	c2 := world.NewCube(3.0, 0.0, 0.0, 1.0, 1.0, 1.0, rl.Maroon)
	worldRef.AddShape(c2)

	rl.SetTargetFPS(60) // Set our game to run at 60 frames-per-second

	// Main game loop
	for !rl.WindowShouldClose() { // Detect window close button or ESC key

		//rl.UpdateCamera(cameraRef, rl.CameraFirstPerson)

		input.ProcessKeyboard()

		input.SetSelectedObjects()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(*cameraRef) // Start 3D mode
		worldRef.DrawGrid()
		worldRef.Draw()
		rl.EndMode3D() // End 3D mode

		rl.DrawFPS(10, 10) // Draw FPS counter in the corner

		rl.EndDrawing()
	}

	rl.CloseWindow() // Close window and OpenGL context
}
