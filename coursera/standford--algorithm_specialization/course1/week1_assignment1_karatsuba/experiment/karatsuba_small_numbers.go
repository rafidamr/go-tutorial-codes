package smallnum

import (
	"fmt"
	"math"
)

func smallnum() {
	//	prev week
	//	init x and y
	//	recursion: split into two
	//		get high and low
	//		need to construct something like this: high*100^2 + middle*100^1 + low^100^0
	//		for some reason the middle starts from (a + b)(c + d) = ac + ad + bc + bd
	//		then middle: ad + bc = (a + b)(c + d) - ac - bd
	//		recurse into high and low while returning high*100^2 + middle*100^1 + low^100^0
	//

	//26/12/2025
	//init x and y
	//check and return if x or y < 2 digits
	//split high low from x is a and b
	//split high low from y is c and d
	//
	//given ab*cd is the template1 (a*10m+b)*(c*10m+d) = a*c*10^2m + (a*d+b*c)*10^m + b*d
	//switch them to template2 (a*d+b*c)*10^m = (a*10m+b)*(c*10m+d) - a*c*10^2m - b*d
	//
	//compute ac and bd and (i dont understand why) (a+b)*(c+d) into recursion
	//combine the result back into template2
	//combine ac, bd, and template2 into template1
	//return result

	var x uint = 11
	var y uint = 123

	result := karatsuba(x, y)

	fmt.Println(result)
}

func karatsuba(x uint, y uint) uint {
	if x < 10 || y < 10 {
		return x * y
	}

	maxDigit := max(countDigit(x), countDigit(y))
	splitPoint := maxDigit / 2
	divisor := uint(math.Pow(10, float64(splitPoint)))

	a := x / divisor
	b := x % divisor

	c := y / divisor
	d := y % divisor

	ac := karatsuba(a, c)
	bd := karatsuba(b, d)
	abcd := karatsuba(a+b, c+d) // this i dont understand

	mid := abcd - ac - bd // this i dont understand

	m1 := uint(math.Pow(10, float64(2*splitPoint)))
	m2 := uint(math.Pow(10, float64(splitPoint)))

	result := ac*m1 + mid*m2 + bd
	return result
}

func countDigit(i uint) uint {
	var count uint = 0
	for i > 0 {
		count++
		i /= 10
	}
	return count
}
