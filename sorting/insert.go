package main

import "fmt"

func Insert (a []interface{}) []interface{} {

	for i := 1; i < len(a); i++ {
		j := i-1
		k := i
		for j > -1 && float32(a[k]) < float32(a[j]) {
			a[k], a[j] = a[j], a[k]
			j--
			k--
			fmt.Println(a)
		}
	}

	return a

}

func main() {
	a := []interface{}{10.10,8,5,10,19,45,29,6,4000000000000}


	fmt.Println(a)
	fmt.Println(Insert(a))
}
