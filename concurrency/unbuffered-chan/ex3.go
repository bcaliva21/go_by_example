package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func main() {
	done := make(chan bool)
	go hello(done)
	<-done // this line is blocking -- read below
	fmt.Println("main function")
}


// -- below --
// until some goroutine writes data to the done channel
// we will not move on

// <-done means we are waiting for done to receive data
// this is called a receiving operation
