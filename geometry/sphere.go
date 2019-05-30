package geometry

import (
	"GoPath/color"
	. "GoPath/core"
	"GoPath/material"
	"math"
)

//Sphere primitive
type Sphere struct {
	CT          Vec
	R           float64
	objectColor color.Color32
	material    material.Material
}

//Returns Normal Vector of Sphere at Hitpoint
func (S *Sphere) Normal(ray Ray, HitPoint Vec) Vec {
	Nhit := HitPoint.Sub(S.CT).Normalized()

	return Nhit
}

//SetMaterial defines a Material for the Sphere
func (S *Sphere) SetMaterial(mat material.Material) {
	S.material = mat
}

//Material returns the Spheres Material
func (S *Sphere) Material() material.Material {
	return S.material
}

//NewSphere creates a new Sphere with default Material
func NewSphere(center Vec, radius float64) *Sphere {
	return &Sphere{
		CT:          center,
		R:           radius,
		objectColor: color.Gray32(0.8),
		material:    material.Default()}
}

//SetObjectColor sets an ObjectColor
func (S *Sphere) SetObjectColor(c color.Color32) {
	S.objectColor = c
}

//ObjectColor returns Objectcolor
func (S *Sphere) ObjectColor() color.Color32 {
	return S.objectColor
}

//Intersect returns true if Ray intersects Sphere and sets t to depth of Hitpoint
func (S Sphere) Intersect(ray Ray, t *float64) bool {
	L := S.CT.Sub(ray.O)
	tca := Dot(L, ray.D)
	d2 := Dot(L, L) - tca*tca

	if d2 > (S.R * S.R) {
		return false
	}

	thc := math.Sqrt((S.R * S.R) - d2)

	t0 := tca - thc
	t1 := tca + thc
	*t = math.Min(t0, t1)

	return true
}
