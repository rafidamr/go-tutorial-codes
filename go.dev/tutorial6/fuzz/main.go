package main

import "fmt"

func Reverse(s string) string {
	bArr := []byte(s)
	for left, right := 0, len(bArr)-1; left < len(bArr)/2; left, right = left+1, right-1 {
		bArr[left], bArr[right] = bArr[right], bArr[left]
	}
	return string(bArr)
}

func main() {
	s := "Hello, World"
	r := Reverse(s)
	rr := Reverse(r)
	fmt.Println(s)
	fmt.Println(r)
	fmt.Println(rr)
}
