package main

import (. "./core"
		. "./geometry")

func main() {

	const WIDTH, HEIGHT int = 1280, 720
	const ASPECT_RATIO float64 = float64(WIDTH) / float64(HEIGHT)

	//Simple Sphere test-scene
	cam_or := Vec{-5, 0, 0}
	//cam_dir := Vec{-1,0,0}.Normalized()
	sp1 := Sphere{Vec{0,0,0}, 1}
	sp2 := Sphere{Vec{0,2,0}, 1.4}
	sp3 := Sphere{Vec{0,-1.5,0}, 0.8}
	sphereArray := []Sphere{sp1, sp2, sp3}

	render(sphereArray, cam_or)
}
