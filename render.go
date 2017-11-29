package main

import (
	"./geometry"
	"image"
	"./util"
	"./color"
	. "./core")


func render(ObjectArray []geometry.Sphere, cam_or Vec) {

	const WIDTH, HEIGHT int = 1280, 720
	const ASPECT_RATIO float64 = float64(WIDTH) / float64(HEIGHT)

	var t float64
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))

	//Loop through whole Image
	//TODO: make Interface for other Objects besides Spheres

	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {

			//Map X and Y from -0.5 to 0.5 and respect Aspect Ratio 1- for correct rotation
			X := 1-float64(x)/float64(WIDTH) - 0.5
			Y := 1-float64(y)/float64(HEIGHT) - 0.5
			X *= ASPECT_RATIO

			//Background Color
			col := color.Gray32(float32(-Y+0.5)*0.5).Clamped()

			//Shoot Ray through each Pixel TODO: Make use of Camera Direction
			ray := Ray{cam_or, Vec{1,X,Y}.Normalized()}

			intersect := false

			//tmin = depth of Hitpoint nearest to camera
			var tMin float64 = 10000
			var ObjectHitIndex int
			//Loop thorough Objects
			for i, currentSphere := range ObjectArray {

				//Check for intersection in every Object
				if currentSphere.Intersect(ray, &t) == true {
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
				col = color.Gray32(float32(ObjectHitIndex) / 5)
				//Show Surface Normal
				col = color.FromVec(ObjectArray[ObjectHitIndex].GetNormal(ray, tMin))
				col.Clamp()
			}

			//convert float32 colors to 24 bit (0-255) color and save
			img.Set(x,y, col.Get24Bit())
		}
	}
	util.SavePNG("testimage", img)
}
