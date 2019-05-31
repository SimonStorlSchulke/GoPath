package color

//different Blendmodes for Color32

func Multiply(c1, c2 Color32) Color32 {
	return Color32{c1.R * c2.R, c1.G * c2.G, c1.B * c2.B}
}

func MultiplyF(c1 Color32, val float32) Color32 {
	return Color32{c1.R * val, c1.G * val, c1.B * val}
}

func Divide(c1, c2 Color32) Color32 {
	return Color32{c1.R / c2.R, c1.G / c2.G, c1.B / c2.B}
}

func DivideF(c1 Color32, val float32) Color32 {
	return Color32{c1.R / val, c1.G / val, c1.B / val}
}

func Add(c1, c2 Color32) Color32 {
	return Color32{c1.R + c2.R, c1.G + c2.G, c1.B + c2.B}
}

func Subtract(c1, c2 Color32) Color32 {
	return Color32{c1.R - c2.R, c1.G - c2.G, c1.B - c2.G}
}

func Mix(c1, c2 Color32, fac float32) Color32 {
	return Color32{
		(1-fac)*c1.R + fac*c2.R,
		(1-fac)*c1.G + fac*c2.G,
		(1-fac)*c1.B + fac*c2.B}
}
