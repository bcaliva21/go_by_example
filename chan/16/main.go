package main

import "fmt"

func digits(num int, dchnl chan int) {
	for num != 0 {
		dig := num % 10
		dchnl <- dig
		num /= 10
	}
	close(dchnl)
}

func calcSquares(num int, sqop chan int) {
	s := 0
	dch := make(chan int)
	go digits(num, dch)
	for dig := range dch {
		s += dig * dig
	}
	sqop <- s
}

func calcCubes(num int, cuop chan int) {
	s := 0
	dch := make(chan int)
	go digits(num, dch)
	for dig := range dch {
		s += dig * dig * dig
	}
	cuop <- s
}

func main() {
	n := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(n, sqrch)
	go calcCubes(n, cubech)
	squares, cubes := <-sqrch,<-cubech
	fmt.Println("Final output: ", squares+cubes)
}
