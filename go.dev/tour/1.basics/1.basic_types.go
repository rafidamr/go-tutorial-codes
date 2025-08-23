package main

import (
	"fmt"
	"math/cmplx"
)

var (
	Bool   bool       = true
	MaxInt uint64     = 1<<64 - 1
	z1     complex128 = 5 + 12i
	z2     complex128 = cmplx.Sqrt(z1)
)

func main() {
	fmt.Printf("Type=%T Value=%v\n", Bool, Bool)
	fmt.Printf("Type=%T Value=%v\n", MaxInt, MaxInt)
	fmt.Printf("Type=%T Value=%v\n", z1, z1)
	fmt.Printf("Type=%T Value=%v\n", z2, z2)
}
