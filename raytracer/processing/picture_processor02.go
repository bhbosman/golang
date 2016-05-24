package processing

import (
	"image"
	"github.com/bhbosman/golang-raytracer"
	"github.com/go-gl/mathgl/mgl64"
)


// ImageProcessor02PixelTask is a struct that describes the pixel that needs to be processed
type ImageProcessor02PixelTask struct {
	x int
	y int
}

// ImageProcessor02ColorTask is a struct that describes the color value of the pixel
type ImageProcessor02ColorTask struct {
	x int
	y int
	c mgl64.Vec3
}

// ImageProcessor02 is the structure for the second processor
type ImageProcessor02 struct {
	ImageProcessor
	picture *raytracer.Picture
	scene *raytracer.Scene
}

// CreateImageProcessor02 creates an instance of ImageProcessor02
func CreateImageProcessor02(p *raytracer.Picture, s *raytracer.Scene) ImageProcessor02 {
	result := ImageProcessor02{
		ImageProcessor{
			image.NewRGBA(image.Rect(0, 0, p.GetWidth(), p.GetHeight()))},
		p,
		s}

	return result
}

const imageProcessor02Workers = 8

// Process generates the image
func (ip ImageProcessor02) Process() {
	operationCount := ip.picture.GetWidth() * ip.picture.GetHeight()
	PixelTaskChannel := make(chan ImageProcessor02PixelTask, operationCount)
	ColorTaskChannel := make(chan ImageProcessor02ColorTask, operationCount)
	DoneChannel := make(chan bool, operationCount)

	for i := 0; i <= imageProcessor02Workers-1; i++ {
		go ip.DealWithPixel(PixelTaskChannel, ColorTaskChannel, DoneChannel)
	}
	go func() {
		for x := 0; x < ip.picture.GetWidth(); x++ {
			for y := 0; y < ip.picture.GetHeight(); y++ {
				pixelTask := ImageProcessor02PixelTask{x, y}
				PixelTaskChannel <- pixelTask
			}
		}
		close(PixelTaskChannel)
	}()
	go func() {
		for task := range ColorTaskChannel {
			ip.SetPixel(task.x, task.y, task.c)
			DoneChannel <- true
		}
	}()
	for i := 0; i < operationCount; i++ {
		<-DoneChannel
	}
}

// DealWithPixel is the go function
func (ip ImageProcessor02) DealWithPixel(
	pixelTasks <-chan ImageProcessor02PixelTask,
	colorTasks chan<- ImageProcessor02ColorTask,
	doneTasks chan<- bool) {

	for task := range pixelTasks {
		r :=  ip.picture.RayAt(task.x, task.y)
		b, c := ip.scene.RayIntersect(r)
		if b {
			colorTasks <- ImageProcessor02ColorTask{task.x, task.y, c}
		} else {
			doneTasks <- false
		}
	}
}
