package raytracer

import (
	"fmt"
	"testing"

	"github.com/go-gl/mathgl/mgl64"
	bhbosmanTest "github.com/bhbosman/golang/testing"
)

func TestPlane00001(t *testing.T) {
	data := struct {
		p StandardPlane
		r Ray
	}{
		StandardPlane{
			Plane{
				Origin: mgl64.Vec3{0, 1, 0},
				Normal: mgl64.Vec3{0, 1, 0}},
			MaterialParams{}},
		Ray{
			mgl64.Vec3{0, 0, 0},
			mgl64.Vec3{1.0, 0, 0},
			0,
			0}}

	b, _ := data.p.GetIntersect(data.r)

	testCase := bhbosmanTest.MyTestingT{t}
	testCase.CheckBool(b != wiNoHit, "must intersect")
}

func TestPlane00002(t *testing.T) {
	data := struct {
		p StandardPlane
		r Ray
	}{
		StandardPlane{
			Plane{
				mgl64.Vec3{0, 1, 0},
				mgl64.Vec3{0, 1, 0}},
			MaterialParams{}},
		Ray{
			mgl64.Vec3{0, 0, 0},
			mgl64.Vec3{0, 0, 1},
			0,
			0}}

	b, tt := data.p.GetIntersect(data.r)

	fmt.Println(tt)
	testCase := bhbosmanTest.MyTestingT{t}
	testCase.CheckBool(b != wiNoHit, "must intersect")
}

func TestPlane00003(t *testing.T) {

	c := &Camera{}
	c.Init2(
		mgl64.Vec3{-2.0, +1.5, 0.0},
		mgl64.Vec3{-2.0, -1.5, 0.0},
		mgl64.Vec3{0.0, 0.0, -1.0}.Normalize(),
		5.0,
		4.0/3.0)

	p := CreatePicture(c, 640, 480)
	r := p.RayAt(320, 480)
	plane := StandardPlane{
		Plane{
			mgl64.Vec3{0, -10, 0},
			mgl64.Vec3{0, -1, 0}},
		MaterialParams{}}
	b, _ := plane.GetIntersect(r)

	bhbtest := bhbosmanTest.MyTestingT{t}
	bhbtest.CheckVector(mgl64.Vec3{-2.0, 1.5, 0.0}, c.TopLeft(), "TopLeft incorrect")
	bhbtest.CheckVector(mgl64.Vec3{-2.0, -1.5, 0.0}, c.BottomLeft(), "BottomLeft incorrect")
	bhbtest.CheckVector(mgl64.Vec3{2.0, 1.5, 0.0}, c.TopRight(), "TopRight")
	bhbtest.CheckVector(mgl64.Vec3{2.0, -1.5, 0.0}, c.BottomRight(), "TopRight")
	bhbtest.CheckVector(mgl64.Vec3{0.0, 0.0, 5.0}, c.ViewPoint(), "ViewPoint")
	bhbtest.CheckVector(mgl64.Vec3{1, 0, 0}, c.direction.x, "XX")
	bhbtest.CheckVector(mgl64.Vec3{0, 1, 0}, c.direction.y, "YY")
	bhbtest.CheckVector(mgl64.Vec3{0, 0, -1}, c.direction.z, "ZZ")

	bhbtest.CheckVector(mgl64.Vec3{0, -1.5, 0}, r.Origin, "A")
	bhbtest.CheckVector(mgl64.Vec3{0, -0.2873478855663454, -0.9578262852211514}, r.Direction, "A")
	bhbtest.CheckBool(b != wiNoHit, "ddddd")

}

func TestPlane00004(t *testing.T) {

	c := &Camera{}
	c.Init2(
		mgl64.Vec3{-2.0, +1.5, 0.0},
		mgl64.Vec3{-2.0, -1.5, 0.0},
		mgl64.Vec3{0.0, 0.0, -1.0}.Normalize(),
		5.0,
		4.0/3.0)

	p := CreatePicture(c, 640, 480)
	r := p.RayAt(320, 80)
	plane := StandardPlane{
		Plane{
			mgl64.Vec3{0, -10, 0},
			mgl64.Vec3{0, -1, 0}}, MaterialParams{}}
	b, _ := plane.GetIntersect(r)

	bhbtest := bhbosmanTest.MyTestingT{t}
	bhbtest.CheckVector(mgl64.Vec3{-2.0, 1.5, 0.0}, c.TopLeft(), "TopLeft incorrect")
	bhbtest.CheckVector(mgl64.Vec3{-2.0, -1.5, 0.0}, c.BottomLeft(), "BottomLeft incorrect")
	bhbtest.CheckVector(mgl64.Vec3{2.0, 1.5, 0.0}, c.TopRight(), "TopRight")
	bhbtest.CheckVector(mgl64.Vec3{2.0, -1.5, 0.0}, c.BottomRight(), "TopRight")
	bhbtest.CheckVector(mgl64.Vec3{0.0, 0.0, 5.0}, c.ViewPoint(), "ViewPoint")
	bhbtest.CheckVector(mgl64.Vec3{1, 0, 0}, c.direction.x, "XX")
	bhbtest.CheckVector(mgl64.Vec3{0, 1, 0}, c.direction.y, "YY")
	bhbtest.CheckVector(mgl64.Vec3{0, 0, -1}, c.direction.z, "ZZ")

	bhbtest.CheckVector(mgl64.Vec3{0, 1, 0}, r.Origin, "A")
	bhbtest.CheckVector(mgl64.Vec3{0, 0.19611613513818404, -0.9805806756909202}, r.Direction, "A")
	bhbtest.CheckBool(b == wiNoHit, "ddddd")

}
