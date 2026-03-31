package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"rjfield.com/graphics/config"
)

// Are we showing the message box? Show it when the button is clicked,
// and hide it when the user clicks "Okay" or "Cancel"
var showMessageBox bool

func init() {
	showMessageBox = false
}

func ShowDialogOrButton() {
	configRef := config.GetConfig()

	// Dialog box parameters
	boxBounds := rl.NewRectangle(float32(configRef.ScreenWidth/2-200), float32(configRef.ScreenHeight/2-50), 400, 100)

	// Main screen controls
	if gui.Button(rl.NewRectangle(750, 30, 120, 30), "About...") {
		showMessageBox = true
	}

	if showMessageBox {
		// Draw the dialog box
		// The return value indicates which button was pressed (0: none, 1: Ok, 2: Cancel - depends on implementation)
		result := gui.MessageBox(boxBounds, "About ShapeShifter", "ShapeShifter is a simple 3D graphics application built with Go and Raylib.", "Okay;Cancel")

		if result > 0 {
			// Handle button click (e.g., close the dialog)
			showMessageBox = false
			// You can add logic here based on 'result' (e.g., if result == 1 { ... })
		}
	}
}
