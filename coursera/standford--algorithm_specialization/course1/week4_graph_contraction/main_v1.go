//go:build ignore

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
	fmt.Println("Running V1")
	graph, _cluster := readFromFile("./graph.txt")
	var manyK []int

	for i := 0; i < 1000; i++ {
		cluster := make(Cluster, len(_cluster))
		copy(cluster, _cluster)
		for len(cluster) > 2 {
			contract(&cluster)
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
//   - rand choose cluster A and B from registry
//   - copy all B's vertices into A's list.
//   - Delete B from registry
//   - stop if len(C) == 2, else repeat
func contract(cluster *Cluster) {
	clrLen := len(*cluster)
	iA, iB := 0, 0
	for iA == iB {
		iA = rand.Intn(clrLen)
		iB = rand.Intn(clrLen)
	}
	(*cluster)[iA] = append((*cluster)[iA], (*cluster)[iB]...)
	// replace B with last element then trim the cluster
	(*cluster)[iB] = (*cluster)[clrLen-1]
	*cluster = slices.Delete(*cluster, clrLen-1, clrLen)
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
