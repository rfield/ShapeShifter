package world

// import rl "github.com/gen2brain/raylib-go/raylib"

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

func (w *World) AddCube(cube *Cube) {
	w.Cubes = append(w.Cubes, cube)
}

func (w *World) Draw() {
	for _, cube := range w.Cubes {
		cube.Draw()
	}
}
