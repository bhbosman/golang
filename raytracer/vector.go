package raytracer

import (
	"fmt"
	"math"
	"github.com/go-gl/mathgl/mgl64"
)

// SpecularReflection ...
func SpecularReflection(n mgl64.Vec3, d mgl64.Vec3) mgl64.Vec3 {
	return d.Sub(n.Mul(2.0 * d.Dot(n))).Normalize()
}

// SpecularRefraction ...
func SpecularRefraction(n mgl64.Vec3, d mgl64.Vec3, n1 float64, n2 float64) (bool, mgl64.Vec3) {
	dn := d.Dot(n)
	underSquare := 1.0 - math.Pow(n1, 2.0)/math.Pow(n2, 2.0)*(1.0-math.Pow(dn, 2))
	if underSquare < 0.0 {
		fmt.Println("zzz")
		return false, mgl64.Vec3{0.0, 0.0, 0.0}
	}
	term01 := d.Sub(n.Mul(dn)).Mul(n1 / n2)
	term02 := n.Mul(math.Sqrt(underSquare))

	return true, term01.Sub(term02)
}


// Transform ...
func Transform(v mgl64.Vec3) mgl64.Mat4 {
	ident := mgl64.Ident4()
	c1, c2, c3, _ := ident.Cols()
	return mgl64.Mat4FromCols(c1, c2, c3, v.Vec4(1))
}

// MultiScalar ...
func MultiScalar(v, m mgl64.Vec3) mgl64.Vec3 {
	return mgl64.Vec3{m[0] * v[0], m[1] * v[1], m[2] * v[2]}
}

// Distance returns the Euclidean distance between v and ov.
func Distance(v mgl64.Vec3, ov mgl64.Vec3) float64 {
	d := v.Sub(ov)
	return math.Sqrt(d.Dot(d))
}
