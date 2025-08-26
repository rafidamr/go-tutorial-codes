package main

import "fmt"

func pointer() {
	var p *int
	var i = 123
	p = &i
	fmt.Println(p)
	fmt.Println(*p)
	*p += 321
	fmt.Println(i)
}
