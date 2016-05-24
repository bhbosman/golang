package raytracer

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

// PhongLight ...
type PhongLight struct {
}

//	dotln := nUnit.Dot(lUnit)
//	rUnit := nUnit.Mul(2.0 * dotln).Sub(lUnit)
//


// CalculateLight ...
func (pl PhongLight) CalculateLight(p, l, v mgl64.Vec3, light SceneLight, sceneObject SceneObjectLightCalculation) mgl64.Vec3 {
	nUnit := sceneObject.GetNormalAt(p)
	m := sceneObject.GetMaterial(p)
	vUnit := v.Sub(p).Normalize()
	lUnit := l.Sub(p).Normalize()
	compA := MultiScalar(light.GetIa().Mul(m.GetPa()), m.GetColor())
	compD := MultiScalar(light.GetId().Mul(m.GetPd()).Mul(math.Max(0, lUnit.Dot(nUnit))), m.GetColor())
	compS := MultiScalar((light.GetIs().Mul(m.GetPs())).Mul(math.Pow(math.Max(0, vUnit.Dot(lUnit)), light.GetF())), m.GetColor())
	if vUnit.Dot(nUnit) < 0.0 {
		compS = mgl64.Vec3{0, 0, 0}
	}
	result := compA.Add(compD).Add(compS)
	return result
}
