package main

import (
	"fmt"
)

func type_conversion() {
	var x, y int = 3, 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	var f float64 = float64(x*x + y*y)
	var z1 uint = uint(f)
	var z2 uint = -uint(f)
	var z3 uint = uint(-f)
	fmt.Println(f, z1, z2, z3)
	fmt.Printf("z2 == z3: %v\n", z2 == z3)
}
