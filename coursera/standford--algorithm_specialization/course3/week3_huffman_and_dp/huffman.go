package main

import (
	"bufio"
	"fmt"
	"os"
	"shared/heap"
	"strconv"
	"strings"
)

func runGreedyHuffman() {
	minheap := buildHeap("course3/week3_huffman_and_dp/huffman_codewords_weights.txt")
	maxLen, minLen := countLen(minheap)
	fmt.Println(maxLen, minLen)
}

func buildHeap(filename string) *heap.DistMinHeap {
	weights := make(map[int]int)
	codwordSize := 0
	codeword := 0

	ptr, _ := os.Open(filename)
	defer ptr.Close()

	scanner := bufio.NewScanner(ptr)
	for scanner.Scan() {
		line := scanner.Text()
		codwordSize, _ = strconv.Atoi(strings.TrimSpace(line))
		break
	}

	for scanner.Scan() {
		line := scanner.Text()
		w, _ := strconv.Atoi(strings.TrimSpace(line))
		weights[codeword] = w
		codeword++
	}

	minheap := heap.DistMinHeap{
		Data: make([]int, 0),
		Dist: weights,
		Loc:  make(map[int]int),
	}

	for i := 0; i < codwordSize; i++ {
		minheap.Push(i)
	}

	return &minheap
}

func countLen(h *heap.DistMinHeap) (int, int) {
	var maxMergeCounts = make(map[int]int)
	var minMergeCounts = make(map[int]int)

	for len(h.Data) > 1 {
		// pop two codewords w/ smallest weight then put back one of them with weight updated
		c1, _ := h.Pop()
		c2, _ := h.Pop()
		newWeight := h.Dist[c1] + h.Dist[c2]
		h.Dist[c1] = newWeight
		h.Push(c1)
		// update merge counts for one of them, remove the other
		maxMergeCounts[c1] = 1 + max(maxMergeCounts[c1], maxMergeCounts[c2])
		minMergeCounts[c1] = 1 + min(minMergeCounts[c1], minMergeCounts[c2])
		delete(maxMergeCounts, c2)
		delete(minMergeCounts, c2)
	}

	var maxLen = -1
	for _, v := range maxMergeCounts {
		maxLen = max(maxLen, v)
	}

	var minLen = 1_000_000
	for _, v := range minMergeCounts {
		minLen = min(minLen, v)
	}

	return maxLen, minLen
}
