package core

import ("math" 
		"fmt")

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

//Returns the Length of a Vector as float64
func (v *Vec) Length() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

//Normalizes an Vector and returns it
func (v *Vec) Normalize() Vec {

	if v.Length() > 0 {
		invLen := 1/v.Length()
		v.X *= invLen
		v.Y *= invLen
		v.Z *= invLen
	}
	return *v
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