package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/bhbosman/golang/raytracer"
	"github.com/bhbosman/golang/raytracer/processing/processing03"
	"github.com/go-gl/mathgl/mgl64"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2) // Use all the machine's cores
	scene := raytracer.CreateDefaultScene()
	scene00002(scene)
	camera := raytracer.DefaultCamera()
	picture := raytracer.CreatePicture(camera, 2560, 1920)
	ss := processing03.CreateImageProcessor(picture, scene)
	t0 := time.Now()
	ss.Process(8, 12)
	t1 := time.Now()
	fmt.Println(t1.Sub(t0))
	ss.Save("")
}

func scene00004(scene *raytracer.Scene) {
	for x := -2; x <= 2; x++ {
		for y := -2; y <= 2; y++ {
			for z := -10; z <= -6; z++ {
				scene.AddObject(&raytracer.StandardSphere{
					Sphere: raytracer.Sphere{
						Origin: mgl64.Vec3{
							float64(x * 20),
							float64(y * 20),
							float64(z * 20)},
						Radius: 10},
					StandardMaterial: raytracer.MaterialParams{
						RefractionValue: 0,
						Pd:              0.15,
						Ps:              1.0,
						Color:           mgl64.Vec3{0, 1, 0}}})
			}
		}
	}
}

func scene00003(scene *raytracer.Scene) {

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{0, 0, 0},
			Radius: 40},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 5,
			Pd:              0.15,
			Ps:              1.0,
			Color:           mgl64.Vec3{0, 1, 0}}})

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{0, 0, 0},
			Radius: 5},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 0,
			Pd:              0.5,
			Ps:              1.0,
			Color:           mgl64.Vec3{1, 0, 0}}})

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{10, 0, 0},
			Radius: 5},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 0,
			Pd:              0.5,
			Ps:              0.0,
			Color:           mgl64.Vec3{1, 0, 0}}})

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{-10, 0, 0},
			Radius: 5},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 0,
			Pd:              0.5,
			Ps:              0.0,
			Color:           mgl64.Vec3{1, 0, 0}}})

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{0, -10, 0},
			Radius: 5},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 0,
			Pd:              0.5,
			Ps:              0.0,
			Color:           mgl64.Vec3{1, 0, 0}}})

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{0, 10, 0},
			Radius: 5},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 0,
			Pd:              0.5,
			Ps:              0.0,
			Color:           mgl64.Vec3{1, 0, 0}}})

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{0, 0, -90},
			Radius: 45},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 0,
			Pd:              0.3,
			Ps:              0.0,
			Color:           mgl64.Vec3{1, 0, 0}}})

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{90, 0, -90},
			Radius: 45},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 0,
			Pd:              0.3,
			Ps:              0.0,
			Color:           mgl64.Vec3{1, 0, 0}}})

	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{-90, 0, -90},
			Radius: 45},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 0,
			Pd:              0.3,
			Ps:              0.0,
			Color:           mgl64.Vec3{1, 0, 0}}})
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

func scene00001(scene *raytracer.Scene) {
	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{0, 0, -50},
			Radius: 30},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 1,
			Pd:              0.3,
			Ps:              .4,
			Color:           mgl64.Vec3{0, 0, 1}}})

	//*
	scene.AddObject(&raytracer.StandardSphere{
		Sphere: raytracer.Sphere{
			Origin: mgl64.Vec3{0, 0, -50},
			Radius: 50},
		StandardMaterial: raytracer.MaterialParams{
			RefractionValue: 1,
			Pd:              0.5,
			Ps:              .0,
			Color:           mgl64.Vec3{1, 0, 0}}})
	//*/
	//*
	for i := 1; i < 5; i++ {
		for k := 1; k < 5; k++ {

			scene.AddObject(&raytracer.StandardSphere{
				Sphere: raytracer.Sphere{
					Origin: mgl64.Vec3{float64(i*40 - 100), float64(k*40 - 100), -100},
					Radius: 20},
				StandardMaterial: raytracer.MaterialParams{
					RefractionValue: 1.2,
					Pd:              0.7,
					Ps:              0.5,
					Color:           mgl64.Vec3{0, 1, 0}}})
		}
	}
	//*/
	/*
		scene.AddObject(&raytracer.StandardPlane{
			Plane: raytracer.Plane{
				Origin: mgl64.Vec3{0, -50, 0},
				Normal: mgl64.Vec3{0,-1, 0}},
			StandardMaterial: raytracer.MaterialParams{
				Pd:    1.0,
				Ps:    1.0,
				Color: mgl64.Vec3{1, 0, 0}}})
		//*/
	//*
	scene.AddObject(&raytracer.BlackAndWhiteChessBoard{
		Plane: raytracer.Plane{
			Origin: mgl64.Vec3{0, -100, 0},
			Normal: mgl64.Vec3{0, -1, 0}},
		BlockSize: mgl64.Vec2{50, 50},
		Color01:   nil,
		Color02:   nil})
	//*/

}
