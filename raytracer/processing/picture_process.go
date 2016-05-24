package processing

import (
	"fmt"
	"math"
	"os"

	"image/color"
	"image/draw"
	"image/jpeg"

	"github.com/go-gl/mathgl/mgl64"
)

// ImageProcessor is a base class for image processor
type ImageProcessor struct {
	Data draw.Image
}

// SetPixel ...
func (ip *ImageProcessor) SetPixel(x int, y int, c mgl64.Vec3) {
	red := uint8(math.Min(255.0, c[0]*255.0))
	green := uint8(math.Min(255.0, c[1]*255.0))
	blue := uint8(math.Min(255.0, c[2]*255.0))
	ip.Data.Set(
		x,
		y,
		color.RGBA{
			R: red,
			G: green,
			B: blue,
			A: 255})
}

// Save will save the image to file.
func (ip ImageProcessor) Save(filename string) {
	f, err := os.Create("./ss.jpg")
	if err != nil {
		fmt.Println(err)
	} else {
		defer f.Close()
		error := jpeg.Encode(
			f,
			ip.Data,
			&jpeg.Options{
				Quality: 100})
		if error != nil {
			fmt.Println(error)
		}
	}
}
