package raytracer

import (
	"github.com/go-gl/mathgl/mgl64"
)

// Plane ...
type Plane struct {
	Origin mgl64.Vec3
	Normal mgl64.Vec3
}

// GetOrigin ...
func (p Plane) GetOrigin() mgl64.Vec3 {
	return p.Origin
}

// GetNormal ...
func (p Plane) GetNormal() mgl64.Vec3 {
	return p.Normal
}

// GetIntersect ...
func (p Plane) GetIntersect(r Ray) (TIntersection, float64) {
	d := p.Normal.Dot(r.Direction)
	if d < 0.0 {
		return wiNoHit, 0.0
	}
	p0l0 := p.Origin.Sub(r.Origin)
	t := p0l0.Dot(p.Normal) / d
	return wiOutSide, t
}

// Rotate ...
func (p *Plane) Rotate(angle float64, axis mgl64.Vec3) {
	rot := mgl64.HomogRotate3D(angle, axis)
	p.Normal = rot.Mul4x1(p.Normal.Vec4(1)).Vec3()
	p.Origin = rot.Mul4x1(p.Origin.Vec4(1)).Vec3()
}

// GetNormalAt ...
func (p Plane) GetNormalAt(point mgl64.Vec3) mgl64.Vec3 {
	return (p.Normal.Normalize()).Mul(-1.0)
}
