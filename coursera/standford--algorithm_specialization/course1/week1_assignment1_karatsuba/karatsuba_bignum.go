package main

import (
	"fmt"
	"strings"
)

func main() {
	// Test with very large numbers that can't fit in uint64
	x := "3141592653589793238462643383279502884197169399375105820974944592"
	y := "2718281828459045235360287471352662497757247093699959574966967627"

	result := karatsubaBigNum(x, y)
	fmt.Println("Result:", result)

	// Test with smaller numbers to verify correctness
	fmt.Println("\nVerification with small numbers:")
	fmt.Println("11 * 123 =", karatsubaBigNum("11", "123"))
	fmt.Println("Expected: 1353")
}

// karatsubaBigNum multiplies two very large numbers represented as strings
// Uses the same Karatsuba algorithm as the original implementation
func karatsubaBigNum(x string, y string) string {
	// Remove leading zeros
	x = removeLeadingZeros(x)
	y = removeLeadingZeros(y)

	// Base case: if either number is single digit, use simple multiplication
	if len(x) <= 1 || len(y) <= 1 {
		return multiplySmallNumbers(x, y)
	}

	// Find the maximum digit count and split point
	maxDigit := max(len(x), len(y))
	splitPoint := maxDigit / 2

	// Split x into high (a) and low (b) parts
	// For example: 1234 splits into a=12, b=34
	a, b := splitNumber(x, splitPoint)

	// Split y into high (c) and low (d) parts
	c, d := splitNumber(y, splitPoint)

	// Recursive calls - same as original algorithm
	ac := karatsubaBigNum(a, c)           // high * high
	bd := karatsubaBigNum(b, d)           // low * low
	abcd := karatsubaBigNum(addStrings(a, b), addStrings(c, d)) // (a+b)*(c+d)

	// Calculate middle term: (a+b)*(c+d) - ac - bd = ad + bc
	mid := subtractStrings(abcd, ac)
	mid = subtractStrings(mid, bd)

	// Combine results: ac*10^(2*splitPoint) + mid*10^splitPoint + bd
	result := addStrings(
		addStrings(
			multiplyByPowerOf10(ac, 2*splitPoint),
			multiplyByPowerOf10(mid, splitPoint),
		),
		bd,
	)

	return result
}

// splitNumber splits a number string at the given position from the right
// For example: splitNumber("12345", 2) returns ("123", "45")
func splitNumber(num string, splitPoint int) (string, string) {
	// Pad with zeros on the left if needed
	if len(num) < splitPoint {
		num = strings.Repeat("0", splitPoint-len(num)) + num
	}

	splitPos := len(num) - splitPoint
	if splitPos <= 0 {
		return "0", num
	}

	high := num[:splitPos]
	low := num[splitPos:]

	if high == "" {
		high = "0"
	}
	if low == "" {
		low = "0"
	}

	return high, low
}

// addStrings adds two numbers represented as strings
func addStrings(a string, b string) string {
	// Make sure a is the longer string
	if len(a) < len(b) {
		a, b = b, a
	}

	result := make([]byte, 0, len(a)+1)
	carry := 0
	i := len(a) - 1
	j := len(b) - 1

	// Add from right to left
	for i >= 0 || j >= 0 || carry > 0 {
		digitA := 0
		if i >= 0 {
			digitA = int(a[i] - '0')
			i--
		}

		digitB := 0
		if j >= 0 {
			digitB = int(b[j] - '0')
			j--
		}

		sum := digitA + digitB + carry
		carry = sum / 10
		result = append(result, byte(sum%10+'0'))
	}

	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return removeLeadingZeros(string(result))
}

// subtractStrings subtracts b from a (assumes a >= b)
func subtractStrings(a string, b string) string {
	// Pad b with zeros on the left to match length of a
	if len(b) < len(a) {
		b = strings.Repeat("0", len(a)-len(b)) + b
	}

	result := make([]byte, 0, len(a))
	borrow := 0
	i := len(a) - 1

	// Subtract from right to left
	for i >= 0 {
		digitA := int(a[i] - '0')
		digitB := 0
		if i < len(b) {
			digitB = int(b[i] - '0')
		}

		diff := digitA - digitB - borrow
		if diff < 0 {
			diff += 10
			borrow = 1
		} else {
			borrow = 0
		}

		result = append(result, byte(diff+'0'))
		i--
	}

	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return removeLeadingZeros(string(result))
}

// multiplySmallNumbers multiplies single-digit or small numbers
// This is the base case for the recursion
func multiplySmallNumbers(a string, b string) string {
	if a == "0" || b == "0" {
		return "0"
	}

	// Convert to integers (safe for small numbers)
	numA := 0
	for _, ch := range a {
		numA = numA*10 + int(ch-'0')
	}

	numB := 0
	for _, ch := range b {
		numB = numB*10 + int(ch-'0')
	}

	product := numA * numB
	return fmt.Sprintf("%d", product)
}

// multiplyByPowerOf10 multiplies a number by 10^power by appending zeros
func multiplyByPowerOf10(num string, power int) string {
	if num == "0" || power == 0 {
		return num
	}
	return num + strings.Repeat("0", power)
}

// removeLeadingZeros removes leading zeros from a number string
func removeLeadingZeros(num string) string {
	if num == "" {
		return "0"
	}

	i := 0
	for i < len(num)-1 && num[i] == '0' {
		i++
	}

	return num[i:]
}