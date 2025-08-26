package main

import (
	"fmt"
	"strings"
)

func mapsFunc() {
	s := "string1 string2 string3 string1 string3"
	parts := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range parts {
		m[string(v)] += 1
	}
	fmt.Println(m)
	delete(m, "string1")
	fmt.Println(m)
	val, found := m["nonExistent"]
	fmt.Println(val, found)
}
