package raytracer

import (
	"testing"

	"github.com/go-gl/mathgl/mgl64"
)

func TestRayIntersectWithOldMechanism(t *testing.T) {
	bhbTest := MyTestingT{t}

	scene := CreateDefaultScene()

	scene.AddObject(&StandardSphere{
		Sphere: Sphere{
			Origin: mgl64.Vec3{0, 0, -50},
			Radius: 50},
		StandardMaterial: MaterialParams{
			RefractionValue: 0.0,
			Pd:              0.5,
			Ps:              1.0,
			Color:           mgl64.Vec3{1, 0, 0}}})

	camera := DefaultCamera()
	picture := CreatePicture(camera, 640, 480)

	r := picture.RayAt(320, 240)
	intersected, _, color := scene.internalRayIntersect(r, 0, 1.0, 1.0)
	bhbTest.CheckBool(intersected, "Must intersect")
	bhbTest.CheckVector(mgl64.Vec3{0.56347531233456, 0.00000000000000, 0.00000000000000}, color, "Color incorrect")
}
