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

	rl.SetTargetFPS(60) // Set our game to run at 60 frames-per-second

	// Main game loop
	for !rl.WindowShouldClose() { // Detect window close button or ESC key

		worldRef.DumpInfo()

		ui.DrawAndHandleAddItemMenu()

		input.ProcessKeyboard()
		input.SetSelectedObjects()
		input.DragSelectedObjects()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(*cameraRef)
		worldRef.DrawGrid()
		worldRef.Draw()
		rl.EndMode3D()

		ui.ShowDialogOrButton()
		ui.DrawEditPanel()

		rl.EndDrawing()
	}

	rl.CloseWindow() // Close window and OpenGL context
}
