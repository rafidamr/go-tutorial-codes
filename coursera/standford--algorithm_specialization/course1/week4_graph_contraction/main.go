package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Graph map[int][]int
type Cluster [][]int

func main() {
	graph, _cluster := readFromFile("./graph.txt")
	var manyK []int

	for i := 0; i < 1000; i++ {
		cluster := make(Cluster, len(_cluster))
		copy(cluster, _cluster)
		for len(cluster) > 2 {
			contract(&graph, &cluster)
		}
		k := countCrossingEdges(&graph, &cluster)
		manyK = append(manyK, k)
	}

	top := slices.Min(manyK)
	fmt.Println(top)
}

// read the file then build the graph G and cluster registry C
func readFromFile(filename string) (Graph, Cluster) {
	var graph = make(Graph)
	var cluster Cluster

	fPointer, _ := os.Open(filename)
	defer fPointer.Close()

	scanner := bufio.NewScanner(fPointer)
	for scanner.Scan() {
		var vertexId int
		var iArr []int

		line := scanner.Text()
		numSeq := strings.Split(line, "\t")

		for i, tNum := range numSeq {
			tNum = strings.TrimSpace(tNum)
			if tNum == "" {
				continue
			}
			iNum, _ := strconv.Atoi(tNum)
			if i == 0 {
				vertexId = iNum
			} else {
				iArr = append(iArr, iNum)
			}
		}

		graph[vertexId] = iArr
		cluster = append(cluster, []int{vertexId})
	}

	return graph, cluster
}

// contract
//   - pick a random edge that crosses two different clusters
//   - merge the two clusters
//   - stop if len(C) == 2, else repeat
func contract(graph *Graph, cluster *Cluster) {
	// build membership map: vertex -> cluster index
	membership := make(map[int]int)
	for i, c := range *cluster {
		for _, v := range c {
			membership[v] = i
		}
	}

	// collect all cross-cluster edges (each edge counted once via u < v)
	type Edge struct{ u, v int }
	var edges []Edge
	for u, neighbors := range *graph {
		for _, v := range neighbors {
			if u < v && membership[u] != membership[v] {
				edges = append(edges, Edge{u, v})
			}
		}
	}

	if len(edges) == 0 {
		return
	}

	// pick a random crossing edge, then merge its two clusters
	e := edges[rand.Intn(len(edges))]
	iA := membership[e.u]
	iB := membership[e.v]

	(*cluster)[iA] = append((*cluster)[iA], (*cluster)[iB]...)
	clrLen := len(*cluster)
	(*cluster)[iB] = (*cluster)[clrLen-1]
	*cluster = (*cluster)[:clrLen-1]
}

// count crossing k edges
//   - choose the first of two clusters of C
//   - for each vertex i of the cluster, retrieve all its neighboring vertices from G
//   - count k-i: how many i's neighbors not in the same cluster
//   - sum all k-i then save in result
func countCrossingEdges(graph *Graph, cluster *Cluster) int {
	var k int
	clusterA := (*cluster)[0]
	table := make(map[int]bool)

	for _, vertex := range clusterA {
		table[vertex] = true
	}

	for _, vertex := range clusterA {
		neighbors := (*graph)[vertex]
		for _, neighbor := range neighbors {
			if !table[neighbor] {
				k += 1
			}
		}
	}

	return k
}
