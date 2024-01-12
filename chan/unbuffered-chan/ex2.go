package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello goroutine")
}

func main() {
	fmt.Println("main function")
	time.Sleep(1 * time.Second)
	go hello()
}

