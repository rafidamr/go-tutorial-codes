package heap

import (
	"fmt"
	"slices"
)

type Mode string

const (
	MinMode Mode = "min"
	MaxMode Mode = "max"
)

type Heap struct {
	data []int
	mode Mode
}

func NewHeap(m Mode) Heap {
	switch m {
	case MinMode, MaxMode:
		return Heap{mode: m}
	}
	panic("Invalid Heap Mode")
}

func (h *Heap) Compare(idx1 int, idx2 int) bool {
	if h.mode == MinMode {
		return h.data[idx1] < h.data[idx2]
	} else {
		return h.data[idx1] > h.data[idx2]
	}
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
	h.heapifyUp(h.Len() - 1)
}

func (h *Heap) Pop() int {
	smallest, ok := h.Peek()
	if !ok {
		panic("Pop Error")
	}
	h.Swap(0, h.Len()-1)
	h.data = slices.Delete(h.data, h.Len()-1, h.Len())
	h.heapifyDown(0)
	return smallest
}

func (h *Heap) heapifyUp(idx int) {
	for idx > 0 {
		p := parentIdxOf(idx)
		if h.Compare(idx, p) {
			h.Swap(idx, p)
			idx = p
			continue
		}
		break
	}
}

func (h *Heap) heapifyDown(idx int) {
	n := h.Len()
	for {
		l := leftChildIdxOf(idx)
		r := rightChildIdxOf(idx)
		iWithSmallerV := idx

		if l < n && h.Compare(l, iWithSmallerV) {
			iWithSmallerV = l
		}
		if r < n && h.Compare(r, iWithSmallerV) {
			iWithSmallerV = r
		}
		if iWithSmallerV == idx {
			break
		}

		h.Swap(iWithSmallerV, idx)
		idx = iWithSmallerV
	}
}

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
