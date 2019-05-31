package render

import (
	"GoPath/color"
	. "GoPath/core"
	"GoPath/geometry"
	"GoPath/util"
	"image"
	"math"
)

func Render(objectArray []geometry.Geometry, cam_or Vec, WIDTH, HEIGHT int, filename string) {

	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	ASPECT_RATIO := float64(WIDTH) / float64(HEIGHT)
	var ray Ray
	var intersect bool
	var t float64

	//HARDCODED LIGHTS

	lights := []Light{
		NewLightDirectional(Vec{0.7, -0.4, -0.6}, color.Color32{1, 0.9, 0.6}, 3),
		NewLightDirectional(Vec{-0.3, 0.4, -0.2}, color.Color32{0.8, 0.8, 1}, 6),
	}

	//Loop through whole Image
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {

			//Map X and Y from -0.5 to 0.5 and respect Aspect Ratio 1- for correct rotation
			X := 1 - float64(x)/float64(WIDTH) - 0.5
			Y := 1 - float64(y)/float64(HEIGHT) - 0.5
			X *= ASPECT_RATIO

			//Background Color
			col := color.Gray32(float32(0.5-Y) * 0.5).Clamped()

			//Shoot Ray through each Pixel TODO: Make use of Camera Direction
			ray = Ray{cam_or, Vec{1, X, Y}.Normalized()}

			intersect = false

			//tmin = depth of Hitpoint nearest to camera
			var tMin float64 = 10000
			var ObjectHitIndex int

			//Loop thorough Objects
			var cObj geometry.Geometry

			for i, currentObject := range objectArray {

				//Check for intersection in every Object
				if currentObject.Intersect(ray, &t) {
					intersect = true
					//check if new t is smaller then old tmin and set tmin to new t if it is
					tminOld := tMin

					if t < tminOld {
						tMin = t
						//what Object got hit
						ObjectHitIndex = i
						cObj = currentObject
					}
				}
			}

			if intersect == true {
				HitPoint := ray.O.Add(ray.D.Multiply(tMin))
				Normal := objectArray[ObjectHitIndex].Normal(ray, HitPoint)

				col = Diffuse(cObj, lights, HitPoint, Normal)

				col.Clamp()
			}
			//convert float32 colors to 24 bit (0-255) color and save
			img.Set(x, y, col.Get24Bit())
		}
	}
	util.SavePNG(filename, img)
}

//TODO REFACTOR
func Diffuse(cObj geometry.Geometry, lights []Light, hitPoint, Normal Vec) color.Color32 {

	col := color.Color32{}

	for _, L := range lights {
		c1 := color.Divide(cObj.Material().BaseColor(), color.Gray32(3.14))
		cMax := float32(math.Max(0.0, Dot(Normal, L.Dir().Negated())))
		c2 := color.MultiplyF(L.Color(), L.Intinsity()*cMax)
		col = color.Add(col, color.Multiply(c1, c2))
	}
	return col
}

type Light interface {
	Dir() Vec
	Color() color.Color32
	Intinsity() float32
}

type LightDirectional struct {
	dir       Vec
	color     color.Color32
	intensity float32
}

func NewLightDirectional(dir Vec, color color.Color32, intensity float32) *LightDirectional {
	return &LightDirectional{dir.Normalized(), color, intensity}
}

func (L *LightDirectional) Dir() Vec {
	return L.dir
}

func (L *LightDirectional) Color() color.Color32 {
	return L.color
}

func (L *LightDirectional) Intinsity() float32 {
	return L.intensity
}
