package main

import (. "./core"
		. "./geometry"
		"image"
		"./util"
		"./color")


func main() {

	const WIDTH, HEIGHT int = 1280, 720
	const ASPECT_RATIO float64 = float64(WIDTH) / float64(HEIGHT)

	//Simple Sphere test-scene
	cam_or := Vec{-5, 0, 0}
	//cam_dir := Vec{-1,0,0}.Normalized()
	sp1 := Sphere{Vec{0,0,0}, 1}
	sp2 := Sphere{Vec{0,2,0}, 1.4}
	sp3 := Sphere{Vec{0,-1.5,0}, 0.8}
	var t float64

	sphereArray := [3]Sphere{sp1, sp2, sp3}

	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	//Loop through whole Image
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {

			//Map X and Y from -0.5 to 0.5 and respect Aspect Ratio
			X := 1-float64(x)/float64(WIDTH) - 0.5
			Y := 1-float64(y)/float64(HEIGHT) - 0.5
			X *= ASPECT_RATIO

			//Background Color
			col := color.Gray32(float32(-Y+0.5)*0.5).Clamped()

			//Shoot Ray through each Pixel TODO: Make use of Camera Direction
			ray := Ray{cam_or, Vec{1,X,Y}.Normalized()}

			//TODO: implement in separate function integrator and make Interface for other Objects besides Spheres
			intersect := false

			//tmin = depth of Hitpoint nearest to camera
			var tMin float64 = 10000
			var ObjectHitIndex int
			//Loop thorough Objects
			for i, currentSphere := range sphereArray {

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
				col = color.FromVec(sphereArray[ObjectHitIndex].GetNormal(ray, tMin))
				col.Clamp()
			}

			//convert float32 colors to 24 bit (0-255) color and save
			img.Set(x,y, col.Get24Bit())
		}
	}
	
	
	util.SavePNG("testimage", img)


}
