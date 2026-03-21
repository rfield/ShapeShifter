package camera

import rl "github.com/gen2brain/raylib-go/raylib"

var camera rl.Camera3D

func init() {
	// Define the camera to look into our 3D world
	camera.Position = rl.NewVector3(10.0, 10.0, 10.0) // Camera position
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)      // Camera looking at point
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)          // Camera up vector (rotation towards target)
	camera.Fovy = 45.0                                // Camera field-of-view Y
	camera.Projection = rl.CameraPerspective          // Camera projection type
}

func GetInstance() *rl.Camera3D {
	return &camera
}
