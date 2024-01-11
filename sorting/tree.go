package main

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func (n *Node) Insert(v int) {
	if n.Value < v {
		if n.Right == nil {
			n.Right = &Node{Value:v}
		} else {
			n.Right.Insert(v)
		}
	} else {
		if n.Left == nil {
			n.Left = &Node{Value:v}
		} else {
			n.Left.Insert(v)
		}
	}
}

func (n *Node) PrintPreOrder() []int {
	o := []int{}
	if n == nil {
		return []int{}
	}

	if n.Left != nil {
		o = append(o,n.Left.PrintPreOrder()...)
	}
	o = append(o,n.Value)

	if n.Right != nil {
		o = append(o,n.Right.PrintPreOrder()...)
	}

	return o
}

func TreeSort(a []int) []int {
	t := &Node{Value:a[0]}
	for i := 1; i < len(a); i++ {
		t.Insert(a[i])
	}

	s := t.PrintPreOrder()
	return s
}

func main() {
	a := []int{10,0,16,6,4,2,20,8,3,17}

	fmt.Println(a)
	fmt.Println(TreeSort(a))
}

