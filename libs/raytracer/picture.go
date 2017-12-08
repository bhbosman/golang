package raytracer

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
		pixelWidth:          width,
		pixelHeight:         height,
		pixelWidthDistance:  camera.GetWidth() / float64(width),
		pixelHeightDistance: camera.GetHeight() / float64(height),
		camera:              camera}

	return p
}

// RayAt will return the ray that will be used the coordinate. This ray will be projected into the screen
func (p *Picture) RayAt(x, y int) Ray {
	deltaX := p.camera.direction.x.Mul(float64(x) * p.pixelWidthDistance)
	deltaY := p.camera.direction.y.Mul(float64(y) * p.pixelHeightDistance)
	o := p.camera.plane.topLeft.Add(deltaX).Sub(deltaY)
	r := o.Sub(p.camera.viewPoint).Normalize()
	return Ray{
		Origin:    o,
		Direction: r}
}

// GetWidth returns the pixel width of the picture
func (p *Picture) GetWidth() int {
	return p.pixelWidth
}

// GetHeight returns the pixel height of the picture
func (p *Picture) GetHeight() int {
	return p.pixelHeight
}
