package world

import rl "github.com/gen2brain/raylib-go/raylib"

var world World

type World struct {
	Cubes []*Cube
}

func init() {
	world = World{
		Cubes: []*Cube{},
	}
}

func GetInstance() *World {
	return &world
}

func (w *World) AddShape(cube *Cube) {
	w.Cubes = append(w.Cubes, cube)
}

func (w *World) Draw() {
	for _, shape := range w.Cubes {
		(*shape).Draw()
	}
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
