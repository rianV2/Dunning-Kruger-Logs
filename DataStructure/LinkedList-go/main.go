package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) PushAtHead(value int) {
	node := &Node{value: value}

	if l.head == nil {
		l.tail = node
	} else {
		l.head.prev = node
		node.next = l.head
		// node.next.prev = node
		// l.head = &Node{value: value, next: l.tail, prev: l.head} //asign next as current head
	}

	l.head = node
}

func (l *List) PushAtTail(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (n *Node) Next() *Node {
	return n.next
}
func (n *Node) Prev() *Node {
	return n.prev
}

func main() {
	list := &List{}
	list.PushAtTail(2)
	list.PushAtTail(3)
	list.PushAtHead(4)
	list.PushAtTail(5)

	fmt.Println("first: ", list.First().value)

	n := list.First()
	for {
		fmt.Println(n.value, "|")

		if n.Next() == nil {
			break
		}
		n = n.Next()
	}

	// fmt.Println(list.First().Next().value)
	fmt.Println("last: ", list.Last().value)
	// for i := 0; i <= 2; i++ {
	// 	fmt.Println(list.First().value)
	// }
}
