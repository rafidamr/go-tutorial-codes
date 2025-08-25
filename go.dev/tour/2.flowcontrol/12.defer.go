package main

import "fmt"

func defer_func() {
	fmt.Println(1)
	// pushed two calls into stack
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println(4)
}
