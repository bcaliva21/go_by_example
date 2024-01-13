package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	ch <- "from server1"
}

func server2(ch chan string) {
	ch <- "from server2"
}

func main() {
	out1 := make(chan string)
	out2 := make(chan string)
	go server1(out1)
	go server2(out2)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <- out1:
		fmt.Println(s1)
	case s2 := <- out2:
		fmt.Println(s2)
	}
	fmt.Println("--main terminated--")
}

