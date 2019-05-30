package core

import (
	"fmt"
	"math"
)

type Vec struct {
	X, Y, Z float64
}

//Adds Vector a + Vector b and returns it
func (a *Vec) Add(b Vec) Vec {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z

	return *a
}

//Subs Vector a - Vector b and returns it
func (a *Vec) Sub(b Vec) Vec {
	a.X -= b.X
	a.Y -= b.Y
	a.Z -= b.Z

	return *a
}

//Returns the Length of a Vector as float64
func (v *Vec) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

//Normalizes an Vector and returns it
func (v *Vec) Normalize() Vec {
	len := v.Length()
	if len > 0 {
		invLen := 1 / len
		v.X *= invLen
		v.Y *= invLen
		v.Z *= invLen
	}
	return *v
}

//Method - returns normalized Vector
func (a Vec) Normalized() Vec {
	len := a.Length()
	V := a

	if len > 0 {
		invLen := 1 / len
		V.X *= invLen
		V.Y *= invLen
		V.Z *= invLen
	}
	return V
}

//Method - Negated
func (a Vec) Negated() Vec {
	V := a
	V.X = -a.X
	V.Y = -a.Y
	V.Z = -a.Z
	return V
}

func Dot(a, b Vec) float64 {
	//function - Dot Product
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Cross(a, b Vec) Vec {
	//function - Cross Product
	return Vec{a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X}
	//TODO: Eventuell normalized() zurückgeben
}

//Multiplies a Vector with a float64 and returns it.
func (v *Vec) Multiply(f float64) Vec {
	v.X *= f
	v.Y *= f
	v.Z *= f

	return *v
}

//Prints X,Y,Z and Length of a Vector
func (v *Vec) Info() {
	fmt.Printf("x=%.3f y=%.3f z=%.3f length %.3f \n", v.X, v.Y, v.Z, v.Length())
}
