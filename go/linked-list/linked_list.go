// Package linkedlist implements a solution for the linked-list exercism problem
package linkedlist

import "fmt"

// Node represents a node in a doubly linked list
type Node struct {
	Val  interface{}
	prev *Node
	next *Node
}

// Next gives the next node in the list
func (n *Node) Next() *Node {
	return n.next
}

// Prev gives the previous node in the list
func (n *Node) Prev() *Node {
	return n.prev
}

// First gives the first node in the list
func (n *Node) First() *Node {
	for n.prev != nil {
		n = n.prev
	}
	return n
}

// Last gives the last node in the list
func (n *Node) Last() *Node {
	for n.next != nil {
		n = n.next
	}
	return n
}

// List represents a doubly linked list
type List struct {
	head *Node
	tail *Node
}

// NewList creates a doubly linked list
func NewList(args ...interface{}) *List {
	l := &List{}
	for i, a := range args {
		if i == 0 {
			l.tail = &Node{Val: a}
			l.head = l.tail
			continue
		}
		l.tail.next = &Node{Val: a, prev: l.tail}
		l.tail = l.tail.next
	}
	return l
}

// First gives the first node in the list
func (l *List) First() *Node {
	return l.head
}

// Last gives the last node in the list
func (l *List) Last() *Node {
	return l.tail
}

// PushBack pushes a node at the end of the list
func (l *List) PushBack(v interface{}) {
	l.tail.next = &Node{Val: v, prev: l.tail}
}

// PushFront pushes a node at the start of the list
func (l *List) PushFront(v interface{}) {
	l.head = &Node{Val: v, next: l.head}
	l.head.next.prev = l.head
}

// PopFront pops a node from the start of the list
func (l *List) PopFront() (interface{}, error) {
	h := l.First()
	if h == nil {
		return 0, fmt.Errorf("attempt to pop an empty list")
	}
	l.head = l.head.Next()
	l.head.prev = nil
	return h.Val, nil
}

// PopBack pops a node from the end of the list
func (l *List) PopBack(v interface{}) (interface{}, error) {
	h := l.Last()
	if h == nil {
		return 0, fmt.Errorf("attempt to pop an empty list")
	}
	l.tail = l.tail.Prev()
	l.tail.next = nil
	return h.Val, nil

}

// Reverse reverses the linked list
func (l *List) Reverse() *List {
	for n := l.First(); n != nil; n = n.Prev() {
		n.prev, n.next = n.next, n.prev
	}
	l.head, l.tail = l.tail, l.head
	return l
}
