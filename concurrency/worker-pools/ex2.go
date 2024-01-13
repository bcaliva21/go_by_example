package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id        int
	randomnum int
}

type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func calcSum(n int) int {
	s := 0
	num := n
	for num != 0 {
		d := num % 10
		s += d
		num /= 10
	}
	time.Sleep(2 * time.Second)
	return s
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		out := Result{job, calcSum(job.randomnum)}
		results <- out
	}
	wg.Done()
}

func createWorkerPool(numOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(numOfJobs int) {
	for i := 0; i < numOfJobs; i++ {
		randomnum := rand.Intn(1000)
		job := Job{i, randomnum}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.randomnum, result.sumofdigits)
	}
	done <- true
}

func main() {
	start := time.Now()
	numOfJobs := 100
	go allocate(numOfJobs)
	done := make(chan bool)
	go result(done)
	numOfWorkers := 100
	createWorkerPool(numOfWorkers)
	<-done
	end := time.Now()
	diff := end.Sub(start)
	fmt.Println("total time taken: ", diff.Seconds(), " seconds")
}

