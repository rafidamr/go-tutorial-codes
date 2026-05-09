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

type Graph struct {
	edges        []Edge
	clusters     map[int][]int // lead_id to its members
	leaders      map[int]int   // member_id to its lead_id
	clusters_num int
}

var K = 3

func main() {
	graph := buildGraph("course3/week2_clustering_kruskal_mst/data_500.txt")
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
		v1_lead := g.leaders[e.v1]
		v2_lead := g.leaders[e.v2]

		if v1_lead == v2_lead {
			continue
		}

		var master, absorbed int
		if len(g.clusters[v1_lead]) > len(g.clusters[v2_lead]) {
			master = v1_lead
			absorbed = v2_lead
		} else {
			master = v2_lead
			absorbed = v1_lead
		}

		for _, v := range g.clusters[absorbed] {
			g.leaders[v] = master
			g.clusters[master] = append(g.clusters[master], v)
		}
		delete(g.clusters, absorbed)
		g.clusters_num -= 1
		if g.clusters_num == K {
			spacing = e.cost
			break
		}
	}

	for k := range g.clusters {
		fmt.Println(len(g.clusters[k]))
	}
	fmt.Println("----")

	return spacing
}
