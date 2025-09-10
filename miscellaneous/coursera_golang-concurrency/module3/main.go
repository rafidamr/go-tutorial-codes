package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	f3()
}

func prod(c chan int) {
	c <- 1
	fmt.Println("Reached")
}

func con(c chan int) {
	<-c
}

func f3() {
	// c := make(chan int) // this will block the second p call
	c := make(chan int, 3)
	go con(c)
	go prod(c)
	go prod(c)
	time.Sleep(100 * time.Millisecond)
}

func f2go(wg *sync.WaitGroup, str string) {
	fmt.Println(str)
	wg.Done()
}

func f2() {
	var wg sync.WaitGroup
	wg.Add(2)
	go f2go(&wg, "task1")
	go f2go(&wg, "task2")
	wg.Wait()
	fmt.Println("Main goroutine")
}

func f1() {
	go fmt.Println("New goroutine")
	fmt.Println("Main goroutine")
	time.Sleep(1000 * time.Millisecond)
}
