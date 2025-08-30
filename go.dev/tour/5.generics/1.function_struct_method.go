package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l List[T]) PrintVal() {
	fmt.Println(l.val)
}

func PrintVal[X any](l *List[X]) {
	fmt.Println(l.val)
}

func main() {
	l := &List[int]{val: 10}
	l.next = &List[int]{val: 30}
	l2 := &List[string]{val: "content"}

	l.PrintVal()
	l = l.next
	PrintVal(l)
	l2.PrintVal()
}
