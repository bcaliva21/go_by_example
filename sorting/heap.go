package main

import "fmt"

type MaxHeap struct {
	List []int
}

func (h *MaxHeap) Insert(n int) {
	h.List = append(h.List, n)
	h.HeapifyUp()
}

func (h *MaxHeap) Extract() int {
	l := len(h.List) - 1
	o := h.List[0]

	if l < 1 {
		return -1
	}

	h.List[0] = h.List[l]
	h.List = h.List[:l]

	h.HeapifyDown(len(h.List)-1)

	return o
}

func (h *MaxHeap) HeapifyDown(last int) {
	i := 0
	c := 0
	l := h.GetLeft(i)
	r := h.GetRight(i)

	for l <= last {
		if l == last {
			c = l
		} else if h.List[l] > h.List[r] {
			c = l
		} else {
			c = r
		}

		if h.List[i] < h.List[c] {
			h.Swap(i,c)
			i = c
			l,r = h.GetLeft(i), h.GetRight(i)
		} else {
			return
		}
	}
}

func (h *MaxHeap) HeapifyUp() {
	i := len(h.List) - 1
	p := h.GetParent(i)

	for p > -1 && h.List[i] > h.List[p] {
		h.Swap(i,p)
		i = p
		p = h.GetParent(i)
	}
}

func (h *MaxHeap) GetParent(i int) int {
	return (i-1)/2 
}

func (h *MaxHeap) GetLeft(i int) int {
	return (i*2)+1
}

func (h *MaxHeap) GetRight(i int) int {
	return (i*2)+2
}

func (h *MaxHeap) Swap(i1, i2 int) {
	h.List[i1], h.List[i2] = h.List[i2], h.List[i1]
}

func HeapSort (a []int) []int {
	h := &MaxHeap{}
	for _,v := range a {
		h.Insert(v)
	}

	for i := len(h.List) - 1; i > -1; i-- {
		h.Swap(0,i)
		h.HeapifyDown(i-1)
	}

	return h.List
}

func main() {
	a := []int{10,39,8,6,34,4,21,11,0,-10,94,0,-4,99}
	// heaptest := &MaxHeap{}

	fmt.Println(a)
	fmt.Println(HeapSort(a))
	// fmt.Println(heaptest)
	// heaptest.Insert(10)
	// heaptest.Insert(20)
	// heaptest.Insert(30)
	// heaptest.Insert(40)
	// heaptest.Insert(50)
	// fmt.Println(heaptest)
	// fmt.Println(heaptest.Extract())
	// fmt.Println(heaptest.Extract())
	// fmt.Println(heaptest)
}
