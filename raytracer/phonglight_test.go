package raytracer

import (
	"testing"

	"github.com/go-gl/mathgl/mgl64"
	bhbosmanTest "github.com/bhbosman/golang/testing"
)

//	dotln := nUnit.Dot(lUnit)
//	rUnit := nUnit.Mul(2.0 * dotln).Sub(lUnit)
//

type test_material struct {
	color mgl64.Vec3
	pa    float64
	pd    float64
	ps    float64
}

func (sm test_material) GetColor() mgl64.Vec3 {
	return sm.color
}

func (sm test_material) GetPa() float64 {
	return sm.pa
}

func (sm test_material) GetPs() float64 {
	return sm.ps
}

func (sm test_material) GetPd() float64 {
	return sm.pd
}

func (sm test_material) GetRefraction() float64 {
	return 0.0
}

type TestSphereData struct {
	n     mgl64.Vec3
	pa    float64
	ps    float64
	pd    float64
	color mgl64.Vec3
}

func (data TestSphereData) GetIntersect(r Ray) (float64, bool) {
	return 0.0, false
}

func (data TestSphereData) GetNormalAt(point mgl64.Vec3) mgl64.Vec3 {
	return data.n
}

func (data TestSphereData) GetMaterial(p mgl64.Vec3) Material {
	sm := &test_material{data.color, data.pa, data.pd, data.ps}
	return sm
}

/*
func (data TestSphereData) GetColor() mgl64.Vec3 {
	return data.color
}
func (data TestSphereData) GetPa() float64 {
	return data.pa
}

func (data TestSphereData) GetPs() float64 {
	return data.ps
}

func (data TestSphereData) GetPd() float64 {
	return data.pd
}
*/
func TestPhongLight002(t *testing.T) {
	f := 8.0
	pointOfInterest := mgl64.Vec3{5, 5, 2}
	lightSourcePosition := mgl64.Vec3{20, 7, 2}
	viewerPosition := mgl64.Vec3{20, 0, 10}

	light := PointLight{mgl64.Vec3{1, 1, 1}, lightSourcePosition, 1.0, f}
	data := TestSphereData{}
	data.n = mgl64.Vec3{0.928477, 0, 0.371391}
	data.pa = 0.1
	data.ps = 0.5
	data.pd = 0.0
	data.color = mgl64.Vec3{1, 1, 1}
	phongLight := PhongLight{}
	value := phongLight.CalculateLight(pointOfInterest, lightSourcePosition, viewerPosition, light, data)

	bhbTest := bhbosmanTest.MyTestingT{t}
	bhbTest.CheckVector(mgl64.Vec3{0.08539199538899721, 0.08539199538899721, 0.08539199538899721}, value, "")
}
