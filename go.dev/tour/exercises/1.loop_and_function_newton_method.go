package main

import (
	"fmt"
	"math"
)

func SqrtType1(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func SqrtType2(x float64) float64 {
	z := 1.0
	diff := 1.0
	i := 0
	for math.Abs(diff) > 1e-10 {
		diff = (z*z - x) / (2 * z)
		z -= diff
		i++
	}
	fmt.Printf("SqrtType2 takes %v iteration(s)\n", i)
	return z
}

func main() {
	var x = []float64{1, 2, 3, 40000, 9000000, 18000000}
	for _, v := range x {
		// fmt.Printf("Sqrt=%v math.Sqrt=%v\n", Sqrt(v), math.Sqrt(v))
		// fmt.Printf("Diff=%v\n", SqrtType1(v)-math.Sqrt(v))
		fmt.Printf("Diff2=%v\n", SqrtType2(v)-math.Sqrt(v))
	}
}
