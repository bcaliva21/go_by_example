package main
import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	size := len(s.items) - 1
	popped := s.items[size]
	s.items = s.items[:size]
	return popped
}

// main
func main() {
	mystack := &Stack{}

	fmt.Println(mystack)
	mystack.Push(10)
	mystack.Push(20)
	mystack.Push(30)
	fmt.Println(mystack)

	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}
