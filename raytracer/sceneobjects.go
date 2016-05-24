package raytracer

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
	bhbosmanVector "github.com/bhbosman/golang/vector"
)

// Material ...
type Material interface {
	GetColor() mgl64.Vec3
	GetPa() float64
	GetPs() float64
	GetPd() float64
	GetRefraction() float64
}

// SceneObjectMaterial ...
type SceneObjectMaterial interface {
	GetMaterial(point mgl64.Vec3) Material
}

// SceneObjectLightCalculation ...
type SceneObjectLightCalculation interface {
	SceneObjectMaterial
	GetNormalAt(point mgl64.Vec3) mgl64.Vec3 // this vector must be normalized
}

// TIntersection ...
type TIntersection int

const (
	wiInside  = -1
	wiNoHit   = 0
	wiOutSide = 1
)

// SceneObjectIntersectCalculation ...
type SceneObjectIntersectCalculation interface {
	GetIntersect(r Ray) (TIntersection, float64)
}

// SceneObject ...
type SceneObject interface {
	Init()
	SceneObjectLightCalculation
	SceneObjectIntersectCalculation
	Description() string
	Rotate(angle float64, axis mgl64.Vec3)
}

// SceneLight ...
type SceneLight interface {
	GetPosition() mgl64.Vec3
	GetColor() mgl64.Vec3
	GetIa() mgl64.Vec3
	GetIs() mgl64.Vec3
	GetId() mgl64.Vec3
	GetF() float64
}

// Scene ...
type Scene struct {
	objs   []SceneObject
	lights []SceneLight
}

// CreateDefaultScene ...
func CreateDefaultScene() *Scene {
	scene := &Scene{}
	scene.AddLight(
		&PointLight{
			Color:    mgl64.Vec3{1, 1, 1},
			Position: mgl64.Vec3{-300, 300, 300},
			B:        500,
			F:        10})
	scene.AddLight(
		&PointLight{
			Color:    mgl64.Vec3{1, 1, 1},
			Position: mgl64.Vec3{+300, 300, 300},
			B:        500,
			F:        10})

	return scene
}

// AddLight ...
func (scene *Scene) AddLight(light SceneLight) {
	scene.lights = append(scene.lights, light)

}

// AddObject ...
func (scene *Scene) AddObject(obj SceneObject) {
	obj.Init()
	scene.objs = append(scene.objs, obj)
}

// FindClosestObject ...
func (scene Scene) FindClosestObject(r Ray) (intersected TIntersection, shortestTime float64, sceneObject SceneObject) {
	shortestTime = math.MaxFloat64
	for _, obj := range scene.objs {
		intersect, intersectTime := obj.GetIntersect(r)
		if intersect != wiNoHit {

			if intersectTime > 0 && intersectTime < shortestTime {
				intersected = intersect
				shortestTime = intersectTime
				sceneObject = obj
			}
		}
	}
	return intersected, shortestTime, sceneObject
}

// CalculateLight ...
func (scene Scene) CalculateLight(sceneObject SceneObject, pointOfInterest mgl64.Vec3, viewerPosition mgl64.Vec3) (color mgl64.Vec3) {
	for _, lightObj := range scene.lights {
		phongLight := PhongLight{}

		lightcolor := phongLight.CalculateLight(
			pointOfInterest,
			lightObj.GetPosition(),
			viewerPosition,
			lightObj,
			sceneObject)
		lightcolor = lightcolor.Mul(1.0 / bhbosmanVector.Distance(lightObj.GetPosition(), pointOfInterest))

		color = color.Add(lightcolor)
	}
	return color

}

// RayIntersect ...
func (scene Scene) RayIntersect(r Ray) (intersected bool, color mgl64.Vec3) {
	intersected, _, color = scene.internalRayIntersect(r, 0, 1.0, 1.0)
	return intersected, color
}

// internal_RayIntersect ...
func (scene Scene) internalRayIntersect(r Ray, depth int, colorValueIntensity float64, arindex float64) (bool, float64, mgl64.Vec3) {
	intersected := false
	dist := 0.0
	color := mgl64.Vec3{}
	localIntersected, shortestTime, sceneObject := scene.FindClosestObject(r)
	if localIntersected != wiNoHit && (sceneObject != nil) {
		intersected = true
		dist = shortestTime
		phit := r.PositionAt(shortestTime)
		nhit := sceneObject.GetNormalAt(phit)
		sceneObjectMaterial := sceneObject.GetMaterial(phit)
		sceneObjectColor := scene.CalculateLight(sceneObject, phit, r.Origin).Mul(colorValueIntensity)

		// reflection
		reflectionColor := mgl64.Vec3{}
		if sceneObject.GetMaterial(phit).GetPs() > 0.0 {
			colorValueIntensity *= (sceneObjectMaterial.GetPs() * 0.80)
			if colorValueIntensity > 1e-08 {
				reflectionDirection := bhbosmanVector.SpecularReflection(nhit, r.Direction)
				reflectionRay := Ray{
					phit.Add(reflectionDirection.Mul(1e-8)),
					reflectionDirection,
					r.x,
					r.y}
				reflectionIntersected, _, localReflectionColor := scene.internalRayIntersect(
					reflectionRay,
					depth+1,
					colorValueIntensity, arindex)
				if reflectionIntersected {
					reflectionColor = localReflectionColor
				}
			}
		}
		// refraction
		refractionColor := mgl64.Vec3{}
		if (sceneObject.GetMaterial(phit).GetRefraction() > 0.0) && (depth < 5) {

			rindex := sceneObjectMaterial.GetRefraction()
			n := arindex / rindex
			N := nhit.Mul(float64(localIntersected))
			cosI := -N.Dot(r.Direction)
			cosT2 := 1.0 - n*n*(1.0-cosI*cosI)
			if cosT2 > 0.0 {
				T := r.Direction.Mul(n).Add(N.Mul(n*cosI - math.Sqrt(cosT2)))
				refractionRay := Ray{phit.Add(T.Mul(1.0e-04)), T, r.x, r.y}
				refractionIntersected, refractionDistance, localRefractionColor := scene.internalRayIntersect(
					refractionRay,
					depth+1,
					colorValueIntensity,
					rindex)
				if refractionIntersected {
					if refractionDistance > 0.0 {
					}

					absorbance := sceneObjectMaterial.GetColor().Mul(0.25 * (-refractionDistance))
					transparency := mgl64.Vec3{math.Exp(absorbance[0]), math.Exp(absorbance[1]), math.Exp(absorbance[2])}
					refractionColor = bhbosmanVector.MultiScalar(localRefractionColor, transparency)
				}
			}

		}
		// this needs some work as reflection and refraction must be on a sliding scale.
		color = color.Add(sceneObjectColor).Add(reflectionColor.Add(refractionColor))
		// color = color.Add(sceneObjectColor).Add(reflectionColor.Mul(1.0)).Add(refractionColor.Mul(1.0))
	}
	return intersected, dist, color
}
