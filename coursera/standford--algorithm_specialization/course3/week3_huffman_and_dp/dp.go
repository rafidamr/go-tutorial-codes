package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func runDP() {
	stream := buildStream("course3/week3_huffman_and_dp/path_graph.txt")
	maxWIS := buildMaxWIS(stream)
	WISmembership := buildWISMembership(maxWIS, stream)
	// fmt.Println(WISmembership)
	for _, t := range []int{1, 2, 3, 4, 17, 117, 517, 997} {
		i := t - 1
		if (*WISmembership)[i] {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}
}

func buildStream(filename string) *[]int {
	var stream []int

	ptr, _ := os.Open(filename)
	defer ptr.Close()

	scanner := bufio.NewScanner(ptr)
	for scanner.Scan() {
		scanner.Text()
		break
	}

	for scanner.Scan() {
		line := scanner.Text()
		w, _ := strconv.Atoi(strings.TrimSpace(line))
		stream = append(stream, w)
	}

	return &stream
}

func buildMaxWIS(stream *[]int) *[]int {
	var s = *stream
	var maxWIS []int // Max Weight Independent Set of Path Graph (tracks the max weight at vertex id)
	var vId = 0

	// first element always has max weight from itself
	maxWIS = append(maxWIS, s[vId])
	vId++

	// handles second element
	maxWIS = append(maxWIS, max(s[vId-1], s[vId]))
	vId++

	// handles 3rd element to the end
	for i := vId; i < len(s); i++ {
		maxWIS = append(maxWIS, max(maxWIS[i-1], s[i]+maxWIS[i-2]))
	}

	return &maxWIS
}

func buildWISMembership(maxWIS *[]int, stream *[]int) *map[int]bool {
	var s = *stream
	var is = *maxWIS
	var membership = make(map[int]bool) // true if vertex id is in WIS
	var i int

	for i = len(is) - 1; i > 1; {
		if is[i-1] >= s[i]+is[i-2] {
			i--
		} else {
			membership[i] = true
			i = i - 2
		}
	}

	// handles base (i == 1 or i == 0)
	membership[i] = true

	return &membership
}
