package main

import (
	"fmt"
	"math"
	"strconv"
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

func getDigitCount(n int) int {
	s := strconv.Itoa(n)
	return len(s)
}

func radixSort(a []int) []int {
	m := int(math.Inf(-1))

	for _,v := range a {
		if v > m {
			m = v
		}
	}
	md := getDigitCount(m)
	
	for i := 0; i < md; i++ {
	    da := make([]int,len(a))
		for j := range a {
			da[j] = getDigitAtPlace(a[j],i)
		}
		da = countSort(da)

		// prefix arr, do we need to do this?
		for i := 1; i < len(da); i++ {
			da[i] = da[i-1]+da[i]
		}

	}

	return []int{0}
}

// takes in a number,n and a place, p
// we can build an array of single digits
// based on the place we are planning to sort

func getDigitAtPlace(n, p int) int {
	s := strconv.Itoa(n)
	if p > len(s)-1 {
		return 0
	}
	o := []int{}
	for i := len(s)-1; i > -1; i-- {
		num,_ := strconv.Atoi(string(s[i]))
		o = append(o, num)
	}

	return o[p]
}

func main() {
	a := []int{53,89,150,36,633,233}
	
	fmt.Println(a)
	fmt.Println(radixSort(a))
}

