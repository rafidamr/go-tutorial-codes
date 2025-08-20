package main

import "fmt"

type Number interface { // Type Constraint
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"f": 34,
		"s": 12,
	}

	floats := map[string]float64{
		"f": 35.98,
		"s": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumsInts(ints),
		SumsFloats(floats),
	)

	fmt.Printf("Generic Sums: %v and %v\n",
		SumsInt64orFloat64[string, int64](ints),
		SumsInt64orFloat64[string, float64](floats),
	)

	fmt.Printf("Generic Sums with contraint: %v and %v\n",
		SumsNumbers(ints),
		SumsNumbers(floats),
	)
}

func SumsInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumsFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumsInt64orFloat64[K comparable, Vtype int64 | float64](m map[K]Vtype) Vtype {
	var s Vtype
	for _, v := range m {
		s += v
	}
	return s
}

func SumsNumbers[K comparable, Vtype Number](m map[K]Vtype) Vtype {
	var s Vtype
	for _, v := range m {
		s += v
	}
	return s
}
