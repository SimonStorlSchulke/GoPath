package geometry

import (. "../core"
		"math")

type Sphere struct {
	CT Vec
	R float64
}

//Returns Normal Vector of Sphere at Hitpoint
func (S *Sphere) getNormal(ray Ray) Vec {
	v := ray.O.Sub(S.CT)
	v.Normalize()
	return v
}

//returns true if Ray intersects Sphere and sets t to depth of Hitpoint
func (S *Sphere) Intersect(ray Ray, t *float64) bool {
	oc := ray.O.Sub(S.CT)
	b := 2 * Dot(oc, ray.D)
	c := Dot(oc, oc) - S.R * S.R
	discriminant := b * b - 4 * c

	if discriminant < 0 {
		return false
	}
	sqrt_discriminant := math.Sqrt(discriminant)
	t0 := -b - sqrt_discriminant
	t1 := -b + sqrt_discriminant

	*t = math.Min(t0, t1)

	return true
}