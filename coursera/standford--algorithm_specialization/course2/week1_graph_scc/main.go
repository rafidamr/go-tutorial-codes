package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Read and parse the file. Store them in reverse graph dict. Keep max and min id.
// Init visited, finish_time, leader (dicts), and max_time
// first pass: idx from max_id to min_id
// second pass: time from max_time to 1. count the vertices per unique leader

type Graph map[int][]int

var visited map[int]bool
var finishTime = make(map[int]int)
var maxVtx, minVtx = -1, math.MaxInt
var time, maxTime = 0, 0
var top5 = make([]int, 5)

func main() {
	graph, reverseGraph := buildGraph("./course2/week1_graph_scc/graph.txt")

	firstPass(reverseGraph)
	secondPass(graph)

	for _, i := range top5 {
		fmt.Printf("%v,", i)
	}
}

// Its objective is to track a correct finishing time for each vertex
func firstPass(reverseGraph Graph) {
	visited = make(map[int]bool)
	for i := maxVtx; i >= minVtx; i-- { // making sure no vertex is skipped
		if _, ok := reverseGraph[i]; ok {
			if !visited[i] {
				firstDfs(reverseGraph, i)
			}
		}
	}
}

func firstDfs(reverseGraph Graph, i int) {
	// Duplicate elements: first element to explore neighbors, second to record finishing time
	stack := []int{i, i}
	popCount := make(map[int]int)
	for len(stack) > 0 {
		lastEl := stack[len(stack)-1]
		stack = slices.Delete(stack, len(stack)-1, len(stack))
		popCount[lastEl]++
		// Record its finishing time once all of its children time are recorded
		if popCount[lastEl] == 2 {
			time++
			finishTime[time] = lastEl
			maxTime = max(maxTime, time)
		}
		visited[lastEl] = true
		for _, nbr := range reverseGraph[lastEl] {
			if !visited[nbr] {
				stack = append(stack, nbr)
				stack = append(stack, nbr)
			}
		}
	}
}

// Its objective is to count the member of SCCs
func secondPass(graph Graph) {
	visited = make(map[int]bool)
	// one iteration is for a cluster of SCCs because it "peels off" SCCs one by one,
	// starting from the sink SCCs
	for t := maxTime; t >= 1; t-- {
		i := finishTime[t]
		if !visited[i] {
			// vertex-i is the initiator for a cluster of SCCs, aka. the sink vertex
			memberCount := secondDfs(graph, i)
			addToTop5(memberCount)
		}
	}
}

func secondDfs(graph Graph, i int) int {
	stack := []int{i}
	visited[i] = true
	memberCount := 1
	for len(stack) > 0 {
		lastEl := stack[len(stack)-1]
		stack = slices.Delete(stack, len(stack)-1, len(stack))
		for _, neighbor := range graph[lastEl] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
				visited[neighbor] = true
				memberCount++
			}
		}
	}
	return memberCount
}

// Alternative.
// This one may push the same element more than once.“
// But will be useful if the use case is tracking the finishing time like in the DFS of the first pass
func _secondDfs(graph Graph, i int) int {
	stack := []int{i}
	memberCount := 0
	for len(stack) > 0 {
		lastEl := stack[len(stack)-1]
		stack = slices.Delete(stack, len(stack)-1, len(stack))
		if !visited[lastEl] {
			memberCount++
		}
		visited[lastEl] = true
		for _, neighbor := range graph[lastEl] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}
	return memberCount
}

func addToTop5(count int) {
	slices.SortFunc(top5, func(a, b int) int {
		return cmp.Compare(b, a)
	})
	if top5[4] < count {
		top5[4] = count
	}
}

func buildGraph(filename string) (Graph, Graph) {
	graph := make(Graph)
	reverseGraph := make(Graph)
	ptr, _ := os.Open(filename)
	defer ptr.Close()

	scanner := bufio.NewScanner(ptr)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		tail, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
		head, _ := strconv.Atoi(strings.TrimSpace(arr[1]))
		if tail != head {
			graph[tail] = append(graph[tail], head)
			reverseGraph[head] = append(reverseGraph[head], tail)
			maxVtx = max(maxVtx, head, tail)
			minVtx = min(minVtx, head, tail)
		}
	}
	return graph, reverseGraph
}
