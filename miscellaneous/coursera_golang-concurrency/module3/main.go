package main

import (
	"fmt"
	"sync"
	"time"
)

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

func main() {
	f2()
}
