package main

import (
	"./geometry"
	"image"
	"./util"
	"./color"
	. "./core"
)

func Render(objectArray []geometry.Geometry, cam_or Vec, WIDTH, HEIGHT int, filename string) {
	
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))	
	ASPECT_RATIO := float64(WIDTH) / float64(HEIGHT)
	var ray Ray
	var intersect bool
	var t float64

	//Loop through whole Image
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {

			//Map X and Y from -0.5 to 0.5 and respect Aspect Ratio 1- for correct rotation
			X := 1 - float64(x)/float64(WIDTH) - 0.5
			Y := 1 - float64(y)/float64(HEIGHT) - 0.5
			X *= ASPECT_RATIO

			//Background Color
			col := color.Gray32(float32(0.5 - Y) * 0.5).Clamped()

			//Shoot Ray through each Pixel TODO: Make use of Camera Direction
			ray = Ray{cam_or, Vec{1, X, Y}.Normalized()}

			intersect = false

			//tmin = depth of Hitpoint nearest to camera
			var tMin float64 = 10000
			var ObjectHitIndex int

			//Loop thorough Objects
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
					}
				}
			}

			if intersect == true {
				HitPoint := ray.O.Add(ray.D.Multiply(tMin))

				//Light Attendance
				LightPos := Vec{-4, -1, 3}
				LightAttendance := Dot(HitPoint.Normalized(), LightPos.Normalized())

				//Placeholder Color
				//Color Spheres (color multiplied with Light Attendance)
				col = color.Gray32(float32(ObjectHitIndex) / 5)
				col = color.Gray32(float32(LightAttendance))
				//col = color.FromVec(objectArray[ObjectHitIndex].Normal(ray, HitPoint))
				col.Clamp()
			}
			//convert float32 colors to 24 bit (0-255) color and save
			img.Set(x, y, col.Get24Bit())
		}
	}
	util.SavePNG(filename, img)
}