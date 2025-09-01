package main

import (
	"database/sql"
	"fmt"
	"log"
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

func tests5() {
	nums := []int{1, 2, 3, 4, 5}
	resultChan := make(chan int)
	go processNumbers(nums, resultChan)
	for res := range resultChan {
		fmt.Print(res, " ")
	}
}

func test6() {
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for r := 1; r <= numJobs; r++ {
		<-results
	}
}

func test7() {
	nums := []int{10, 20, 30, 40, 50}
	var sum int
	var mu sync.Mutex
	for _, num := range nums {
		go func() {
			mu.Lock()
			sum += num
			mu.Unlock()
		}()
	}
	time.Sleep(time.Millisecond * 500)
	fmt.Println(sum)
}

func test8() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var user User
	userID := 1
	err = db.QueryRow("SELECT id, username, email FROM users WHERE id = ?",
		userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User ID: %d\nUsername: %s\nEmail: %s\n", user.ID, user.Username,
		user.Email)
}

func test9() {
	rect := Rectangle{Width: 10, Height: 5}
	doubleArea(rect)
	fmt.Println(rect.Width, rect.Height)
}

func main() {
}
