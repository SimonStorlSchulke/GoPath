package main

import (
	. "./core"
	"./geometry"
	. "./render"
)

func main() {

	const WIDTH, HEIGHT int = 1200, 800
	const ASPECT_RATIO float64 = float64(WIDTH) / float64(HEIGHT)

	//Simple Sphere test-scene
	cam_or := Vec{-5, 0, 0}
	//cam_dir := Vec{-1,0,0}.Normalized() // - TODO
	sp1 := geometry.Sphere{Vec{0, -0.5, -1.8}, 0.6}
	sp2 := geometry.Sphere{Vec{0, 1.5, -1.8}, 1}
	sp3 := geometry.Sphere{Vec{0, -2, -1.8}, 0.5}

	//Cube
	v1 := Vec{-0.5, 0, -0.3}
	v2 := Vec{0.5, -0.75, -0.3}
	v3 := Vec{0.5, 0.75, -0.3}
	v4 := Vec{0, 0, 0.8}

	tri1 := geometry.Tri{v2, v1, v3}
	tri2 := geometry.Tri{v4, v2, v1}
	tri3 := geometry.Tri{v2, v3, v4}
	tri4 := geometry.Tri{v1, v3, v4}

	//tri2 := geometry.Tri{Vec{0,0,1}, Vec{0,1,0}, Vec{0,0,1}}
	//tri3 := geometry.Tri{Vec{0,0,1}, Vec{0,1,0}, Vec{0,0,1}}

	ObArray := []geometry.Geometry{&sp1, &sp2, &sp3, &tri1, &tri2, &tri3, &tri4}

	Render(ObArray, cam_or, WIDTH, HEIGHT, "image")
}
