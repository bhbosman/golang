package raytracer

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

// Sphere ...
type Sphere struct {
	Origin mgl64.Vec3
	Radius float64
}

// GetIntersect ...
func (s Sphere) GetIntersect(r Ray) (TIntersection, float64) {
	return s.calculateIntersection(r)
}

// Rotate ...
func (s *Sphere) Rotate(angle float64, axis mgl64.Vec3){
	rot := mgl64.HomogRotate3D(angle, axis)
	s.Origin = rot.Mul4x1(s.Origin.Vec4(1)).Vec3()
}

// calculateIntersection ...
func (s Sphere) calculateIntersection(r Ray)(TIntersection, float64){
	L := s.Origin.Sub(r.Origin)
	tca := L.Dot(r.Direction)
	if tca < 0 {
		return wiNoHit, 0.0
	}
	dd := L.Dot(L) - math.Pow(tca, 2)
	if dd < 0 {
		return wiNoHit, 0.0
	}
	if dd > math.Pow(s.Radius, 2) {
		return wiNoHit, 0.0
	}
	thc := math.Sqrt(math.Pow(s.Radius, 2.0) - dd)
	i1 := tca - thc
	i2 := tca + thc
	if i1 < 0.0 {
		return wiInside, i2
	} else {
		return wiOutSide, i1
	}
}

// GetNormalAt ...
func (s Sphere) GetNormalAt(p mgl64.Vec3) mgl64.Vec3 {
	result := p.Sub(s.Origin)
	return result.Normalize()
}
