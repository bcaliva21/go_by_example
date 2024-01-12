package main

import (
	"fmt"
	"time"
)

func calcSquares(n int, sqop chan int) {
	s := 0
	for n != 0 {
		d := n % 10
		s += d * d
		n /= 10
	}
	time.Sleep(4 * time.Second)
	sqop <- s
}

func calcCubes(n int, cuop chan int) {
	s := 0
	for n != 0 {
		d := n % 10
		s += d * d * d
		n /= 10
	}
	cuop <- s
}

func main() {
	num := 589
	sqch := make(chan int)
	cuch := make(chan int)
	go calcSquares(num, sqch)
	go calcCubes(num, cuch)
	squares, cubes := <-sqch,<-cuch
	fmt.Println("Final output", squares + cubes)
}
