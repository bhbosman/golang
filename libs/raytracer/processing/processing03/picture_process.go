package processing03

import (
	"image"

	"github.com/bhbosman/code/golang/libs/raytracer"
	"github.com/bhbosman/code/golang/libs/raytracer/processing"
	"github.com/go-gl/mathgl/mgl64"
)

// PixelTask is a struct that describes the pixel that needs to be processed
type PixelTask struct {
	x      int
	y      int
	width  int
	height int
}

// ColorTask is a struct that describes the color value of the pixel
type ColorTask struct {
	x      int
	y      int
	width  int
	height int
	c      []mgl64.Vec3
}

// ImageProcessor is the structure for the second processor
type ImageProcessor struct {
	processing.ImageProcessor
	picture *raytracer.Picture
	scene   *raytracer.Scene
}

// CreateImageProcessor creates an instance of ImageProcessor02
func CreateImageProcessor(p *raytracer.Picture, s *raytracer.Scene) ImageProcessor {
	result := ImageProcessor{
		processing.ImageProcessor{
			Data: image.NewRGBA(image.Rect(0, 0, p.GetWidth(), p.GetHeight()))},
		p,
		s}

	return result
}

// Workers ...
const Workers = 8

// Process generates the image
func (ip ImageProcessor) Process(width int, height int) {
	imageWidth := ip.picture.GetWidth() - (ip.picture.GetWidth() % width)
	imageHeight := ip.picture.GetHeight() - (ip.picture.GetHeight() % height)

	blockX := imageWidth / width
	blockY := imageHeight / height

	operationCount := blockX * blockY

	PixelTaskChannel := make(chan PixelTask, operationCount)
	ColorTaskChannel := make(chan ColorTask, operationCount)
	DoneChannel := make(chan bool, operationCount)

	for i := 0; i <= Workers-1; i++ {
		go ip.DealWithPixel(PixelTaskChannel, ColorTaskChannel)
		go ip.DealWithColor(ColorTaskChannel, DoneChannel)
	}
	go func() {
		for y := 0; y < blockY; y++ {
			for x := 0; x < blockX; x++ {
				pixelTask := PixelTask{
					x:      x * width,
					y:      y * height,
					width:  width,
					height: height}
				PixelTaskChannel <- pixelTask
			}
		}
		close(PixelTaskChannel)
	}()
	for i := 0; i < operationCount; i++ {
		<-DoneChannel
	}
}

// DealWithColor ...
func (ip ImageProcessor) DealWithColor(colorTasks <-chan ColorTask, doneTasks chan<- bool) {
	for task := range colorTasks {
		for y := 0; y < task.height; y++ {
			offsetY := y * task.width
			for x := 0; x < task.width; x++ {
				ip.SetPixel(
					task.x+x,
					task.y+y,
					task.c[offsetY+x])
			}
		}
		doneTasks <- true
	}
}

// DealWithPixel is the go function
func (ip ImageProcessor) DealWithPixel(
	pixelTasks <-chan PixelTask,
	colorTasks chan<- ColorTask) {

	for task := range pixelTasks {
		colorTask := ColorTask{
			x:      task.x,
			y:      task.y,
			width:  task.width,
			height: task.height}
		colorTask.c = make(
			[]mgl64.Vec3,
			task.width*task.height,
			task.width*task.height)

		for y := 0; y < task.height; y++ {
			offsetY := y * task.width
			for x := 0; x < task.width; x++ {
				r := ip.picture.RayAt(
					task.x+x,
					task.y+y)
				_, c := ip.scene.RayIntersect(r)
				colorTask.c[offsetY+x] = c
			}
		}
		colorTasks <- colorTask
	}
}
