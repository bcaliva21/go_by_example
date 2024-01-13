package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println(<-chnl)
}

// what's going on here?

// we init a bidirectional chan on line 7
// in our function on line 13 we convert the
// channel into a send only channel because of the props
// in teh function signature

// so chnl is bidirectional in main, but inside of sendData
// it is a send only channel
