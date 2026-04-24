package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// need to construct array because iterating over ht map is ~2x slower
	ht, arr := buildDataStructure("course2/week4_hash_table/numbers.txt")
	count := syncCount2Sum(ht, arr)
	// count := exhConcurCount2Sum(ht, arr)
	// count := anConcurCount2Sum(ht, arr)

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
