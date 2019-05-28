package geometry

import (
	"GoPath/color"
	. "GoPath/core"
	"math"
)

type Sphere struct {
	CT          Vec
	R           float64
	objectColor color.Color32
}

//Returns Normal Vector of Sphere at Hitpoint
func (S Sphere) Normal(ray Ray, HitPoint Vec) Vec {
	Nhit := HitPoint.Sub(S.CT).Normalized()

	return Nhit
}

func NewSphere(center Vec, radius float64) Sphere {
	return Sphere{center, radius, color.Gray32(0.8)}
}

func (S *Sphere) SetObjectColor(c color.Color32) {
	S.objectColor = c
}

func (S Sphere) ObjectColor() color.Color32 {
	return S.objectColor
}

//returns true if Ray intersects Sphere and sets t to depth of Hitpoint
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
