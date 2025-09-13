package main

import (
	"fmt"
)

func produce(c chan int, abort chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	abort <- 1
	close(c)
}

func iterate(c chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

func firstComeFirstServe(c chan int, abort chan int) {
	for {
		select {
		case a := <-c:
			fmt.Printf("a: %v\n", a)
		case abort <- 20:
			fmt.Println("Assign to abort")
		case x := <-abort:
			y := x * (<-c + 1)
			fmt.Printf("y = x * (<-c + 1) = %v\n", y)
			if y >= 30 {
				fmt.Println("y >= 30. Aborted.")
				return
			}
		default:
			fmt.Println("nonblocking default")
		}
	}
}

func f1() {
	c := make(chan int, 5)
	abort := make(chan int, 1)
	go produce(c, abort)
	go firstComeFirstServe(c, abort)
	iterate(c)
}

func main() {
	f1()
}
