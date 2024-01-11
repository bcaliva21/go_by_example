package main

import "fmt"

type Quick struct {
	List []int
}

func (q *Quick) Partition(lo, hi int) int {
	p := q.List[hi]

	i := lo - 1

	for j := lo; j < hi; j++ {
		if q.List[j] <= p {
			i++
			q.List[i], q.List[j] = q.List[j], q.List[i]
		}
	}
	i++
	q.List[i], q.List[hi] = q.List[hi], q.List[i]
	return i
}

func (q *Quick) Sort(lo, hi int) {
	if lo >= hi || lo < 0 {
		return
	}

	p := q.Partition(lo, hi)

	q.Sort(lo, p-1) 
	q.Sort(p+1, hi)
}

func main() {
	q := &Quick{ List: []int{10,8,6,4,2,0} }

	fmt.Println(q)
	q.Sort(0, len(q.List)-1)
	fmt.Println(q)
}

