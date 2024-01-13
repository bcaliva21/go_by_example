package main

import (
	"fmt"
	"math/rand"
)

var greetings = [5]string{"hello", "ni hao", "hola", "halo", "lei hou"}

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		rand := rand.Intn(len(greetings))
		ch <- rand
	}
	close(ch)
}

func consumer(nums chan int, strings chan string) {
	for num := range nums {
		strings <- greetings[num]
	}
	close(strings)
}

func printer(strings chan string) {
	for s := range strings {
		fmt.Println(s)
	}
}

func main() {
	randonums := make(chan int)
	randogreetings := make(chan string)
	go producer(randonums)
	go consumer(randonums, randogreetings)
	<-randogreetings
	printer(randogreetings)
	fmt.Println("--main terminated--")
}

