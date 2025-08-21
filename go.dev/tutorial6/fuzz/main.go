package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("not an utf8 encoded string")
	}
	rArr := []rune(s)
	fmt.Printf("runes: %q\n", rArr)
	for left, right := 0, len(rArr)-1; left < len(rArr)/2; left, right = left+1, right-1 {
		rArr[left], rArr[right] = rArr[right], rArr[left]
	}
	return string(rArr), nil
}

func main() {
	s := "Hello, World"
	r, _ := Reverse(s)
	rr, _ := Reverse(r)
	fmt.Println(s)
	fmt.Println(r)
	fmt.Println(rr)
}
