package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	id   int
	dist int // a distance from a tail to this vertex
}

type Graph map[int][]Vertex

var graph Graph
var heap DijkstraHeap

var conquered = make(map[int]bool)
var start, maxIdx = 1, 200

func main() {
	graph = buildGraph("course2/week2_dijkstra_shortest_path/weighted_graph.txt")
	heap = initHeap()
	expandTerritory()
	inspectKeys()
}

func buildGraph(filename string) Graph {
	g := make(Graph)
	ptr, _ := os.Open(filename)
	defer ptr.Close()

	scanner := bufio.NewScanner(ptr)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "\t")
		tail, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
		for _, strTuple := range arr[1:] {
			if strings.TrimSpace(strTuple) == "" {
				continue
			}
			tuple := strings.Split(strTuple, ",")
			head, _ := strconv.Atoi(strings.TrimSpace(tuple[0]))
			dist, _ := strconv.Atoi(strings.TrimSpace(tuple[1]))
			g[tail] = append(g[tail], Vertex{head, dist})
		}
	}
	return g
}

func initHeap() DijkstraHeap {
	k := make(map[int]int)
	// init all to 1 million distance
	for i := start; i <= maxIdx; i++ {
		k[i] = 1000000
	}
	// actual distance from start vertex to its unconquered neighbors
	k[start] = 0
	for _, v := range graph[start] {
		k[v.id] = v.dist
	}

	h := DijkstraHeap{data: make([]int, 0), totalDist: k, loc: make(map[int]int)}
	for i := start; i <= maxIdx; i++ {
		h.Push(i)
	}
	return h
}

func expandTerritory() {
	for heap.Len() > 0 {
		vId, _ := heap.Pop()
		conquered[vId] = true
		for _, w := range graph[vId] {
			if !conquered[w.id] {
				heap.Remove(w.id)
				greedyScore := min(heap.totalDist[w.id], heap.totalDist[vId]+w.dist)
				heap.totalDist[w.id] = greedyScore
				heap.Push(w.id)
			}
		}
	}
}

func inspectKeys() {
	for _, vId := range []int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197} {
		fmt.Print(heap.totalDist[vId])
		fmt.Print(",")
	}
	fmt.Println()
}
