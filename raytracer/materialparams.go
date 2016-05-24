package raytracer



import (
	"github.com/go-gl/mathgl/mgl64"
)

// MaterialParams ...
type MaterialParams struct {
	Pa float64
	Ps float64
	Pd float64
	Color mgl64.Vec3
	RefractionValue float64
}

// GetColor ...
func(mp MaterialParams) GetColor() mgl64.Vec3 {
	return mp.Color
	
}


// GetPa ...
func(mp MaterialParams) GetPa() float64 {
	return mp.Pa
}

// GetPd ...
func(mp MaterialParams) GetPd() float64 {
	return mp.Pd
}

/// GetPs ...
func(mp MaterialParams) GetPs() float64 {
	return mp.Ps
}

// GetRefraction ...
func (mp MaterialParams) GetRefraction() float64 {
	return mp.RefractionValue
}

