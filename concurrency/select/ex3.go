package main

import "fmt"

func main() {
	ch := make(chan string)
	select {
	case <-ch:
	}
	fmt.Println("--main terminated--")
}
