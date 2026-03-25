package world

import rl "github.com/gen2brain/raylib-go/raylib"

type Sphere struct {
	Position rl.Vector3
	Radius   float32
	Color    rl.Color
	Selected bool
}

func NewSphere(position rl.Vector3, radius float32, color rl.Color) *Sphere {
	return &Sphere{
		Position: position,
		Radius:   radius,
		Color:    color,
		Selected: false,
	}
}

func (s *Sphere) Draw() {
	rl.DrawSphere(s.Position, s.Radius, s.Color)
}

func (s *Sphere) Move(delta rl.Vector3) {
	s.Position = rl.Vector3Add(s.Position, delta)
}

func (s *Sphere) SetPosition(pos rl.Vector3) {
	s.Position.X = pos.X
	s.Position.Y = pos.Y
	// s.Position.Z = pos.Z
}

func (s *Sphere) GetPosition() rl.Vector3 {
	return s.Position
}

func (s *Sphere) Grow(factor float32) {
	s.Radius *= factor
}

func (s *Sphere) SetColor(color rl.Color) {
	s.Color = color
}

func (s *Sphere) SetSelected(isSelected bool) {
	s.Selected = isSelected
}

func (s *Sphere) GetSelected() bool {
	return s.Selected
}

func (s *Sphere) GetBoundingBox() (min, max rl.Vector3) {
	min = rl.Vector3{
		X: s.Position.X - s.Radius,
		Y: s.Position.Y - s.Radius,
		Z: s.Position.Z - s.Radius,
	}
	max = rl.Vector3{
		X: s.Position.X + s.Radius,
		Y: s.Position.Y + s.Radius,
		Z: s.Position.Z + s.Radius,
	}
	return min, max
}

func (s *Sphere) GetBoundingBoxRayCollision(ray rl.Ray) (collision rl.RayCollision) {
	return rl.GetRayCollisionSphere(ray, s.Position, s.Radius)
}

func (s *Sphere) GetShapeType() string {
	return "Sphere"
}
