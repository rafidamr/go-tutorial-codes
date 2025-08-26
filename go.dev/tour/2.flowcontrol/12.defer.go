package main

import "fmt"

func defer_func() {
	fmt.Print(1)
	// pushed two calls into stack
	defer fmt.Print(2)
	defer fmt.Print(3)
	fmt.Print(4)
}
