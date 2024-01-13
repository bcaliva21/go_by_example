package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "joe brown"
	ch <- "sally smith"
	ch <-"steve(third wheel)gonzales"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
