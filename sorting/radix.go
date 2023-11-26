package main

import (
	"fmt"
	"math"
	"strconv"
)

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

func getDigitCount(n int) int {
	s := strconv.Itoa(n)
	return len(s)
}

// this implementation uses counting sort
func radixSort(input []int) []int {
	maxVal := int(math.Inf(-1))

	for _,v := range input {
		if v > maxVal {
			maxVal = v
		}
	}
	numOfMaxValDigits := getDigitCount(maxVal)
	
	for i := 0; i < numOfMaxValDigits; i++ {
	    digits := make([]int,len(input))
		for j := range input {
			digits[j] = getDigitAtPlace(input[j],i)
		}
		occurences := [10]int{}

		for _,v := range digits {
			occurences[v]++
		}

		for i := 1; i < len(occurences); i++ {
			occurences[i] += occurences[i-1]
		}

		orderedDigits := make([]int,len(input))
		reconstructedInput := make([]int,len(input))

		for i := len(digits)-1; i > -1; i-- {
			val := digits[i]
			occurences[val] -= 1
			idx := occurences[val]
			orderedDigits[idx] = val
			reconstructedInput[idx] = input[i]
		}
		input = reconstructedInput
	}

	return input
}

func main() {
	a := []int{53,89,150,36,633,233,9000,18,17,0,887}
	
	fmt.Println(a)
	fmt.Println(radixSort(a))
}

