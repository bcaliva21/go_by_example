package main

import (
	"fmt"
)

func Insertion (a []int) []int {
	for i := 0; i < len(a); i++ {
		m := a[i]
		mi := i
		for j := i; j < len(a); j++ {
			if a[j] < m {
				m = a[j]
			    mi = j
			}
		}
		a[i], a[mi] = a[mi], a[i]
	}

	return a
}

func main() {
	arr := []int{10,8,78,64,20,0}

	fmt.Println(arr)
	
	fmt.Println(Insertion(arr))

}
