package main

import "fmt"

func merge (a, b []int) []int {
	o := []int{}
	i, j := 0,0
	for i <= len(a)-1 || j <= len(b)-1 {
		if i == len(a) {
			o = append(o,b[j:]...)
			break
		}
		if j == len(b) {
			o = append(o,a[i:]...)
			break
		}
		if a[i] <= b[j] {
			o = append(o,a[i])
			i++
		} else {
			o = append(o,b[j])
			j++
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

