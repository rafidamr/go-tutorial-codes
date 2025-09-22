package main

import (
	"fmt"
)

// MinHeap is a simple array-backed binary min-heap for ints.
type MinHeap struct {
	data []int
}

// NewMinHeap creates an empty MinHeap.
func NewMinHeap() *MinHeap {
	return &MinHeap{data: make([]int, 0)}
}

// FromSlice builds a MinHeap from the provided slice in O(n).
func FromSlice(values []int) *MinHeap {
	h := &MinHeap{data: append([]int(nil), values...)}
	// Heapify down from the last parent to the root
	for i := parentIndex(len(h.data) - 1); i >= 0; i-- {
		h.heapifyDown(i)
	}
	return h
}

// Len returns the number of elements in the heap.
func (h *MinHeap) Len() int {
	return len(h.data)
}

// Peek returns the minimum element without removing it.
// The second return value is false when the heap is empty.
func (h *MinHeap) Peek() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	return h.data[0], true
}

// Push inserts a value into the heap.
func (h *MinHeap) Push(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

// Pop removes and returns the minimum element.
// The second return value is false when the heap is empty.
func (h *MinHeap) Pop() (int, bool) {
	n := len(h.data)
	if n == 0 {
		return 0, false
	}
	minVal := h.data[0]
	h.swap(0, n-1)
	h.data = h.data[:n-1]
	if len(h.data) > 0 {
		h.heapifyDown(0)
	}
	return minVal, true
}

// Internal helpers

func leftChildIndex(i int) int  { return 2*i + 1 }
func rightChildIndex(i int) int { return 2*i + 2 }
func parentIndex(i int) int     { return (i - 1) / 2 }

func (h *MinHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinHeap) heapifyUp(i int) {
	for i > 0 {
		p := parentIndex(i)
		if h.data[i] < h.data[p] {
			h.swap(i, p)
			i = p
			continue
		}
		break
	}
}

func (h *MinHeap) heapifyDown(i int) {
	n := len(h.data)
	for {
		l := leftChildIndex(i)
		r := rightChildIndex(i)
		smallest := i
		if l < n && h.data[l] < h.data[smallest] {
			smallest = l
		}
		if r < n && h.data[r] < h.data[smallest] {
			smallest = r
		}
		if smallest == i {
			break
		}
		h.swap(i, smallest)
		i = smallest
	}
}

func main() {
	// Example usage
	values := []int{7, 3, 10, 1, 8, 2, 5}
	h := FromSlice(values)
	fmt.Println("heap size:", h.Len())
	if v, ok := h.Peek(); ok {
		fmt.Println("min:", v)
	}
	for h.Len() > 0 {
		v, _ := h.Pop()
		fmt.Print(v, " ")
	}
	fmt.Println()
}
