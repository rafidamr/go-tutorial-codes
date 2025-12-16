package main

import (
	"fmt"
)

type BigNum struct {
	digits []int
}

func NewBigNum(s string) *BigNum {
	digits := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		digits[len(s)-1-i] = int(s[i] - '0')
	}
	return &BigNum{digits: digits}
}

func main() {
	x := NewBigNum("3141592653589793238462643383279502884197169399375105820974944592")
	y := NewBigNum("2718281828459045235360287471352662497757247093699959574966967627")

	result := multiply(x, y)

	// Print result in correct order (reverse)
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i])
	}
}

func multiply(x *BigNum, y *BigNum) []int {
	// Result can be at most len(x) + len(y) digits
	result := make([]int, len(x.digits)+len(y.digits))

	// Multiply each digit of x by each digit of y
	for i := 0; i < len(x.digits); i++ {
		for j := 0; j < len(y.digits); j++ {
			product := x.digits[i] * y.digits[j]
			result[i+j] += product
		}
	}

	// Handle carries
	carry := 0
	for i := 0; i < len(result); i++ {
		result[i] += carry
		carry = result[i] / 10
		result[i] %= 10
	}

	// Remove leading zeros
	for len(result) > 1 && result[len(result)-1] == 0 {
		result = result[:len(result)-1]
	}

	return result
}
