package world

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"rjfield.com/graphics/camera"
)

var world World

type World struct {
	Shapes []Shape
}

func init() {
	world = World{
		Shapes: []Shape{},
	}
}

func GetInstance() *World {
	return &world
}

func (w *World) AddShape(shape Shape) {
	w.Shapes = append(w.Shapes, shape)
}

func (w *World) Draw() {
	for _, shape := range w.Shapes {
		shape.Draw()
	}
}

// For now just return the first shape, but we will want to track which shape is selected in the future
func (w *World) GetSelectedShape() Shape {
	// for _, shape := range w.Shapes {
	// 	if shape.GetSelected() {
	// 		return shape
	// 	}
	// }
	return w.Shapes[0] // Placeholder: always return the first shape for now
	// return nil
}

func (w *World) DrawGrid() {
	rl.DrawGrid(10, 1.0)
	rl.DrawLine3D(
		rl.NewVector3(0, -5, 0),
		rl.NewVector3(0, 5, 0),
		rl.Blue,
	)
	rl.DrawLine3D(
		rl.NewVector3(0, 0, -5),
		rl.NewVector3(0, 0, 5),
		rl.Blue,
	)
	rl.DrawLine3D(
		rl.NewVector3(-5, 0, 0),
		rl.NewVector3(5, 0, 0),
		rl.Blue,
	)
}

func (w *World) DumpInfo() {
	for i, shape := range w.Shapes {
		position := shape.GetPosition()
		rl.DrawText(
			fmt.Sprintf("%s %d: Pos(%.2f, %.2f, %.2f)",
				shape.GetShapeType(), i, position.X, position.Y, position.Z),
			10,
			int32(30+i*20),
			14,
			rl.Black,
		)
	}

	cameraRef := camera.GetInstance()
	rl.DrawText(
		fmt.Sprintf("Camera: Pos(%.2f, %.2f, %.2f) Target(%.2f, %.2f, %.2f)",
			cameraRef.Position.X, cameraRef.Position.Y, cameraRef.Position.Z,
			cameraRef.Target.X, cameraRef.Target.Y, cameraRef.Target.Z),
		10,
		int32(30+len(w.Shapes)*20),
		14,
		rl.Black,
	)
}

func (w *World) CreateDefaultShapes() {
	c1 := NewCube(0.0, 0.0, 0.0, 2.0, 2.0, 2.0, rl.Maroon)
	c1.IsSolid = true
	w.AddShape(c1)
	c2 := NewCube(3.0, 0.0, 0.0, 1.0, 1.0, 1.0, rl.Maroon)
	w.AddShape(c2)
	s1 := NewSphere(rl.Vector3{X: -3.0, Y: 0.0, Z: 0.0}, 1.0, rl.Blue)
	w.AddShape(s1)
}
