package main

import (. "./core"
		. "./geometry"
		"fmt"
		"image"
		"./util"
		"./color")


func main() {

	const WIDTH, HEIGHT int = 1280, 720
	const ASPECT_RATIO float64 = float64(WIDTH) / float64(HEIGHT)

	//Simple Sphere test-scene
	cam_or := Vec{-5, 0, 0}
	cam_dir := Vec{1, 0, 0}
	sp := Sphere{Vec{0,0,0}, 1.3}
	t := 0.5

	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	//Loop through whole Image
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {

			//Map X and Y from -0.5 to 0.5 and respect Aspect Ratio
			X := float64(x)/float64(WIDTH) - 0.5
			Y := float64(y)/float64(HEIGHT) - 0.5
			X *= ASPECT_RATIO

			//Shoot Ray through each Pixel
			ray := Ray{cam_or, Vec{1,X,Y}.Normalized()}

			//Background Color
			col := color.Gray32(float32(Y)).Clamped()

			//colorize Sphere
			if sp.Intersect(ray, &t) == true {
				col = color.Color32{0.4,0,0.6}
			}

			//convert float32 colors to 24 bit (0-255) color and save
			img.Set(x,y, col.Get24Bit())
		}
	}
	
	
	util.SavePNG("test", img)

	ray := Ray{cam_or, cam_dir}

	fmt.Println(sp.Intersect(ray, &t))

}