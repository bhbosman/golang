package raytracer

import (
	"github.com/go-gl/mathgl/mgl64"
)

// Ray ...
type Ray struct {
	Origin    mgl64.Vec3
	Direction mgl64.Vec3
	// x         int
	// y         int
}

// // NewRay ...
// func NewRay(Origin mgl64.Vec3, Direction mgl64.Vec3, x int, y int, bias float64) Ray {
// 	return Ray{
// 		Origin:    Origin.Add(Direction.Mul(bias)),
// 		Direction: Direction,
// 		x:         x,
// 		y:         y}
// }

// PositionAt ...
func (r Ray) PositionAt(time float64) (result mgl64.Vec3) {
	return r.Origin.Add(r.Direction.Mul(time))
}

// GetX ...
func (r Ray) GetX() int {
	return r.x
}

// GetY ...
func (r Ray) GetY() int {
	return r.y
}
