package main

import "fmt"

func array_func() {
	var arr = [7]int{1, 2, 3, 4}
	defer fmt.Println(arr)

	s1 := arr[0:2]
	s2 := arr[1:]
	fmt.Println(s1, s2)
	s2[0] = 100
	fmt.Println(s1, s2)
	s3 := arr[:]
	s3[0] = 100000
	fmt.Println(arr)
}
