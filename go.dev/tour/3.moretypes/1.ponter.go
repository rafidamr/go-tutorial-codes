package main

import "fmt"

func pointer() {
	var p *float32
	var i, j float32 = 123, 9
	p = &i
	fmt.Println(p)
	fmt.Println(*p)
	*p += 321
	fmt.Println(i)

	p = &j
	*p = *p / 2
	fmt.Println(j)
}
