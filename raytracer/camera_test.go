package raytracer

import (
	//"math"
	"testing"

	"github.com/go-gl/mathgl/mgl64"
	bhbosmanTest "github.com/bhbosman/golang/testing"
)

func TestCamera0001(t *testing.T) {
	c := Camera{}
	c.Init2(
		mgl64.Vec3{-2.0, +1.5, 0.0},
		mgl64.Vec3{-2.0, -1.5, 0.0},
		mgl64.Vec3{0.0, 0.0, -1.0}.Normalize(),
		5.0,
		4.0/3.0)
		
		

	testCase := bhbosmanTest.MyTestingT{t}
	testCase.CheckVector(mgl64.Vec3{-2.0, 1.5, 0.0}, c.TopLeft(), "TopLeft incorrect")
	testCase.CheckVector(mgl64.Vec3{-2.0, -1.5, 0.0}, c.BottomLeft(), "BottomLeft incorrect")
	testCase.CheckVector(mgl64.Vec3{2.0, 1.5, 0.0}, c.TopRight(), "TopRight incorrect")
	testCase.CheckVector(mgl64.Vec3{2.0, -1.5, 0.0}, c.BottomRight(), "TopRight incorrect")
	testCase.CheckVector(mgl64.Vec3{0.0, 0.0, 5.0}, c.ViewPoint(), "ViewPoint incorrect")
	testCase.CheckVector(mgl64.Vec3{1, 0, 0}, c.direction.x, "Camera X vector incorrect ")
	testCase.CheckVector(mgl64.Vec3{0, 1, 0}, c.direction.y, "Camera Y vector incorrect")
	testCase.CheckVector(mgl64.Vec3{0, 0, -1}, c.direction.z, "Camera Z vector incorrect")
}

func TestCamera0002(t *testing.T) {
	c := Camera{}
	c.Init2(
		mgl64.Vec3{-12.0, +1.5, 0.0},
		mgl64.Vec3{-12.0, -1.5, 0.0},
		mgl64.Vec3{0.0, 0.0, -1.0}.Normalize(),
		5.0,
		4.0/3.0)

	testCase := bhbosmanTest.MyTestingT{t}
	testCase.CheckVector(mgl64.Vec3{-12.0, 1.5, 0.0}, c.TopLeft(), "TopLeft incorrect")
	testCase.CheckVector(mgl64.Vec3{-12.0, -1.5, 0.0}, c.BottomLeft(), "BottomLeft incorrect")
	testCase.CheckVector(mgl64.Vec3{-8.0, 1.5, 0.0}, c.TopRight(), "TopRight")
	testCase.CheckVector(mgl64.Vec3{-8.0, -1.5, 0.0}, c.BottomRight(), "TopRight")
	testCase.CheckVector(mgl64.Vec3{-10.0, 0.0, 5.0}, c.ViewPoint(), "ViewPoint")
	testCase.CheckVector(mgl64.Vec3{1, 0, 0}, c.direction.x, "XX")
	testCase.CheckVector(mgl64.Vec3{0, 1, 0}, c.direction.y, "YY")
	testCase.CheckVector(mgl64.Vec3{0, 0, -1}, c.direction.z, "ZZ")
}

func TestCamera0003(t *testing.T) {
	c := Camera{}
	c.Init2(
		mgl64.Vec3{-12.0, +1.5, 0.0},
		mgl64.Vec3{-12.0, -1.5, 0.0},
		mgl64.Vec3{0.0, 0.0, -1.0}.Normalize(),
		5.0,
		4.0/3.0)

	testCase := bhbosmanTest.MyTestingT{t}
	testCase.CheckVector(mgl64.Vec3{-12.0, 1.5, 0.0}, c.TopLeft(), "TopLeft incorrect")
	testCase.CheckVector(mgl64.Vec3{-12.0, -1.5, 0.0}, c.BottomLeft(), "BottomLeft incorrect")
	testCase.CheckVector(mgl64.Vec3{-8.0, 1.5, 0.0}, c.TopRight(), "TopRight")
	testCase.CheckVector(mgl64.Vec3{-8.0, -1.5, 0.0}, c.BottomRight(), "TopRight")
	testCase.CheckVector(mgl64.Vec3{-10.0, 0.0, 5.0}, c.ViewPoint(), "ViewPoint")
	testCase.CheckVector(mgl64.Vec3{1, 0, 0}, c.direction.x, "XX")
	testCase.CheckVector(mgl64.Vec3{0, 1, 0}, c.direction.y, "YY")
	testCase.CheckVector(mgl64.Vec3{0, 0, -1}, c.direction.z, "ZZ")
}
