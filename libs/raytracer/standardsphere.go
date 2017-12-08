package raytracer

import (
	"github.com/go-gl/mathgl/mgl64"
)

// StandardSphere ...
type StandardSphere struct {
	Sphere
	StandardMaterial MaterialParams
}

// Init ...
func (s *StandardSphere) Init() {
}

// GetMaterial ...
func (s StandardSphere) GetMaterial(point mgl64.Vec3) Material {
	return s.StandardMaterial
}

// Description ...
func (s StandardSphere) Description() string {
	return "StandardSphere"
}

// CrystalBallSphere ...
type CrystalBallSphere struct {
	Sphere
	StandardMaterial MaterialParams
	Children         []StandardSphere
}

// Init ...
func (s CrystalBallSphere) Init() {

}

// AddChild ...
func (s *CrystalBallSphere) AddChild(child StandardSphere) {

	s.Children = append(s.Children, child)
}

// GetIntersect ...
func (s CrystalBallSphere) GetIntersect(r Ray) (TIntersection, float64) {
	return s.calculateIntersection(r)
}
