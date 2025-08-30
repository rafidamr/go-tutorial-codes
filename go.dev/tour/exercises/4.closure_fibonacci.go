package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	arr := make([]int, 0, 5)
	return func() int {
		l := len(arr)
		if l == 0 || l == 1 {
			arr = append(arr, l)
			return l
		}
		n := arr[l-1] + arr[l-2]
		arr = append(arr, n)
		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(i, ":", f())
	}
}
