package raytracer

import (
	"testing"

	bhbosmanTest "github.com/bhbosman/golang/testing"
	"github.com/go-gl/mathgl/mgl64"
)

func TestPicture00001(t *testing.T) {
	c := &Camera{}
	c.Init2(
		mgl64.Vec3{-2.0, +1.5, 0.0},
		mgl64.Vec3{-2.0, -1.5, 0.0},
		mgl64.Vec3{0.0, 0.0, -1.0}.Normalize(),
		5.0,
		4.0/3.0)

	p := CreatePicture(c, 640, 480)

	r := p.RayAt(320, 240)

	bhbTest := bhbosmanTest.MyTestingT{T: t}
	bhbTest.CheckVector(mgl64.Vec3{-2.0, 1.5, 0.0}, c.TopLeft(), "TopLeft incorrect")
	bhbTest.CheckVector(mgl64.Vec3{-2.0, -1.5, 0.0}, c.BottomLeft(), "BottomLeft incorrect")
	bhbTest.CheckVector(mgl64.Vec3{2.0, 1.5, 0.0}, c.TopRight(), "TopRight")
	bhbTest.CheckVector(mgl64.Vec3{2.0, -1.5, 0.0}, c.BottomRight(), "TopRight")
	bhbTest.CheckVector(mgl64.Vec3{0.0, 0.0, 5.0}, c.ViewPoint(), "ViewPoint")
	bhbTest.CheckVector(mgl64.Vec3{1, 0, 0}, c.direction.x, "XX")
	bhbTest.CheckVector(mgl64.Vec3{0, 1, 0}, c.direction.y, "YY")
	bhbTest.CheckVector(mgl64.Vec3{0, 0, -1}, c.direction.z, "ZZ")

	bhbTest.CheckVector(mgl64.Vec3{0, 0, 0}, r.Origin, "A")
	bhbTest.CheckVector(c.GetDirectionZ(), r.Direction, "B")
}

func TestPicture00002(t *testing.T) {
	c := &Camera{}
	c.Init2(
		mgl64.Vec3{-2.0, +1.5, 0.0},
		mgl64.Vec3{-2.0, -1.5, 0.0},
		mgl64.Vec3{0.0, 0.0, -1.0}.Normalize(),
		5.0,
		4.0/3.0)

	p := CreatePicture(c, 640, 480)
	r := p.RayAt(320, 240)

	s := StandardSphere{
		Sphere{
			mgl64.Vec3{0, 0, -5},
			3},
		//mgl64.Vec3{1,0,0},
		MaterialParams{}}
	intersect, time := s.GetIntersect(r)

	bhbtest := bhbosmanTest.MyTestingT{T: t}
	bhbtest.CheckVector(mgl64.Vec3{-2.0, 1.5, 0.0}, c.TopLeft(), "TopLeft incorrect")
	bhbtest.CheckVector(mgl64.Vec3{-2.0, -1.5, 0.0}, c.BottomLeft(), "BottomLeft incorrect")
	bhbtest.CheckVector(mgl64.Vec3{2.0, 1.5, 0.0}, c.TopRight(), "TopRight")
	bhbtest.CheckVector(mgl64.Vec3{2.0, -1.5, 0.0}, c.BottomRight(), "TopRight")
	bhbtest.CheckVector(mgl64.Vec3{0.0, 0.0, 5.0}, c.ViewPoint(), "ViewPoint")
	bhbtest.CheckVector(mgl64.Vec3{1, 0, 0}, c.direction.x, "XX")
	bhbtest.CheckVector(mgl64.Vec3{0, 1, 0}, c.direction.y, "YY")
	bhbtest.CheckVector(mgl64.Vec3{0, 0, -1}, c.direction.z, "ZZ")

	bhbtest.CheckVector(mgl64.Vec3{0, 0, 0}, r.Origin, "A")
	bhbtest.CheckVector(c.GetDirectionZ(), r.Direction, "B")

	bhbtest.CheckBool(intersect != wiNoHit, "This must Intersect")
	bhbtest.CheckFloat64(2, time, "djdkdj")
}
