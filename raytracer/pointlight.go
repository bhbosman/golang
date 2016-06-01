package raytracer

import (
	"github.com/go-gl/mathgl/mgl64"
)

// PointLight ...
type PointLight struct {
	Color    mgl64.Vec3
	Position mgl64.Vec3
	B        float64
	F        float64
}

// GetF ...
func (pointLight PointLight) GetF() float64 {
	return pointLight.F
}

// GetPosition ...
func (pointLight PointLight) GetPosition() mgl64.Vec3 {
	return pointLight.Position
}

// GetColor ...
func (pointLight PointLight) GetColor() mgl64.Vec3 {
	return pointLight.Color
}

// GetIa ...
func (pointLight PointLight) GetIa() mgl64.Vec3 {
	return mgl64.Vec3{}
}

// GetId ...
func (pointLight PointLight) GetId() mgl64.Vec3 {
	return pointLight.Color.Mul(pointLight.B)
}

// GetIs ...
func (pointLight PointLight) GetIs() mgl64.Vec3 {
	return pointLight.Color.Mul(pointLight.B)
}
