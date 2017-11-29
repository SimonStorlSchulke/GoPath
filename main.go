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
	sp1 := Sphere{Vec{0,0,0}, 1.3}
	sp2 := Sphere{Vec{0,2,0}, 1.3}
	sp3 := Sphere{Vec{112,0,10}, 43}
	sp4 := Sphere{Vec{-2,-2,1}, 0.5}
	var t float64

	sphereArray := [4]Sphere{sp1, sp2, sp3, sp4}

	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	//Loop through whole Image
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {

			//Map X and Y from -0.5 to 0.5 and respect Aspect Ratio
			X := float64(x)/float64(WIDTH) - 0.5
			Y := float64(y)/float64(HEIGHT) - 0.5
			X *= ASPECT_RATIO

			//Background Color
			col := color.Gray32(float32(Y)).Clamped()

			//Shoot Ray through each Pixel
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
				//dirty color variation
				col = color.Color32{float32(ObjectHitIndex) / 5, (1-float32(ObjectHitIndex) / 2), float32(ObjectHitIndex) / 2}
			}


			//convert float32 colors to 24 bit (0-255) color and save
			img.Set(x,y, col.Get24Bit())
		}
	}
	
	
	util.SavePNG("test2", img)


}