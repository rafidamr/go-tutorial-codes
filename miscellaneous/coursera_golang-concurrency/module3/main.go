package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("New goroutine")
	fmt.Println("Main goroutine")
	time.Sleep(1000 * time.Millisecond)
}
