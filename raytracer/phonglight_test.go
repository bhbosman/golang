package raytracer

import (
	"testing"

	bhbosmanTest "github.com/bhbosman/golang/testing"
	"github.com/go-gl/mathgl/mgl64"
)

type testMaterial struct {
	color mgl64.Vec3
	pa    float64
	pd    float64
	ps    float64
}

func (sm testMaterial) GetColor() mgl64.Vec3 {
	return sm.color
}

func (sm testMaterial) GetPa() float64 {
	return sm.pa
}

func (sm testMaterial) GetPs() float64 {
	return sm.ps
}

func (sm testMaterial) GetPd() float64 {
	return sm.pd
}

func (sm testMaterial) GetRefraction() float64 {
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
	sm := &testMaterial{
		color: data.color,
		pa:    data.pa,
		pd:    data.pd,
		ps:    data.ps}
	return sm
}

func TestPhongLight002(t *testing.T) {
	f := 8.0
	pointOfInterest := mgl64.Vec3{5, 5, 2}
	lightSourcePosition := mgl64.Vec3{20, 7, 2}
	viewerPosition := mgl64.Vec3{20, 0, 10}

	light := PointLight{
		Color:    mgl64.Vec3{1, 1, 1},
		Position: lightSourcePosition,
		B:        1.0,
		F:        f}
	data := TestSphereData{}
	data.n = mgl64.Vec3{0.928477, 0, 0.371391}
	data.pa = 0.1
	data.ps = 0.5
	data.pd = 0.0
	data.color = mgl64.Vec3{1, 1, 1}
	phongLight := PhongLight{}
	value := phongLight.CalculateLight(pointOfInterest, lightSourcePosition, viewerPosition, light, data)

	bhbTest := bhbosmanTest.MyTestingT{T: t}
	bhbTest.CheckVector(mgl64.Vec3{0.08539199538899721, 0.08539199538899721, 0.08539199538899721}, value, "")
}
