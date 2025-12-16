package main

import (
	"errors"
	"fmt"
)

type BigNum struct {
	num     string
	pointer int
}

func (n *BigNum) next() (int, error) {
	n.pointer--
	if n.pointer >= 0 {
		return int(n.num[n.pointer] - '0'), nil
	} else {
		return 1, errors.New("out of range")
	}
}

func main() {
	x := new(BigNum)
	//x.num = "3141592653589793238462643383279502884197169399375105820974944592"
	x.num = "22"
	x.pointer = len(x.num)
	y := new(BigNum)
	//y.num = "2718281828459045235360287471352662497757247093699959574966967627"
	y.num = "16"
	y.pointer = len(y.num)
	result := multiply(x, y, 0, []int{})
	fmt.Println(result)
}

func multiply(x *BigNum, y *BigNum, carry int, result []int) []int {
	lsbX, okX := x.next()
	lsbY, okY := y.next()

	if okX != nil && okY != nil {
		if carry > 0 {
			result = append(result, carry)
		}
		return result
	}

	d := lsbX*lsbY + carry
	result = append(result, d%10)

	if d/10 == 0 {
		return multiply(x, y, 0, result)
	} else {
		return multiply(x, y, d/10, result)
	}
}
