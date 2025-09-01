package main

import "time"

type User struct {
	ID       int
	Username string
	Email    string
}

type Rectangle struct {
	Width  int
	Height int
}

func processNumbers(nums []int, ch chan<- int) {
	for _, num := range nums {
		ch <- num * 2
	}
	close(ch)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func doubleArea(r Rectangle) {
	r.Width *= 2
	r.Height *= 2
}
