package raytracer

import (
	"github.com/go-gl/mathgl/mgl64"
)

// StandardPlane ...
type StandardPlane struct {
	Plane
	StandardMaterial MaterialParams
}

// Init ...
func (p *StandardPlane) Init() {
	p.Normal = p.Normal.Normalize()
}

// GetMaterial ...
func (p StandardPlane) GetMaterial(point mgl64.Vec3) Material {
	return p.StandardMaterial
}
