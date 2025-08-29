package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	diff := 1.0
	i := 0
	for math.Abs(diff) > 1e-10 {
		diff = (z*z - x) / (2 * z)
		z -= diff
		i++
	}
	return z, nil
}

func error_func() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
