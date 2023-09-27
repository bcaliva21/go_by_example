package main

import "fmt"

type MaxHeap struct {
	array []int
	size int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array,key)
	h.size = len(h.array)-1
	h.maxHeapifyUp()
}

func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	h.size = h.size-1

	if h.size < 1 {
		fmt.Println("Nothing to extract")
		return -1
	}

	h.array[0] = h.array[h.size]
	h.array = h.array[:h.size]

	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyUp() {
	index := h.size
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}

}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := h.size
	l,r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func parent(i int) int {
	return (i-1)/2
}

func left (i int) int {
	return 2*i+1
}

func right (i int) int {
	return 2*i+2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10,20,30, 5,7,9,11,13,15,17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}

}

