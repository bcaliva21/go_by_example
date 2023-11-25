package main

import (
	"fmt"
	"math"
)

func countSort(a []int) []int {
	m := int(math.Inf(-1))

	// find max
	for _,v := range a {
		if v > m {
			m = v
		}
	}

	c := make([]int,m+1)

	// count occurences
	for _,v := range a {
		c[v]++
	}

	// calc prefix sum
	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	o := make([]int,len(a)-1)

	// populate output array
	for i := len(a)-1; i > -1; i-- {
		num := a[i]
		idx := c[num]-1
		if c[num] > 0 {
			c[num] = c[num]-1
		}
		if idx > 0 {
		    o[idx-1] = num
		}
	}

	return o
}

func main() {
	// adding one very large number slows the algo considerably
	a := []int{10,8,6,4,2,9,7,5,3,1,1000000000,10,6,8,4,3,10,2,0}

	fmt.Println(a)
	fmt.Println(countSort(a))
}

