package main

import (
	"GoPath/color"
	. "GoPath/core"
	"GoPath/geometry"
	"GoPath/material"
	. "GoPath/render"
	"fmt" //"strconv"
	//"time"
)

func main() {

	const WIDTH, HEIGHT int = 1200, 800
	const ASPECT_RATIO float64 = float64(WIDTH) / float64(HEIGHT)

	//Simple Sphere test-scene
	cam_or := Vec{-5, 0, 0}
	//cam_dir := Vec{-1,0,0}.Normalized() // - TODO
	sp1 := geometry.NewSphere(Vec{0, -0.5, -0.9}, 0.6)
	sp2 := geometry.NewSphere(Vec{0, 1.5, -0.9}, 1)
	sp3 := geometry.NewSphere(Vec{0, -2, -0.9}, 0.5)

	sp2.SetMaterial(material.StandardMaterial{
		Color:     color.GREEN,
		Specular:  0.5,
		Roughness: 0,
		Rim:       0,
	})

	//CuPyramid
	v1 := Vec{-0.5, 0, -0.3}
	v2 := Vec{0.5, -0.75, -0.3}
	v3 := Vec{0.5, 0.75, -0.3}
	v4 := Vec{0, 0, 0.8}

	tri1 := geometry.NewTri(v2, v1, v3)
	tri2 := geometry.NewTri(v4, v2, v1)
	tri3 := geometry.NewTri(v2, v3, v4)
	tri4 := geometry.NewTri(v1, v3, v4)

	ObArray := []geometry.Geometry{sp1, sp2, sp3, tri1, tri2, tri3, tri4}

	/*
		//Benchmark
		numberOfRenders := 20
		start := time.Now()
		for i := 0; i < numberOfRenders; i++ {
			name := "image" + strconv.Itoa(i)
			Render(ObArray, cam_or, WIDTH, HEIGHT, name)
		}
		end := time.Now()
		rendertime := end.Sub(start)
		rendertime /= time.Duration(numberOfRenders)
		fmt.Println("average rendertime on", Threads, "Threads:", rendertime)
	*/
	Render(ObArray, cam_or, WIDTH, HEIGHT, "testimage")
	fmt.Println("Rendering done")
}
