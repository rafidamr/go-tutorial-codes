package main

import (
	"bufio"
	"fmt"
	"os"
	mmheap "shared/heap"
	"strconv"
	"strings"
)

func main() {
	stream := buidStream("./course2/week3_heap_median/stream.txt")
	res, perfTime := streamOnHeap(&stream)
	fmt.Println(res, perfTime)
}

func buidStream(filename string) []int {
	var stream []int
	ptr, _ := os.Open(filename)
	defer ptr.Close()

	scanner := bufio.NewScanner(ptr)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(strings.TrimSpace(line))
		stream = append(stream, num)
	}
	return stream
}

func streamOnHeap(stream *[]int) (int, int) {
	var s = *stream
	if len(s) == 0 {
		return 0, -1
	}

	var perfTime = 1
	var maxHeap = mmheap.NewHeap(mmheap.MaxMode)
	var minHeap = mmheap.NewHeap(mmheap.MinMode)
	var totalSum int

	for _, num := range s {
		if v, ok := maxHeap.Peek(); ok && v < num {
			minHeap.Push(num)
		} else {
			maxHeap.Push(num)
		}

		// balancing if necessary
		if maxHeap.Len() < minHeap.Len() {
			// odd invariant len(h1) - 1 == len(h2)
			maxHeap.Push(minHeap.Pop())
		} else if maxHeap.Len()-minHeap.Len() > 1 {
			// even invariant len(h1) == len(h2)
			minHeap.Push(maxHeap.Pop())
		}

		if v, ok := maxHeap.Peek(); ok {
			totalSum = totalSum + v
		}
	}

	res := totalSum % 10000

	return res, perfTime
}
