package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"rjfield.com/graphics/camera"
	"rjfield.com/graphics/config"
	"rjfield.com/graphics/input"
	"rjfield.com/graphics/ui"
	"rjfield.com/graphics/world"
)

func main() {

	configRef := config.GetConfig()
	rl.InitWindow(configRef.ScreenWidth, configRef.ScreenHeight, "ShapeShifter Demo - 3D objects in Go using Raylib")

	// Define the cameraRef to look into our 3D world
	cameraRef := camera.GetInstance()

	worldRef := world.GetInstance()
	worldRef.CreateDefaultShapes()

	// Rotating cube
	model := rl.LoadModelFromMesh(rl.GenMeshCube(1.0, 1.0, 1.0))
	rotation := 0.0

	rl.SetTargetFPS(60) // Set our game to run at 60 frames-per-second

	// Main game loop
	for !rl.WindowShouldClose() { // Detect window close button or ESC key

		worldRef.DumpInfo()

		ui.DrawAndHandleAddItemMenu()

		input.ProcessKeyboard()
		input.SetSelectedObjects()
		input.DragSelectedObjects()

		rotation += 1.0

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(*cameraRef)
		worldRef.DrawGrid()
		worldRef.Draw()

		// Draw the rotating cube
		// TO DO - make part of the world and add to the list of objects to draw
		// TO DO - figure out why it's NOT selectable
		rl.DrawModelEx(
			model,
			rl.NewVector3(-3.0, 2.0, -2.0), // Position
			rl.NewVector3(0.0, 1.0, 0.0),   // Rotation axis (Y-axis for spinning like a top)
			float32(rotation),              // Rotation angle
			rl.NewVector3(1.0, 1.0, 1.0),   // Scale
			rl.Gray,                        // Tint color
		)

		rl.EndMode3D()

		ui.ShowDialogOrButton()
		ui.DrawEditPanel()

		rl.EndDrawing()
	}

	rl.CloseWindow() // Close window and OpenGL context
}
