package geometry

import (
	"GoPath/color"
	. "GoPath/core"
)

//TODO: Figure out how I can use different Objects in an Array for the Scene
type Geometry interface {
	Intersect(ray Ray, t *float64) bool
	Normal(ray Ray, HitPoint Vec) Vec
	ObjectColor() color.Color32
}
