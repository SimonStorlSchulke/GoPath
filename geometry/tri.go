package geometry

import (
	. "../core"
	//"math"
)

type Tri struct {
	A, B, C Vec
}


func (tri Tri) Normal(ray Ray, HitPoint Vec) Vec {

	ab := tri.B.Sub(tri.A)
	ac := tri.C.Sub(tri.A)
	//calculate Normal
	return Cross(ab, ac)
}

func (tri Tri) Intersect(ray Ray, t *float64) bool {

	EPSILON := 0.0000000001
	var a,f,u,v float64

	edge1 := tri.B.Sub(tri.A)
	edge2 := tri.C.Sub(tri.A)

	h := Cross(ray.D, edge2)
	a = Dot(edge1, h)
	if a > -EPSILON && a > EPSILON {
		return false
	}

	f = 1/a
	s := ray.O.Sub(tri.A)
	u = f * Dot(s, h)
	if u < 0.0 || u > 1.0 {
		return false
	}

	q := Cross(s, edge1)
	v = f * Dot(ray.D, q)
	if v < 0.0 || u + v > 1 {
		return false
	}

	*t = f * Dot(edge2, q)

	if *t > EPSILON {
		//outIntersectionPoint := ray.O.Add(ray.D.Multiply(*t))
		return true
	}

	return true
}
