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

	EPSILON := 0.0000001
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

/*
func (tri Tri) Intersect(ray Ray, t *float64) bool {

	//Planes Normal
	ab := tri.B.Sub(tri.A)
	ac := tri.C.Sub(tri.A)

	//calculate Normal
	normal := Cross(ab, ac)
	//area2 := normal.Length()

	//Find P
	//check if ray & Plane parallel
	nDotRay := Dot(normal, ray.D)

	//TODO find more elegant way for kEpsilon
	if math.Abs(nDotRay) < 0.00000000000001 {
		//don't intersect if ray & Tri are parallel
		return false
	}

	// d Parameter
	d := Dot(normal, tri.A)

	//set t
	*t = (Dot(normal, ray.O) + d) / nDotRay

	// check if the triangle is in behind the ray
	if (*t < 0) {
		return false
	}

	//Ray HitPoint
	p := ray.O.Add(ray.D.Multiply(*t))


	//edge 0
	edge0 := tri.B.Sub(tri.A)
	vp0 := p.Sub(tri.A)
	c := Cross(edge0, vp0)

	if Dot(normal, c) < 0 {
		return false // P is on the right side
	}

	//edge 1
	edge1 := tri.C.Sub(tri.B)
	vp1 := p.Sub(tri.B)
	c = Cross(edge1, vp1)

	if Dot(normal, c) < 0 {
		return false // P is on the right side
	}

	//edge 2
	edge2 := tri.A.Sub(tri.C)
	vp2 := p.Sub(tri.C)
	c = Cross(edge2, vp2)

	if Dot(normal, c) < 0 {
		return false // P is on the right side
	}

	return true
}
*/