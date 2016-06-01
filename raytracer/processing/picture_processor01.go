package processing

import (
	"image"

	"github.com/bhbosman/golang/raytracer"
)

// ImageProcessor01 is the structure for the second processor
type ImageProcessor01 struct {
	ImageProcessor
}

// CreateImageProcessor01 creates an instance of ImageProcessor02
func CreateImageProcessor01(p raytracer.Picture) ImageProcessor01 {
	result := ImageProcessor01{
		ImageProcessor{
			image.NewRGBA(image.Rect(0, 0, p.GetWidth(), p.GetHeight()))}}

	return result
}

// Process generates the image
func (ip ImageProcessor01) Process(p raytracer.Picture, s raytracer.Scene) {
	for x := 0; x < p.GetWidth(); x++ {
		for y := 0; y < p.GetHeight(); y++ {
			r := p.RayAt(x, y)
			b, c := s.RayIntersect(r)
			if b {
				ip.SetPixel(x, y, c)
			}
		}
	}
}
