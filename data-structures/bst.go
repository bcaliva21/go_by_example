package main

import "fmt"

type Node struct {
	Key int
	Left *Node
	Right *Node
}

func (n *Node) Insert (k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}

	}
}

func (n *Node) Search (k int) bool {
	cur := n
	for cur.Key != k {
		if cur.Key > k {
			if cur.Left == nil {
				return false
			} else {
				cur = cur.Left
			}
		} else {
			if cur.Right == nil {
				return false
			} else {
				cur = cur.Right
			}
		}
	}

	return true
}

func main() {
	tree := &Node{Key:100}

	tree.Insert(50)
	tree.Insert(150)
	tree.Insert(25)

	fmt.Println(tree.Search(10))
	fmt.Println(tree.Search(50))
}

