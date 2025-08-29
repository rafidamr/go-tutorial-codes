package main

import "fmt"

func interfaces2_func() {
	var i interface{} = 10
	fmt.Printf("v=%v T=%T\n", i, i)

	i = "a long str"
	s, ok := i.(string)
	fmt.Println(s, ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)

	switch t := i.(type) {
	case string:
		fmt.Println("it's a string")
	default:
		fmt.Printf("v=%v T=%T\n", t, t)
	}
}
