package main

import (
	"fmt"
	"slices"
)

// Minimum Heap
type Heap struct {
	data []int
}

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) Swap(idx1 int, idx2 int) {
	h.data[idx1], h.data[idx2] = h.data[idx2], h.data[idx1]
}

func (h *Heap) Peek() (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	return h.data[0], true
}

func (h *Heap) Push(val int) {
	h.data = append(h.data, val)
	h.HeapifyUp(h.Len() - 1)
}

func (h *Heap) Pop() (int, bool) {
	smallest, ok := h.Peek()
	if !ok {
		return 0, false
	}
	h.Swap(0, h.Len()-1)
	h.data = slices.Delete(h.data, h.Len()-1, h.Len())
	h.HeapifyDown(0)
	return smallest, true
}

func (h *Heap) HeapifyUp(idx int) {
	for idx > 0 {
		p := parentIdxOf(idx)
		if h.data[idx] < h.data[p] {
			h.Swap(idx, p)
			idx = p
			continue
		}
		break
	}
}

func (h *Heap) HeapifyDown(idx int) {
	n := h.Len()
	for {
		l := leftChildIdxOf(idx)
		r := rightChildIdxOf(idx)
		iWithSmallerV := idx

		if l < n && h.data[l] < h.data[iWithSmallerV] {
			iWithSmallerV = l
		}
		if r < n && h.data[r] < h.data[iWithSmallerV] {
			iWithSmallerV = r
		}
		if iWithSmallerV == idx {
			break
		}

		h.Swap(iWithSmallerV, idx)
		idx = iWithSmallerV
	}
}

// Helpers
func parentIdxOf(i int) int     { return (i - 1) / 2 }
func leftChildIdxOf(i int) int  { return 2*i + 1 }
func rightChildIdxOf(i int) int { return 2*i + 2 }

// Utils
func buildDummyHeap() *Heap {
	var h Heap
	arr := []int{8, 7, 1, 5, 9}
	for _, v := range arr {
		h.Push(v)
	}
	return &h
}

func testDummyHeap(h *Heap) {
	fmt.Println((*h).data)
	h.Pop()
	fmt.Println((*h).data)
	h.Pop()
	fmt.Println((*h).data)
	h.Pop()
	fmt.Println((*h).data)
	h.Pop()
	fmt.Println((*h).data)
	h.Pop()
	fmt.Println((*h).data)
	h.Pop()
	fmt.Println((*h).data)
}
