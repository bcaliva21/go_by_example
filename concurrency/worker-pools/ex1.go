package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	num := 3
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines have finished executing")
	fmt.Println("--main terminated--")
}
