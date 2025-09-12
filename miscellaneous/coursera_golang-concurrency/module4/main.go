package main

import (
	"fmt"
)

func produce(c chan int, abort chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	abort <- 30
	close(c)
}

func iterate(c chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

func firstComeServed(c chan int, abort chan int) {
	for {
		select {
		case a := <-c:
			fmt.Printf("a: %v\n", a)
		case abort <- 20:
			fmt.Println("Assign to abort")
		case x := <-abort:
			fmt.Printf("Value is %v * <-c = %v\n", x, x*(<-c+1))
			fmt.Println("Value is >= 30. Aborted.")
			return
		default:
			fmt.Println("nonblocking default")
		}
	}
}

func f1() {
	c := make(chan int, 5)
	abort := make(chan int, 1)
	go produce(c, abort)
	go firstComeServed(c, abort)
	iterate(c)
}

func main() {
	f1()
}
