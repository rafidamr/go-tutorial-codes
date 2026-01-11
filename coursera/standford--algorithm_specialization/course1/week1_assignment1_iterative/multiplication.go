package main

import (
	"fmt"
)

type BigNum struct {
	digits []int
}

func CreateBigNum(s string) *BigNum {
	bn := BigNum{digits: make([]int, len(s))}
	for i := range len(s) {
		bn.digits[len(s)-i-1] = int(s[i] - '0')
	}
	return &bn
}

func main() {
	x := CreateBigNum("3141592653589793238462643383279502884197169399375105820974944592")
	y := CreateBigNum("2718281828459045235360287471352662497757247093699959574966967627")

	result := multiply(x, y)

	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i])
	}
}

func multiply(x *BigNum, y *BigNum) []int {
	result := make([]int, len(x.digits)+len(y.digits))

	for i, a := range x.digits {
		for j, b := range y.digits {
			// a property of middle school multiplication:
			// an index of a result is the sum of operands' indices
			result[i+j] += a * b
		}
	}

	carry := 0
	for i := range result {
		currNumber := carry + result[i]
		result[i] = currNumber % 10
		carry = currNumber / 10
	}

	// traverse from behind until index has value > 0
	i := len(result) - 1
	for head := 0; head == 0; i-- {
		head = result[i]
	}

	return result[:i+2]
}
