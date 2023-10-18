package main

import (
	"fmt"
	"time"
)

// func f(from string) {
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(from, ":", i)
// 	}
// }

func l() {
	for i := 1; i < 101; i++ {
		fmt.Println(i)
	}
}

func main() {
	l()

	time.Sleep(time.Microsecond / 10)
	fmt.Println("done")
}

