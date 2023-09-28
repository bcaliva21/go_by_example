package main
import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	length int
}

func (l *linkedList) Prepend (n *node) {
	temp := l.head
	l.head = n
	l.head.next = temp
	l.length++
}

func (l *linkedList) Print () {
	cur := l.head
	
	for cur != nil {
		fmt.Println(cur)
		cur = cur.next
	}
}

func (l *linkedList) RemoveTail () {
	cur := l.head
	tmp := l.head

	for cur.next != nil {
		tmp = cur
		cur = cur.next
	}

	tmp.next = nil
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 100}
	node2 := &node{data: 101}
	node3 := &node{data: 102}
	mylist.Prepend(node1)
	mylist.Prepend(node2)
	mylist.Prepend(node3)

	fmt.Println(mylist)

	mylist.Print()
	mylist.RemoveTail()
	mylist.Print()
}

