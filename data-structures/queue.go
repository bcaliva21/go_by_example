package main
import "fmt"

type Queue struct {
	items []int
	size int
}

func (q *Queue) Enqueue (item int) int {
	q.items = append(q.items, item)
	if q.size > 0 {
		q.size++
	} else {
		q.size = 1
	}
	return q.size
}

func (q *Queue) Dequeue () int {
	if q.size < 1 {
		fmt.Println("cannot dequeue if queue is empty")
		return -1
	}
	dequeued := q.items[0]
	q.items = q.items[1:]
	q.size--
	return dequeued
}

func main() {
	queue := &Queue{}
	
	fmt.Println(queue.Enqueue(10))
	fmt.Println(queue.Enqueue(20))
	fmt.Println(queue.Enqueue(30))

	fmt.Println(queue)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue)

	fmt.Println(queue.Enqueue(100))
	fmt.Println(queue)
}
