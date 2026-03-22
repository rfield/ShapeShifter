package world

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cube struct {
	PositionX  float32
	PositionY  float32
	PositionZ  float32
	Width      float32
	Height     float32
	Length     float32
	Color      color.RGBA
	IsSolid    bool
	IsSelected bool
}

func NewCube(positionX, positionY, positionZ, width, height, length float32, color color.RGBA) *Cube {
	return &Cube{
		PositionX:  positionX,
		PositionY:  positionY,
		PositionZ:  positionZ,
		Width:      width,
		Height:     height,
		Length:     length,
		Color:      color,
		IsSolid:    false,
		IsSelected: false,
	}
}

func (c *Cube) Draw() {
	position := rl.NewVector3(c.PositionX, c.PositionY, c.PositionZ)
	if c.IsSolid {
		rl.DrawCube(position, c.Width, c.Height, c.Length, c.Color)
	} else {
		rl.DrawCubeWires(position, c.Width, c.Height, c.Length, c.Color)
	}
}

func (c *Cube) Move(delta rl.Vector3) {
	c.PositionX += delta.X
	c.PositionY += delta.Y
	c.PositionZ += delta.Z
}

func (c *Cube) Grow(factor float32) {
	c.Width *= factor
	c.Height *= factor
	c.Length *= factor
}

func (c *Cube) SetColor(color rl.Color) {
	c.Color = color
}

func (c *Cube) SetSelected(isSelected bool) {
	c.IsSelected = isSelected
}

func (c *Cube) GetSelected() bool {
	return c.IsSelected
}

func (c *Cube) GetBoundingBox() (min, max rl.Vector3) {
	min = rl.NewVector3(c.PositionX-c.Width/2, c.PositionY-c.Height/2, c.PositionZ-c.Length/2)
	max = rl.NewVector3(c.PositionX+c.Width/2, c.PositionY+c.Height/2, c.PositionZ+c.Length/2)
	return min, max
}

func (c *Cube) GetBoundingBoxRayCollision(ray rl.Ray) (collision rl.RayCollision) {
	min, max := c.GetBoundingBox()
	cubeBBox := rl.NewBoundingBox(min, max)
	collision = rl.GetRayCollisionBox(ray, cubeBBox)
	return collision
}
