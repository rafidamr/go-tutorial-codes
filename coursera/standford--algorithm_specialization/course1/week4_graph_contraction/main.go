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
type Edge struct {
	A, B int
}

func main() {
	graph, _cluster := buildGraphClusterFromFile("./graph.txt")
	var manyK []int

	for i := 0; i < 100; i++ {
		clusters := make(Cluster, len(_cluster))
		copy(clusters, _cluster)
		for len(clusters) > 2 {
			contract(&graph, &clusters)
		}
		k := countCrossingEdges(&graph, &clusters)
		manyK = append(manyK, k)
	}

	top := slices.Min(manyK)
	fmt.Println(top)
}

// read the file then build the graph G and cluster registry C
func buildGraphClusterFromFile(filename string) (Graph, Cluster) {
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
//   - build edges between clusters
//   - randomly choose one edge (connects A and B clusters)
//   - copy all B's vertices into A's list.
//   - Delete B from registry
//   - stop if len(C) == 2, else repeat
func contract(graph *Graph, clusters *Cluster) {
	membership := make(map[int]int)
	// tracks a vertex membership to a cluster, needed for building edges between clusters
	for clusterId, members := range *clusters {
		for _, vertex := range members {
			membership[vertex] = clusterId
		}
	}

	var edges []Edge
	for v1, neighbors := range *graph {
		for _, v2 := range neighbors {
			var A, B = membership[v1], membership[v2]
			if v1 < v2 && A != B {
				edges = append(edges, Edge{A, B})
			}
		}
	}

	edge := edges[rand.Intn(len(edges))]

	// move all members of B to A
	(*clusters)[edge.A] = append((*clusters)[edge.A], (*clusters)[edge.B]...)
	// Delete B by moving to the last position in list then trim the list
	clustersNum := len(*clusters)
	(*clusters)[edge.B] = (*clusters)[clustersNum-1]
	(*clusters) = slices.Delete(*clusters, clustersNum-1, clustersNum)
}

// count crossing k edges
//   - choose the first of two clusters of C
//   - for each vertex i of the cluster, retrieve all its neighboring vertices from G
//   - count k-i: how many i's neighbors not in the same cluster
//   - sum all k-i then save in result
func countCrossingEdges(graph *Graph, cluster *Cluster) int {
	var k int
	A := (*cluster)[0]
	memberOfA := make(map[int]bool)

	for _, vertex := range A {
		memberOfA[vertex] = true
	}

	for _, vertex := range A {
		neighbors := (*graph)[vertex]
		for _, neighbor := range neighbors {
			if !memberOfA[neighbor] {
				k += 1
			}
		}
	}

	return k
}

// It fails because it merges clusters without considering their connectivity
func _contract(clusters *Cluster) {
	clrLen := len(*clusters)
	iA, iB := 0, 0
	for iA == iB {
		iA = rand.Intn(clrLen)
		iB = rand.Intn(clrLen)
	}
	(*clusters)[iA] = append((*clusters)[iA], (*clusters)[iB]...)
	// replace B with last element then trim the cluster
	(*clusters)[iB] = (*clusters)[clrLen-1]
	*clusters = slices.Delete(*clusters, clrLen-1, clrLen)
}
