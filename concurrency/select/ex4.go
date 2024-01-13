package main

import "fmt"

func main() {
	ch := make(chan int)
	select {
	case <-ch:
	default:
		fmt.Println("default case executed")
	}
	fmt.Println("--main terminated")
}
