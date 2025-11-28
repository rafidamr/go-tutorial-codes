package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Using Go version 1.25.4 darwin/arm64

type CStick struct{ sync.Mutex }
type Phil struct {
	LeftChopstick, RightChopstick *CStick
	name                          string
	EatingProcess                 sync.Mutex
}

func (p *Phil) eat(semaphore chan struct{}) {
	// Use mutex to avoid a same philosopher to eat multiple times at once
	p.EatingProcess.Lock()

	// Acquire and defer release of semaphore
	semaphore <- struct{}{}
	defer func() { <-semaphore }()

	// Each philosopher pick the chopstick in any order
	if rand.Intn(2) == 0 {
		p.LeftChopstick.Lock()
		fmt.Println(p.name, "took the left")
		p.RightChopstick.Lock()
		fmt.Println(p.name, "took the right")
	} else {
		p.RightChopstick.Lock()
		fmt.Println(p.name, "took the right")
		p.LeftChopstick.Lock()
		fmt.Println(p.name, "took the left")
	}

	fmt.Printf("starting to eat %s\n", p.name)
	fmt.Printf("finishing eating %s\n", p.name)

	p.LeftChopstick.Unlock()
	p.RightChopstick.Unlock()
	p.EatingProcess.Unlock()
}

func main() {
	// Initialize 5 philosophers and 5 chopsticks, numbered from 1 to 5
	num := 5
	philArr := make([]*Phil, num)
	cstickArr := make([]*CStick, num)
	for i := 0; i < num; i++ {
		cstickArr[i] = new(CStick)
	}
	for i := 0; i < num; i++ {
		philArr[i] = &Phil{
			name:           fmt.Sprint(i + 1),
			LeftChopstick:  cstickArr[i],
			RightChopstick: cstickArr[(i+1)%num]}
	}

	// Each philosopher eat 3 times. There should be 5 * 3 goroutines.
	turn := 3
	var wg sync.WaitGroup
	// Declare semaphore in the host goroutine to limit max 2 philosophers eating at once
	semaphore := make(chan struct{}, 2)
	for t := 0; t < turn; t++ {
		for i := 0; i < num; i++ {
			wg.Go(func() {
				philArr[i].eat(semaphore)
			})
		}
	}
	wg.Wait()
	fmt.Println("All finished eating")
}
