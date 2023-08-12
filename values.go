package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println([2]bool{true,false})
	fmt.Println([]int{1,2,3})
	fmt.Println(map[string]string{"hello": "world"})
}

