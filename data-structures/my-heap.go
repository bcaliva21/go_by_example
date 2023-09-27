package main

import "fmt"

type MaxHeap struct {
	array     []int
	size      int
}

func (h *MaxHeap) Insert (value int) {
	h.array = append(h.array,value)
	h.size = len(h.array)
    h.MaxHeapifyUp(h.size-1)
}

func (h *MaxHeap) Swap (i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func (h *MaxHeap) Extract () int {
	extracted := h.array[0]
	h.size = h.size - 1

	if h.size < 1 {
		fmt.Println("Nothing to extract")
		return -1
	}

	h.array[0] = h.array[h.size]
	h.array = h.array[:h.size]

	h.MaxHeapifyDown()

	return extracted
}

func (h *MaxHeap) MaxHeapifyUp (i int) {
	par := parent(i)
	for h.array[i] > h.array[par] {
		h.Swap(i, par)
		i, par = par, parent(par)
	}
}

func (h *MaxHeap) MaxHeapifyDown () {
	index := 0
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= h.size - 1 {
		if l == h.size - 1 {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.array[childToCompare] > h.array[index] {
			h.Swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

func parent (index int) int {
	return (index - 1) / 2
}

func left (index int) int {
	return (index * 2) + 1
}

func right (index int) int {
	return (index * 2) + 2
}

func main() {
	heap := &MaxHeap{}
	buildHeap := []int{10,20,30,40,50}

	for _,val := range buildHeap {
		heap.Insert(val)
		fmt.Println(heap)
	}

	for i := 0; i < 5; i++ {
		heap.Extract()
		fmt.Println(heap)
	}

}

