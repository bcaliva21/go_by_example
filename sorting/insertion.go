package main

import "fmt"

func Insertion (a []int) []int {

	for i := 1; i < len(a); i++ {
		val, j := a[i],i

		for j > 0 && a[j-1] > val {
			a[j] = a[j -1]
			j--
		}
		a[j] = val
	}

	return a
}

func main () {
	test := []int{10,8,6,4,2,0}

	fmt.Println(test)
	fmt.Println(Insertion(test))

}

