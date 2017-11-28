package main

import (. "./core"
		. "./geometry"
		"fmt"
		"image"
		"./util")


func main() {

	const WIDTH, HEIGHT int = 512, 256

	cam_or := Vec{-5, 0, 0}
	cam_dir := Vec{1, 0, 0}

	sp := Sphere{Vec{0,0,0}, 3}
	t := 0.5

	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))

	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			X := float64(x)/float64(WIDTH)
			Y := float64(y)/float64(HEIGHT)
			
			//Broken
			X -= 0.5
			Y -= 0.5

			ray := Ray{cam_or, Vec{X,Y,0}}
			col := util.RGBFloat1(float32(Y))

			if sp.Intersect(ray, &t) == true {
				col = util.RGBFloat1(1)
			}

			img.Set(x,y, col)
		}
	}
	
	
	util.SavePNG("test", img)

	ray := Ray{cam_or, cam_dir}

	fmt.Print(sp.Intersect(ray, &t))

}