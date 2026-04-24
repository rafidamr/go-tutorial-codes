package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func main() {
	// need to construct array because iterating over ht map is ~2x slower
	ht, arr := buildDataStructure("course2/week4_hash_table/numbers.txt")
	// count := syncCount2Sum(ht, arr)
	// count := exhConcurCount2Sum(ht, arr)
	count := anConcurCount2Sum(ht, arr)

	// The correct answer is 427
	fmt.Println(count)
}

func buildDataStructure(filename string) (*map[int64]bool, *[]int64) {
	var arr []int64
	ht := make(map[int64]bool)
	ptr, _ := os.Open(filename)
	defer ptr.Close()

	scanner := bufio.NewScanner(ptr)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(strings.TrimSpace(line))
		ht[int64(num)] = true
		arr = append(arr, int64(num))
	}
	return &ht, &arr
}

// Without concurrency, time: 259.25s user and 4:20.26 total
func syncCount2Sum(hashTable *map[int64]bool, array *[]int64) uint {
	var arr = *array
	var ht = *hashTable
	var count uint
	var t int64

	for t = -10000; t <= 10000; t++ {
		for _, x := range arr {
			var y = t - x
			if ht[y] && y != x {
				count++
				break
			}
		}
	}

	return count
}

// Exhaustive concurrency, time: 514.66s user 2.62s system 948% cpu 54.511 total
func exhConcurCount2Sum(hashTable *map[int64]bool, array *[]int64) uint {
	var arr = *array
	var ht = *hashTable
	var count uint

	var wg sync.WaitGroup
	var m sync.Mutex

	for t := int64(-10000); t <= 10000; t++ {
		wg.Go(func() {
			for _, x := range arr {
				var y = t - x
				if ht[y] && y != x {
					m.Lock()
					count++
					m.Unlock()
					break
				}
			}

		})
	}

	wg.Wait()

	return count
}

// Another concurrency, time: 493.46s user 2.40s system 934% cpu 53.067 total
func anConcurCount2Sum(hashTable *map[int64]bool, array *[]int64) uint64 {
	var arr = *array
	var ht = *hashTable
	var count uint64

	var workerNum = runtime.NumCPU()
	var tasks = make(chan int64, workerNum)
	var wg sync.WaitGroup

	for w := 0; w < workerNum; w++ {
		wg.Go(func() {
			var local uint64
			for t := range tasks {
				for _, x := range arr {
					var y = t - x
					if ht[y] && y != x {
						local++
						break
					}
				}
			}
			atomic.AddUint64(&count, local)
		})
	}

	for t := int64(-10000); t <= 10000; t++ {
		tasks <- t
	}

	close(tasks)
	wg.Wait()

	return count
}
