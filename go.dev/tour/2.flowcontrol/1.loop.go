package main

import (
	"fmt"
	"math"
)

func looping() {
	var sum int
	for i := 0; sum < 1000; i++ {
		sum += i
	}
	fmt.Println(sum)
	sum = 0
	for sum < 1000 { // Similar to while
		sum++
	}
	fmt.Println(sum)
	fmt.Println(fmt.Sprint(math.Sqrt(-(-float64(sum)))) + "i")

	if x := math.Sqrt(70); x < 6 {
		fmt.Println(x)
	} else {
		fmt.Printf("%g >= %g\n", x, float64(6))
	}
}
