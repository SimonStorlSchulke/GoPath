package color

import (
	"GoPath/core"
	"GoPath/util"
	"image/color"
)

//Color with RGB values from 0 to 1
type Color32 struct {
	R, G, B float32
}

//clamps RGB to values between 0 and 1
func (col *Color32) Clamp() {
	col.R = util.Clamp32(col.R, 0, 1)
	col.G = util.Clamp32(col.G, 0, 1)
	col.B = util.Clamp32(col.B, 0, 1)
}

//returns clamped (0-1) Color
func (col Color32) Clamped() Color32 {
	col.R = util.Clamp32(col.R, 0, 1)
	col.G = util.Clamp32(col.G, 0, 1)
	col.B = util.Clamp32(col.B, 0, 1)
	return col
}

//returns a grey color based on a 0-1 value
func Gray32(val float32) Color32 {
	return Color32{val, val, val}
}

//returns a color from a Vector
func FromVec(v core.Vec) Color32 {
	return Color32{float32(v.X), float32(v.Y), float32(v.Z)}
}

//returns a standard library 24Bit Color for saving
func (col *Color32) Get24Bit() color.Color {
	return color.RGBA{uint8(col.R * 255), uint8(col.G * 255), uint8(col.B * 255), 255}
}
