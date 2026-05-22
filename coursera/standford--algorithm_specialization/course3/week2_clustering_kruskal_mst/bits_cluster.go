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
	nodes := readNodes(filename)
	g := buildGraphFromNodes(&nodes)
	fmt.Println(g.clustersNum)
}

func readNodes(filename string) []int {
	var nodes []int

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
		nodes = append(nodes, int(bits))
	}

	scanner.Err()

	return nodes
}

// finds the k cluster number such that its minimum edge cost is 3 / spacing at least 3
// aka. when the minimum edge cost is 3 then it's merged, the number of cluster become k-1
func buildGraphFromNodes(nodes *[]int) UnionFindGraph {
	var g UnionFindGraph
	g.clusters = make(map[int][]int)
	g.leaders = make(map[int]int)

	for _, n := range *nodes {
		if lead_id := g.Find(n); lead_id == -1 { // avoid duplicates (d=0)
			g.leaders[n] = n
			g.clusters[n] = append(g.clusters[n], n)
			g.clustersNum++
		}
	}

	for _, n := range *nodes {
		for _, m := range getMasks() { // bits masks (d=1 and d=2)
			if lead_id := g.Find(n ^ m); lead_id != -1 {
				g.Union(lead_id, n)
			}
		}
	}

	return g
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
func getMasks() []int {
	var masks []int

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
