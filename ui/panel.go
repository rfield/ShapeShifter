package ui

import (
	"log"
	"strconv"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"rjfield.com/graphics/config"
	"rjfield.com/graphics/world"
)

var panelBounds rl.Rectangle
var textBoxBounds rl.Rectangle
var inputText string
var textEditMode bool
var buttonClicked bool

func init() {
	configRef := config.GetConfig()
	panelBounds = rl.NewRectangle(0, (float32)(configRef.ScreenHeight-configRef.EditPanelHeight), (float32)(configRef.ScreenWidth), (float32)(configRef.EditPanelHeight))
	textBoxBounds = rl.NewRectangle(panelBounds.X+580, panelBounds.Y+38, 50, 30)
	inputText = "0"
	textEditMode = false
	buttonClicked = false
}

func DrawEditPanel() {
	configRef := config.GetConfig()
	worldRef := world.GetInstance()

	gui.Panel(panelBounds, "Change Attributes")

	// Add elements inside the panel
	// For an immediate mode GUI, you call the functions every frame.
	// Position the button relative to the panel's top-left corner (panelBounds.X, panelBounds.Y)
	buttonBounds := rl.NewRectangle(750, panelBounds.Y+35, 120, 30)
	if gui.Button(buttonBounds, "Apply") {
		buttonClicked = !buttonClicked
		zValue, err := strconv.Atoi(inputText)
		if err == nil {
			shapeRef := worldRef.GetSelectedShape()
			if shapeRef != nil {
				shapeRef.Move(rl.NewVector3(0, 0, float32(zValue)))
			}
		}
	}

	rl.DrawText("Enter a Z-value 'delta' for the selected object and press return or click 'Apply'", int32(panelBounds.X+10), int32(panelBounds.Y+42), 14, rl.Black)

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), textBoxBounds) {
			textEditMode = true
		} else {
			textEditMode = false
		}
	}
	// Draw the text box inside the panel
	// The GuiTextBox function updates the 'inputText' string and returns true if 'Enter' is pressed or defocused.
	// The 'maxInputChars' defines the buffer size and limit.
	if gui.TextBox(textBoxBounds, &inputText, int(configRef.MaxInputChars), textEditMode) {
		// Optional: Handle action when input is finished (e.g., Enter key pressed)
		log.Printf("Input finalized: %s\n", inputText)
		zValue, err := strconv.Atoi(inputText)
		if err == nil {
			shapeRef := worldRef.GetSelectedShape()
			if shapeRef != nil {
				shapeRef.Move(rl.NewVector3(0, 0, float32(zValue)))
			}
		}
		textEditMode = false // Exit edit mode after finalization
	}
}
