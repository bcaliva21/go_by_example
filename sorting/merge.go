package main

import "fmt"

func merge (a, b []int) []int {
	o := []int{}
	for len(a) > 0 || len(b) > 0 {
		if len(a) == 0 {
			for _,v := range b {
				o = append(o,v)
			}
			b = []int{}
			break
		}
		if len(b) == 0 {
			for _,v := range a {
				o = append(o,v)
			}
			a = []int{}
			break
		}
		if a[0] <= b[0] {
			o = append(o,a[0])
			a = a[1:]
		} else {
			o = append(o,b[0])
			b = b[1:]
		}
	}

	return o
}

func mergeSort(a []int, beg, end int) []int {
	if beg == end {
		return a
	}
	mid := (beg+end+1)/2
	lft := mergeSort(a[:mid], beg, mid-1)
	rgt := mergeSort(a[mid:], beg, end-mid)

	return merge(lft,rgt)
}

func main() {
    a := []int{10,8,9,6,7,4,5,9,7,1,11,123,42,0,98,28,33,45,10,1234}

	fmt.Println(a)
	fmt.Println(mergeSort(a, 0, len(a)-1))
}

