package main

import (
	"fmt"
	"math"
)

func struct_func() {
	type Point struct {
		X int
		Y int
	}
	point := Point{3, 4}
	var ptr *Point = &point
	ptr.X = int(math.Sqrt(float64(ptr.X*ptr.X + ptr.Y*ptr.Y)))
	fmt.Println(*ptr)

	x := Point{10, 11}
	y := &Point{Y: 5}
	fmt.Println(x, y)
}
