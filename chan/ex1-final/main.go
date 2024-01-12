package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true // sending operation
}

func main() {
	done := make(chan bool)
	fmt.Println("Main going to call hello goroutine")
	go hello(done)
	<-done // receiving operation
	fmt.Println("Main received data")
}

