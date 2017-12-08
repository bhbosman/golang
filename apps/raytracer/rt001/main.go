package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/bhbosman/code/golang/libs/raytracer"
	"github.com/bhbosman/code/golang/libs/raytracer/processing/processing03"
	"github.com/go-gl/mathgl/mgl64"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2) // Use all the machine's cores
	scene := raytracer.CreateDefaultScene()
	scene00002(scene)
	camera := raytracer.DefaultCamera()
	picture := raytracer.CreatePicture(camera, 800, 600)
	ss := processing03.CreateImageProcessor(picture, scene)
	t0 := time.Now()
	ss.Process(8, 12)
	t1 := time.Now()
	fmt.Println(t1.Sub(t0))
	ss.Save("")
}

func scene00002(scene *raytracer.Scene) {
	scene.AddObject(&raytracer.BlackAndWhiteChessBoard{
		Plane: raytracer.Plane{
			Origin: mgl64.Vec3{0, -25, 0},
			Normal: mgl64.Vec3{0, -1, 0}},
		BlockSize: mgl64.Vec2{50, 50},
		Color01:   nil,
		Color02:   nil})

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{0, 0, 0},
			Radius: 25},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 1.5,
			Pd:              0.3,
			Ps:              0.0,
			Color:           mgl64.Vec3{0, 0, 1}}})

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{-50, 0, 0},
			Radius: 25},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 1.5,
			Pd:              0.3,
			Ps:              0.5,
			Color:           mgl64.Vec3{1, 0, 0}}})

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{50, 0, 0},
			Radius: 25},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 1.5,
			Pd:              0.3,
			Ps:              0.5,
			Color:           mgl64.Vec3{0, 1, 0}}})

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{25, 50, 0},
			Radius: 25},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 0,
			Pd:              0.3,
			Ps:              0.5,
			Color:           mgl64.Vec3{0, 1, 0}}})

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{-25, 50, 0},
			Radius: 25},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 1.0,
			Pd:              0.3,
			Ps:              0.0,
			Color:           mgl64.Vec3{0, 1, 0}}})

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{25, 50, -50},
			Radius: 25},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 0,
			Pd:              0.3,
			Ps:              0.5,
			Color:           mgl64.Vec3{0, 1, 0}}})

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{-25, 50, -50},
			Radius: 25},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 0,
			Pd:              0.3,
			Ps:              0.5,
			Color:           mgl64.Vec3{0, 1, 0}}})
}
