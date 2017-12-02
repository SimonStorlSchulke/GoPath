package main

import (. "./core"
		"./geometry")

func main() {

	const WIDTH, HEIGHT int = 1200, 800
	const ASPECT_RATIO float64 = float64(WIDTH) / float64(HEIGHT)

	//Simple Sphere test-scene
	cam_or := Vec{-5, 0, 0}
	//cam_dir := Vec{-1,0,0}.Normalized() // - TODO
	sp1 := geometry.Sphere{Vec{0,-0.5,-1}, 1}
	sp2 := geometry.Sphere{Vec{0,1.5,-1}, 1.4}
	sp3 := geometry.Sphere{Vec{0,-2,-1}, 0.8}

	//Cube
	v1 := Vec{0,0,0}
	v2 := Vec{0.6,-0.3,1}
	v3 := Vec{-0.6, 0.3, 1}


	tri1 := geometry.Tri{v2,v1,v3}


	//tri2 := geometry.Tri{Vec{0,0,1}, Vec{0,1,0}, Vec{0,0,1}}
	//tri3 := geometry.Tri{Vec{0,0,1}, Vec{0,1,0}, Vec{0,0,1}}

	ObArray := []geometry.Geometry{&sp1, &sp2, &sp3, &tri1}

	Render(ObArray, cam_or, WIDTH, HEIGHT, "image2")

}
