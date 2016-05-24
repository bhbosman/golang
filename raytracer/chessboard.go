package raytracer

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
)




// BlackAndWhiteChessBoard represent a Chessboard
type BlackAndWhiteChessBoard struct {
	Plane
	BlockSize mgl64.Vec2
	Color01   Material
	Color02   Material
	Transform mgl64.Mat4
}

// Description ...
func (p BlackAndWhiteChessBoard) Description() string {
	return "BlackAndWhiteChessBoard"
}
// GetMaterial ...
func (p BlackAndWhiteChessBoard) GetMaterial(point mgl64.Vec3) Material {
	newPos := mgl64.TransformCoordinate(point, p.Transform)

	x := math.Abs((math.Ceil(newPos[0] / p.BlockSize[0])))
	y := math.Abs((math.Ceil(newPos[1] / p.BlockSize[1])))

	r := (int(x + y)) % 2
	if r == 1 {
		return &MaterialParams{
			Pa: 0,
			Ps: 1,
			Pd: 0.2,
			Color: mgl64.Vec3{1, 1, 1}}
	} else {
		return &MaterialParams{
			Pa: 0,
			Ps: 1,
			Pd: 0.2,
			Color: mgl64.Vec3{0, 0, 0}}
	}
}
// Init ...
func (p *BlackAndWhiteChessBoard) Init() {
	p.Normal = p.Normal.Normalize()
	angle := math.Acos(p.Normal.Dot(mgl64.Vec3{0, 0, 1})) / 2
	axis := p.Normal.Cross(mgl64.Vec3{0, 0, 1}).Normalize()
	p.Transform = mgl64.HomogRotate3D(angle, axis)

	//p.Rotate(mgl64.DegToRad(90), mgl64.Vec3{1,0,0})
}

