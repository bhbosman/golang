package raytracer

// import "github.com/go-gl/mathgl/mgl64"

// Picture struct represents a picture
type Picture struct {
	pixelWidth          int
	pixelHeight         int
	pixelWidthDistance  float64
	pixelHeightDistance float64
	camera              *Camera
}

// CreatePicture creates a Picture object
func CreatePicture(camera *Camera, width int, height int) *Picture {
	p := &Picture{
		width,
		height,
		camera.GetWidth() / float64(width),
		camera.GetHeight() / float64(height),
		camera}

	return p
}

// RayAt will return the ray that will be used the coordinate. This ray will be projected into the screen
func (p *Picture) RayAt(x, y int) Ray {
	deltaX := p.camera.direction.x.Mul(float64(x) * p.pixelWidthDistance)
	deltaY := p.camera.direction.y.Mul(float64(y) * p.pixelHeightDistance)
	o := p.camera.plane.topLeft.Add(deltaX).Sub(deltaY)
	r := o.Sub(p.camera.viewPoint).Normalize()
	return Ray{o, r, x, y}
}

// GetWidth returns the pixel width of the picture
func (p *Picture) GetWidth() int {
	return p.pixelWidth
}

// GetHeight returns the pixel height of the picture
func (p *Picture) GetHeight() int {
	return p.pixelHeight
}
