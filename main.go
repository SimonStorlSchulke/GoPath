package main

import (. "./core"
		"./geometry")

func main() {

	const WIDTH, HEIGHT int = 1200, 800
	const ASPECT_RATIO float64 = float64(WIDTH) / float64(HEIGHT)

	//Simple Sphere test-scene
	cam_or := Vec{-5, 0, 0}
	//cam_dir := Vec{-1,0,0}.Normalized() // - TODO
	sp1 := geometry.Sphere{Vec{0,-0.5,0}, 1}
	sp2 := geometry.Sphere{Vec{0,1.5,0}, 1.4}
	sp3 := geometry.Sphere{Vec{0,-2,0}, 0.8}

	ObArray := []geometry.Geometry{&sp1, &sp2, &sp3}

	Render(ObArray, cam_or, WIDTH, HEIGHT, "image")

}
