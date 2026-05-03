package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	dheap "shared/heap"
)

type Vertex struct {
	id   int
	dist int // a distance from a tail to this vertex
}

type Graph map[int][]Vertex

var graph Graph
var heap dheap.DistMinHeap

var conquered = make(map[int]bool)
var vStart, vMaxId int

func RunPrim() {
	graph = buildGraph("course3/week1_greedy_schedule_mst/edges.txt")
	heap = initHeap()
	cost := calculateMstCost()
	fmt.Println(cost)
}

func buildGraph(filename string) Graph {
	g := make(Graph)
	ptr, _ := os.Open(filename)
	defer ptr.Close()

	scanner := bufio.NewScanner(ptr)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		vMaxId, _ = strconv.Atoi(strings.TrimSpace(arr[0]))
		vStart = 1
		break
	}

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		v1, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
		v2, _ := strconv.Atoi(strings.TrimSpace(arr[1]))
		dist, _ := strconv.Atoi(strings.TrimSpace(arr[2]))
		g[v1] = append(g[v1], Vertex{v2, dist})
		g[v2] = append(g[v2], Vertex{v1, dist})
	}
	return g
}

func initHeap() dheap.DistMinHeap {
	dist := make(map[int]int)
	// init all vertices to 1 million distance
	for i := vStart; i <= vMaxId; i++ {
		dist[i] = 1000000
	}
	// set actual distance from start vertex to its unconquered neighbors
	dist[vStart] = 0
	for _, v := range graph[vStart] {
		dist[v.id] = v.dist
	}

	h := dheap.DistMinHeap{
		Data: make([]int, 0),
		Dist: dist,
		Loc:  make(map[int]int),
	}
	for i := vStart; i <= vMaxId; i++ {
		h.Push(i)
	}

	return h
}

func calculateMstCost() int {
	var cost int
	for heap.Len() > 0 {
		vId, _ := heap.Pop()
		conquered[vId] = true
		if vId != vStart {
			for _, w := range graph[vId] {
				if !conquered[w.id] {
					heap.Remove(w.id)
					greedyScore := min(heap.Dist[w.id], w.dist)
					heap.Dist[w.id] = greedyScore
					heap.Push(w.id)
				}
			}
			cost += heap.Dist[vId]
		}
	}
	return cost
}
