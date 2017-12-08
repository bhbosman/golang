package raytracer

import (
	"math"
	"testing"
	//"fmt"

	bhbosmanTest "github.com/bhbosman/code/golang/testing"
	"github.com/go-gl/mathgl/mgl64"
)

func float64Eq(x, y float64) bool { return math.Abs(x-y) < 1e-14 }

func TestSphere0001(t *testing.T) {
	data := struct {
		s         StandardSphere
		r         Ray
		intersect TIntersection
		where     float64
	}{
		StandardSphere{
			Sphere{
				mgl64.Vec3{5.0, 0.0, 0.0},
				1.0},
			//mgl64.Vec3{},
			MaterialParams{}},
		Ray{
			mgl64.Vec3{0, 0, 0},
			mgl64.Vec3{1, 0, 0}.Normalize(),
			0,
			0},
		wiOutSide,
		4.0}

	willIntersect, timeToIntersect := data.s.GetIntersect(data.r)
	testCase := bhbosmanTest.MyTestingT{T: t}
	testCase.CheckBool(float64Eq(timeToIntersect, data.where), "Time not correct")
	testCase.CheckBool(willIntersect == data.intersect, "This will Intersection!")
}

func TestSphere0002(t *testing.T) {
	data := struct {
		s         StandardSphere
		r         Ray
		intersect TIntersection
		where     float64
	}{
		StandardSphere{
			Sphere{
				mgl64.Vec3{5.0, 0.0, 0.0},
				1.0},
			//mgl64.Vec3{},
			MaterialParams{}},
		Ray{
			mgl64.Vec3{0, 0, 0},
			mgl64.Vec3{-1, 0, 0}.Normalize(),
			0,
			0},
		wiNoHit,
		0}

	willIntersect, _ := data.s.GetIntersect(data.r)
	testCase := bhbosmanTest.MyTestingT{T: t}
	testCase.CheckBool(willIntersect == data.intersect, "This will not Intersection!")
}

func TestSphere0003(t *testing.T) {
	data := struct {
		s         StandardSphere
		r         Ray
		intersect TIntersection
		where     float64
	}{
		StandardSphere{
			Sphere{
				mgl64.Vec3{1.0, 0.0, 0.0},
				1.0},
			//mgl64.Vec3{},
			MaterialParams{}},
		Ray{
			mgl64.Vec3{0, 0, 0},
			mgl64.Vec3{0, 1, 0}.Normalize(), 0, 0},
		wiOutSide,
		0.0}

	b, ti := data.s.GetIntersect(data.r)
	testCase := bhbosmanTest.MyTestingT{T: t}
	testCase.CheckBool(b == data.intersect, "This will Intersection!")
	testCase.CheckBool(ti == data.where, "This will not Intersection!")
}

func TestSphere0004(t *testing.T) {
	data := struct {
		s         StandardSphere
		r         Ray
		intersect TIntersection
		where     float64
	}{
		StandardSphere{
			Sphere{
				mgl64.Vec3{0.0, 1.0, 0.0},
				1.0},
			//mgl64.Vec3{},
			MaterialParams{}},
		Ray{
			mgl64.Vec3{0, 0, 0},
			mgl64.Vec3{0, 1, 0}.Normalize(), 0, 0},
		wiOutSide,
		0.0}

	b, ti := data.s.GetIntersect(data.r)
	testCase := bhbosmanTest.MyTestingT{T: t}
	testCase.CheckBool(b == data.intersect, "This will Intersection!")
	testCase.CheckFloat64(data.where, ti, "Wrong intersec time")
}

func TestSphere0005(t *testing.T) {
	data := struct {
		s         StandardSphere
		r         Ray
		intersect TIntersection
		where     float64
	}{
		StandardSphere{
			Sphere{
				mgl64.Vec3{5.0, 5.0, 0.0},
				1.0},
			//mgl64.Vec3{},
			MaterialParams{}},
		Ray{
			mgl64.Vec3{0, 0, 0},
			mgl64.Vec3{1, 1, 0}.Normalize(), 0, 0},
		wiOutSide,
		6.07106781186548}

	b, ti := data.s.GetIntersect(data.r)
	testCase := bhbosmanTest.MyTestingT{T: t}
	testCase.CheckBool(b == data.intersect, "This will Intersection!")
	testCase.CheckFloat64(data.where, ti, "Wrong intersec time")
}

func TestSphere0006(t *testing.T) {

	c := &Camera{}
	c.Init2(
		mgl64.Vec3{-2.0, +1.5, 0.0},
		mgl64.Vec3{-2.0, -1.5, 0.0},
		mgl64.Vec3{0.0, 0.0, -1.0}.Normalize(),
		5.0,
		4.0/3.0)

	p := CreatePicture(c, 640, 480)
	r := p.RayAt(320, 80)
	sphere := StandardSphere{
		Sphere{
			mgl64.Vec3{0, 0, -50},
			25},
		//mgl64.Vec3{},
		MaterialParams{}}
	b, time := sphere.GetIntersect(r)

	bhbtest := bhbosmanTest.MyTestingT{T: t}
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
	bhbtest.CheckFloat64(1.0, r.Direction.Len(), "norm")
	bhbtest.CheckBool(b != wiNoHit, "ddddd")
	bhbtest.CheckFloat64(26.27956210851666, time, "ffffff")

}

func TestSphere0007(t *testing.T) {

	c := &Camera{}
	c.Init2(
		mgl64.Vec3{-2.0, +1.5, 0.0},
		mgl64.Vec3{-2.0, -1.5, 0.0},
		mgl64.Vec3{0.0, 0.0, -1.0}.Normalize(),
		5.0,
		4.0/3.0)

	p := CreatePicture(c, 640, 480)
	r := p.RayAt(320, 240)
	sphere := StandardSphere{Sphere{
		mgl64.Vec3{0, 0, -50},
		25},
		//mgl64.Vec3{},
		MaterialParams{}}
	b, time := sphere.GetIntersect(r)

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
	bhbtest.CheckVector(mgl64.Vec3{0, 0, -1}, r.Direction, "A")
	bhbtest.CheckFloat64(1.0, r.Direction.Len(), "norm")
	bhbtest.CheckBool(b != wiNoHit, "ddddd")
	bhbtest.CheckFloat64(25.0, time, "ffffff")

}

func TestSphere0008(t *testing.T) {

	c := &Camera{}
	c.Init2(
		mgl64.Vec3{-2.0, +1.5, 0.0},
		mgl64.Vec3{-2.0, -1.5, 0.0},
		mgl64.Vec3{0.0, 0.0, -1.0}.Normalize(),
		5.0,
		4.0/3.0)

	p := CreatePicture(c, 640, 480)
	r := p.RayAt(320, 240)
	sphere := StandardSphere{Sphere{
		mgl64.Vec3{0, 0, -50},
		25},
		//mgl64.Vec3{},
		MaterialParams{}}
	b, time := sphere.GetIntersect(r)

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
	bhbtest.CheckVector(mgl64.Vec3{0, 0, -1}, r.Direction, "A")
	bhbtest.CheckFloat64(1.0, r.Direction.Len(), "norm")
	bhbtest.CheckBool(b != wiNoHit, "ddddd")
	bhbtest.CheckFloat64(25.0, time, "ffffff")

	pointOfInterest := r.PositionAt(time)
	bhbtest.CheckVector(mgl64.Vec3{0.0, 0.0, -25.0}, pointOfInterest, "")

	r = Ray{
		Origin:    mgl64.Vec3{0, 0, -25},
		Direction: r.Direction}

	b, time = sphere.GetIntersect(r)
	//	fmt.Println("cccccc", time, b)
	//	bhbtest.CheckBool(false, "")

}
