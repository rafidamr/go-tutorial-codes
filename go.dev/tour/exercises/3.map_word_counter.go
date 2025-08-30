package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	parts := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range parts {
		m[string(v)] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
