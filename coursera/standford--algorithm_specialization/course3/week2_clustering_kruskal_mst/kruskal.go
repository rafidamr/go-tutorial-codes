package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Edge struct {
	v1   int
	v2   int
	cost int
}

// a graph optimized/acting as Union-Find data structure
type Graph struct {
	edges        []Edge
	clusters     map[int][]int // lead_id to its members
	leaders      map[int]int   // member_id to its lead_id
	clusters_num int
}

// return the leader of v
func (g *Graph) Find(v int) int {
	return g.leaders[v]
}

// unify the clusters of a pair of vertices into one
func (g *Graph) Union(v1 int, v2 int) {
	var master, absorbed int

	if len(g.clusters[g.Find(v1)]) > len(g.clusters[g.Find(v2)]) {
		master = g.Find(v1)
		absorbed = g.Find(v2)
	} else {
		master = g.Find(v2)
		absorbed = g.Find(v1)
	}

	for _, v := range g.clusters[absorbed] {
		g.leaders[v] = master
		g.clusters[master] = append(g.clusters[master], v)
	}

	delete(g.clusters, absorbed)
	g.clusters_num -= 1
}

var K = 3

func runKruskal(filename string) {
	graph := buildGraph(filename)
	graph.edges = sortEdges(&graph.edges)
	spacing := cluster(&graph)
	fmt.Println(spacing)
}

func buildGraph(filename string) Graph {
	var g Graph
	g.clusters = make(map[int][]int)
	g.leaders = make(map[int]int)

	ptr, _ := os.Open(filename)
	defer ptr.Close()

	scanner := bufio.NewScanner(ptr)
	for scanner.Scan() {
		line := scanner.Text()
		g.clusters_num, _ = strconv.Atoi(strings.TrimSpace(line))
		break
	}

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		v1, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
		v2, _ := strconv.Atoi(strings.TrimSpace(arr[1]))
		cost, _ := strconv.Atoi(strings.TrimSpace(arr[2]))

		e := Edge{v1, v2, cost}
		g.edges = append(g.edges, e)

		if _, ok := g.clusters[v1]; !ok {
			g.leaders[v1] = v1
			g.clusters[v1] = append(g.clusters[v1], v1)
		}
		if _, ok := g.clusters[v2]; !ok {
			g.leaders[v2] = v2
			g.clusters[v2] = append(g.clusters[v2], v2)
		}
	}

	return g
}

// sort edge in ascending order by its cost
func sortEdges(edges *[]Edge) []Edge {
	sorted := slices.Clone(*edges)
	slices.SortFunc(sorted, func(e1, e2 Edge) int {
		return cmp.Compare(e1.cost, e2.cost)
	})
	return sorted
}

func cluster(g *Graph) int {
	var spacing int

	for _, e := range g.edges {
		if g.Find(e.v1) == g.Find(e.v2) {
			continue
		}

		g.Union(e.v1, e.v2)

		if g.clusters_num == K {
			spacing = e.cost
			break
		}
	}

	return spacing
}
