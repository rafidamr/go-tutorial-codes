package main

import (
	"slices"
)

// DijkstraHeap stores vertex id but Pop and Push based on the vertex's distance in the keys map
type DijkstraHeap struct {
	data      []int       // vertex ids
	totalDist map[int]int // distances of vertex ids from the starting vertex
	loc       map[int]int // location of a vertex id in the data slice
}

func (h *DijkstraHeap) Len() int {
	return len(h.data)
}

func (h *DijkstraHeap) Swap(idx1 int, idx2 int) {
	h.loc[h.data[idx1]] = idx2
	h.loc[h.data[idx2]] = idx1
	// swap vertex id must occur after its new location is set
	h.data[idx1], h.data[idx2] = h.data[idx2], h.data[idx1]
}

func (h *DijkstraHeap) Peek() (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	return h.data[0], true
}

func (h *DijkstraHeap) Push(vId int) {
	h.data = append(h.data, vId)
	h.loc[vId] = h.Len() - 1
	h.HeapifyUp(h.Len() - 1)
}

func (h *DijkstraHeap) Pop() (int, bool) {
	smallest, ok := h.Peek()
	if !ok {
		return 0, false
	}
	h.Swap(0, h.Len()-1)
	h.data = slices.Delete(h.data, h.Len()-1, h.Len())
	h.HeapifyDown(0)
	return smallest, true
}

// vId is the vertex id
func (h *DijkstraHeap) Remove(vId int) {
	idx := h.loc[vId]
	h.Swap(idx, h.Len()-1)
	h.data = slices.Delete(h.data, h.Len()-1, h.Len())
	h.HeapifyDown(idx)
}

// idx is the location of vertex in data slice
func (h *DijkstraHeap) HeapifyUp(idx int) {
	for idx > 0 {
		p := parentIdxOf(idx)
		if h.totalDist[h.data[idx]] < h.totalDist[h.data[p]] {
			h.Swap(idx, p)
			idx = p
			continue
		}
		break
	}
}

func (h *DijkstraHeap) HeapifyDown(idx int) {
	n := h.Len()
	for {
		l := leftChildIdxOf(idx)
		r := rightChildIdxOf(idx)
		iWithSmallerV := idx

		if l < n && h.totalDist[h.data[l]] < h.totalDist[h.data[iWithSmallerV]] {
			iWithSmallerV = l
		}
		if r < n && h.totalDist[h.data[r]] < h.totalDist[h.data[iWithSmallerV]] {
			iWithSmallerV = r
		}
		if iWithSmallerV == idx {
			break
		}

		h.Swap(iWithSmallerV, idx)
		idx = iWithSmallerV
	}
}
