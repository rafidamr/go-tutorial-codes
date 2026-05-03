package heap

import "slices"

// DistMinHeap stores vertex id but Pop and Push based on the min vertex's distance in the Dist map.
type DistMinHeap struct {
	Data []int
	Dist map[int]int
	Loc  map[int]int
}

func (h *DistMinHeap) Len() int {
	return len(h.Data)
}

func (h *DistMinHeap) Swap(idx1 int, idx2 int) {
	h.Loc[h.Data[idx1]] = idx2
	h.Loc[h.Data[idx2]] = idx1
	h.Data[idx1], h.Data[idx2] = h.Data[idx2], h.Data[idx1]
}

func (h *DistMinHeap) Peek() (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	return h.Data[0], true
}

func (h *DistMinHeap) Push(vId int) {
	h.Data = append(h.Data, vId)
	h.Loc[vId] = h.Len() - 1
	h.HeapifyUp(h.Len() - 1)
}

func (h *DistMinHeap) Pop() (int, bool) {
	smallest, ok := h.Peek()
	if !ok {
		return 0, false
	}
	h.Swap(0, h.Len()-1)
	h.Data = slices.Delete(h.Data, h.Len()-1, h.Len())
	h.HeapifyDown(0)
	return smallest, true
}

func (h *DistMinHeap) Remove(vId int) {
	idx := h.Loc[vId]
	h.Swap(idx, h.Len()-1)
	h.Data = slices.Delete(h.Data, h.Len()-1, h.Len())
	h.HeapifyDown(idx)
	h.HeapifyUp(idx)
}

func (h *DistMinHeap) HeapifyUp(idx int) {
	for idx > 0 && idx < h.Len() {
		p := parentIdxOf(idx)
		if h.Dist[h.Data[idx]] < h.Dist[h.Data[p]] {
			h.Swap(idx, p)
			idx = p
			continue
		}
		break
	}
}

func (h *DistMinHeap) HeapifyDown(idx int) {
	n := h.Len()
	for {
		l := leftChildIdxOf(idx)
		r := rightChildIdxOf(idx)
		iWithSmallerV := idx

		if l < n && h.Dist[h.Data[l]] < h.Dist[h.Data[iWithSmallerV]] {
			iWithSmallerV = l
		}
		if r < n && h.Dist[h.Data[r]] < h.Dist[h.Data[iWithSmallerV]] {
			iWithSmallerV = r
		}
		if iWithSmallerV == idx {
			break
		}

		h.Swap(iWithSmallerV, idx)
		idx = iWithSmallerV
	}
}
