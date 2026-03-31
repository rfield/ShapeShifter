package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"rjfield.com/graphics/world"
)

var dropdownItems string
var activeItemIndex int32
var editMode bool

func init() {
	dropdownItems = "Add an item...;Cube;Sphere"
	activeItemIndex = 0
	editMode = false
}

func DrawAndHandleAddItemMenu() {

	worldRef := world.GetInstance()

	if gui.DropdownBox(rl.NewRectangle(750, 80, 120, 30), dropdownItems, &activeItemIndex, editMode) {
		editMode = !editMode // Toggle edit mode on selection
		switch activeItemIndex {
		case 1: // Cube
			newCube := world.NewCube(0.0, 0.0, 0.0, 1.0, 1.0, 1.0, rl.Maroon)
			worldRef.AddShape(newCube)
		case 2: // Sphere
			newSphere := world.NewSphere(rl.Vector3{X: 0.0, Y: 0.0, Z: 0.0}, 1.0, rl.Blue)
			worldRef.AddShape(newSphere)
		}
		activeItemIndex = 0 // Reset to default after selection
	}
}
