package main

import (. "./core")

func main() {

	a := Vec{1,1,1}
	b := Vec{-2,1,7}
	a.Add(b)
	a.Normalize()
	c := Vec{0,0,0}
	c.Normalize()
	a.Info()
	c.Info()
}