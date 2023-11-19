package main

import "fmt"

func binarySort (a []int, t int) int {
	s := 0
	e := len(a)-1
	m := (s+e)/2

	if a[m] == t {
		return m
	}

	for s < e {
		if a[m] > t {
			e = m-1
			m = (s+e)/2
		} else {
			s = m+1
			m = (s+e)/2
		}
	}

	if a[m] == t {
		return m
	} else {
		return -1
	}
}

func main() {
	a := []int{1,3,5,7,9,11,13,15,17,19,21,23,25,27,29}

	fmt.Println(binarySort(a, 1))
	fmt.Println(binarySort(a, 29))
	fmt.Println(binarySort(a, 15))
	fmt.Println(binarySort(a, 100))
}

