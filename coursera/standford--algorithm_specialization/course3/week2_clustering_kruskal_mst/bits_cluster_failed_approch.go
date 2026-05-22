//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nodesSize int
var bitsSize int

func runBitsCluster(filename string) {
	nodes, clusterSize := readNodes(filename)
	eq0 := countDuplicates(&nodes)
	lt3 := countClustersWithHammingDistOneAndTwo(&nodes)
	spacing := clusterSize - eq0 - lt3 + 1 // add by one because it is "kruskal spacing"
	fmt.Println(spacing)
}

func readNodes(filename string) (map[uint32]uint, uint) {
	nodes := make(map[uint32]uint)

	ptr, _ := os.Open(filename)
	defer ptr.Close()

	scanner := bufio.NewScanner(ptr)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		nodesSize, _ = strconv.Atoi(strings.TrimSpace(arr[0]))
		bitsSize, _ = strconv.Atoi(strings.TrimSpace(arr[1]))
		break
	}

	for scanner.Scan() {
		line := scanner.Text()
		bits, _ := strconv.ParseUint(strings.ReplaceAll(line, " ", ""), 2, bitsSize)
		nodes[uint32(bits)]++
	}

	scanner.Err()

	return nodes, uint(nodesSize)
}

func countDuplicates(nodes *map[uint32]uint) uint {
	eq0 := uint(0)
	seen := make(map[uint32]bool)

	for n, _ := range *nodes {
		if !seen[n] && (*nodes)[n] > 1 {
			eq0 += (*nodes)[n]
		}
		seen[n] = true
	}
	return eq0
}

func countClustersWithHammingDistOneAndTwo(nodes *map[uint32]uint) uint {
	var lt3 uint = 0
	seen := make(map[uint32]bool)

	masks := getMasks()
	for n, _ := range *nodes {
		for _, m := range masks {
			if (*nodes)[n^m] > 0 && !seen[n^m] {
				lt3++
			}
		}
		seen[n] = true
	}
	return lt3
}

// Return all possible mask m with popcount(m) == d (only handles d = 1 or 2).
// A mask flips bits so that another bits at a hamming distance d can be found
//
// Example at d = 1:
// m = 010
// a = 111
// b = 101
//
// a == b ^ m OR b == a ^ m
func getMasks() []uint32 {
	var masks []uint32

	// d = 0
	for i := uint32(0); i < uint32(bitsSize); i++ {
		masks = append(masks, 1<<i)
	}

	// d = 1
	for i := uint32(0); i < uint32(bitsSize); i++ {
		for j := i + 1; j < uint32(bitsSize); j++ {
			masks = append(masks, (1<<i)|(1<<j))
		}
	}

	return masks
}
