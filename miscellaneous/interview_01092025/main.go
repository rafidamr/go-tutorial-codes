package main

import (
	"fmt"
	"sync"
	"time"
)

func test1() {
	nums := []int{1, 2, 3, 4, 5}
	for i := range nums {
		nums[i] = nums[i] * 2
	}
	fmt.Println(nums)
}

func test2(n int) int {
	if n <= 1 {
		return n
	}
	return test2(n-1) + test2(n-2)
}

func test3() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("Goroutine %d completed\n", i)
		}()
	}
	wg.Wait()
}

func test4(x, y int) (sum int, product int) {
	sum = x + y
	product = x * y
	return
}

func processNumbers(nums []int, ch chan<- int) {
	for _, num := range nums {
		ch <- num * 2
	}
	close(ch)
}

func tests5() {
	nums := []int{1, 2, 3, 4, 5}
	resultChan := make(chan int)
	go processNumbers(nums, resultChan)
	for res := range resultChan {
		fmt.Print(res, " ")
	}
}

func main() {
}
