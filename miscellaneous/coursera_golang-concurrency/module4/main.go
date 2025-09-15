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

func initAndExecute(waitGroup *sync.WaitGroup, on *sync.Once) {
	on.Do(func() {
		fmt.Println("Init")
	})
	fmt.Println("Execute")
	waitGroup.Done()
}

type CStick struct{ sync.Mutex }
type Phil struct {
	LeftChopstick, RightChopstick *CStick
	name                          string
}

func (p Phil) eat(wg *sync.WaitGroup) {
	// p.LeftChopstick.Lock()
	// p.RightChopstick.Lock()
	fmt.Printf("%s: Iam eating\n", p.name)
	// p.LeftChopstick.Unlock()
	// p.RightChopstick.Unlock()
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
	case "3":
		// Synchronized initialization
		var wg sync.WaitGroup
		var on sync.Once
		wg.Add(2)
		go initAndExecute(&wg, &on)
		go initAndExecute(&wg, &on)
		go initAndExecute(&wg, &on)
		wg.Wait()
	case "4.1":
		// Deadlock: circular dependencies
		var wg sync.WaitGroup
		c1 := make(chan int)
		c2 := make(chan int)
		wg.Go(func() {
			c1 <- 1
			<-c2
		})
		wg.Go(func() {
			c2 <- 1
			<-c1
		})
		wg.Wait()
	case "4.2":
		// Deadlock (The Dining Philosopher)
		num := 3
		cstickArr := make([]*CStick, num)
		philArr := make([]*Phil, num)
		for i := 0; i < num; i++ {
			cstickArr[i] = new(CStick)
		}
		for i := 0; i < num; i++ {
			philArr[i] = &Phil{
				name:           fmt.Sprintf("Phil %v-th", i),
				LeftChopstick:  cstickArr[i],
				RightChopstick: cstickArr[(i+1)%num]}
		}
		var wg sync.WaitGroup
		wg.Add(1)
		wg.Add(1)
		for i := 0; i < num; i++ {
			fmt.Println(i)
			go philArr[i].eat(&wg)
		}
		wg.Wait()
		fmt.Println("All finished eating")
	}
}
