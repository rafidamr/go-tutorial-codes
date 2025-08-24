package main

import (
	"fmt"
	"math"
)

func type_conversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z1 uint8 = uint8(f * 100000)
	var z2 uint = -uint(f)
	var z3 uint = uint(-f)
	fmt.Printf("f=%v z1=%v z2=%v z3=%v\n", f, z1, z2, z3)
	fmt.Printf("z2 == z3: %v\n", z2 == z3)
}
