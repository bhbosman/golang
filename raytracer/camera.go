package raytracer

import (
	"github.com/go-gl/mathgl/mgl64"
	bhbosmanVector "github.com/bhbosman/golang/vector"
)

// Camera ...
type Camera struct {
	viewPoint mgl64.Vec3
	plane     struct {
		topLeft     mgl64.Vec3
		topRight    mgl64.Vec3
		bottomLeft  mgl64.Vec3
		bottomRight mgl64.Vec3
	}
	direction struct {
		x mgl64.Vec3
		y mgl64.Vec3
		z mgl64.Vec3
	}
}

// Init2 ...
func (c *Camera) Init2(
	topLeft mgl64.Vec3,
	bottomLeft mgl64.Vec3,
	direction mgl64.Vec3,
	distance float64,
	aspectRatio float64) {

	// work out the diagonal vectors
	c.direction.y = topLeft.Sub(bottomLeft).Normalize()
	c.direction.z = direction.Normalize()
	c.direction.x = c.direction.z.Cross(c.direction.y)

	// work out the distances of the place
	TopDownDistance := bhbosmanVector.Distance(bottomLeft, topLeft)
	LeftRightDistance := TopDownDistance * aspectRatio

	c.plane.topLeft = topLeft
	c.plane.bottomLeft = bottomLeft
	c.plane.topRight = topLeft.Add(c.direction.x.Mul(LeftRightDistance))
	c.plane.bottomRight = bottomLeft.Add(c.direction.x.Mul(LeftRightDistance))

	viewPoint := c.plane.topLeft.Add(c.plane.topRight.Add(c.plane.bottomLeft.Add(c.plane.bottomRight)))
	viewPoint = viewPoint.Mul(0.25)
	c.viewPoint = viewPoint.Add(c.direction.z.Mul(-distance))
}

// DefaultCamera ...
func DefaultCamera() *Camera {
	c := &Camera{}
	c.Init2(
		mgl64.Vec3{-4.0, +3, 120.0},
		mgl64.Vec3{-4.0, -3, 120.0},
		mgl64.Vec3{0.0, 0.0, -1.0},
		5.0,
		4.0/3.0)

	return c
}

// ViewPoint ...
func (c *Camera) ViewPoint() mgl64.Vec3 {
	return c.viewPoint
}

// TopLeft ...
func (c *Camera) TopLeft() mgl64.Vec3 {
	return c.plane.topLeft
}

// TopRight ...
func (c *Camera) TopRight() mgl64.Vec3 {
	return c.plane.topRight
}

// BottomLeft ...
func (c *Camera) BottomLeft() mgl64.Vec3 {
	return c.plane.bottomLeft
}

// BottomRight ...
func (c *Camera) BottomRight() mgl64.Vec3 {
	return c.plane.bottomRight
}

// GetHeight ...
func (c *Camera) GetHeight() float64 {
	return bhbosmanVector.Distance(c.plane.topLeft, c.plane.bottomLeft)
}

// GetWidth ...
func (c *Camera) GetWidth() float64 {
	return bhbosmanVector.Distance(c.plane.topLeft, c.plane.topRight)
}

// GetDirectionX ...
func (c *Camera) GetDirectionX() mgl64.Vec3 {
	return c.direction.x
}

// GetDirectionY ...
func (c *Camera) GetDirectionY() mgl64.Vec3 {
	return c.direction.y
}

// GetDirectionZ ...
func (c *Camera) GetDirectionZ() mgl64.Vec3 {
	return c.direction.z
}
