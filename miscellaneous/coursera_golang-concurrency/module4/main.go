package main

import (
	"fmt"
	"os"
	"sync"
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

func increment(i *int, wg *sync.WaitGroup) {
	*i = *i + 1
	wg.Done()
}

func incrementMutex(i *int, wg *sync.WaitGroup, mt *sync.Mutex) {
	mt.Lock()
	*i = *i + 1
	mt.Unlock()
	wg.Done()
}

func main() {
	switch os.Args[1] {
	case "1":
		// Synchornized channel communication
		c := make(chan int, 5)
		abort := make(chan int, 1)
		go produce(c, abort)
		go firstComeFirstServe(c, abort)
		iterate(c)
	case "2":
		// Mutex
		var i = 0
		var j = 0
		var wg sync.WaitGroup
		var mt sync.Mutex
		c := 100
		wg.Add(c * 2)
		for ci := 0; ci < c; ci++ {
			go increment(&i, &wg)
			go incrementMutex(&j, &wg, &mt)
		}
		wg.Wait()
		if i != c {
			fmt.Printf("i != %v: %v\n", c, i)
		}
		if j != c {
			fmt.Printf("j != %v: %v\n", c, j)
		}
	}
}
