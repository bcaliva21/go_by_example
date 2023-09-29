package main
import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Extract () int {
	extracted := h.array[0]

	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]

	h.HeapifyDown()
	
	return extracted
}
func (h *MaxHeap) Insert (value int) {
	h.array = append(h.array, value)

	h.HeapifyUp()
}

func (h *MaxHeap) HeapifyUp () {
	lastIndex := len(h.array) - 1
	par := parent(lastIndex)

	for h.array[lastIndex] > h.array[par] {
		h.Swap(lastIndex, par)
		lastIndex = par
		par = parent(lastIndex)
	}
}

func (h *MaxHeap) HeapifyDown () {
	i := 0
	l, r := left(i), right(i)
	childToCompare := 0

	for l <= len(h.array) - 1 {
		if l == len(h.array) - 1 {
			childToCompare = l
		} else if l > r {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.array[childToCompare] > h.array[i] {
			h.Swap(childToCompare, i)
			i = childToCompare
			l, r = left(i), right(i)
		} else {
			return
		}
	}
}

func (h *MaxHeap) Swap (i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func left (i int) int {
	return i * 2 + 1
}

func right (i int) int {
	return i * 2 + 2
}

func parent (i int) int {
	return (i - 1) / 2
}

func main() {
	heap := &MaxHeap{}
	buildHeap := []int{10,20,30}

	for _,value := range buildHeap {
		heap.Insert(value)
		fmt.Println(heap)
	}

	for i := 0; i < 3; i++ {
		heap.Extract()
		fmt.Println(heap)
	}
}

