package main

import "fmt"

type I interface {
	M()
}

type T struct {
	Field int
}

func (s *T) M() {
	if s == nil {
		fmt.Println("nil value")
		return
	}
	fmt.Println(s.Field)
}

func interfaces_func() {
	var i I
	fmt.Printf("%v %T\n", i, i)
	// i.M() // error because of nil interface

	var t *T
	i = t
	fmt.Printf("%v %T\n", i, i)
	i.M()

	i = &T{123}
	fmt.Printf("%v %T\n", i, i)
	i.M()
}
