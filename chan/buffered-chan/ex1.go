package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "bradley"
	ch <- "erica"
	fmt.Println(<- ch)
	fmt.Println(<- ch)
}
