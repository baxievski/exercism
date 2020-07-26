// Package linkedlist implements a solution for the linked-list exercism problem
package linkedlist

import "fmt"

// Node represents a node in a doubly linked list
type Node struct {
	Val  interface{}
	prev *Node
	next *Node
}

// List represents a doubly linked list
type List struct {
	head *Node
	tail *Node
}

// ErrEmptyList represents the error when the pushing or popping an empty list
var ErrEmptyList = fmt.Errorf("list is empty")

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

// NewList creates a doubly linked list
func NewList(args ...interface{}) *List {
	l := &List{}
	for _, a := range args {
		l.PushBack(a)
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
	n := Node{Val: v, prev: l.tail, next: nil}
	if l.tail == nil {
		l.head = &n
	} else {
		l.tail.next = &n
	}
	l.tail = &n
}

// PushFront pushes a node at the start of the list
func (l *List) PushFront(v interface{}) {
	n := &Node{Val: v, next: l.head, prev: nil}
	if l.head == nil {
		l.tail = n
	} else {
		l.head.prev = n
	}
	l.head = n
}

// PopFront pops a node from the start of the list
func (l *List) PopFront() (interface{}, error) {
	n := l.First()
	if n == nil {
		return 0, ErrEmptyList
	}
	l.head = n.Next()
	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	return n.Val, nil
}

// PopBack pops a node from the end of the list
func (l *List) PopBack() (interface{}, error) {
	n := l.Last()
	if n == nil {
		return 0, ErrEmptyList
	}
	l.tail = n.Prev()
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	return n.Val, nil
}

// Reverse reverses the linked list
func (l *List) Reverse() {
	for n := l.First(); n != nil; n = n.Prev() {
		n.prev, n.next = n.next, n.prev
	}
	l.head, l.tail = l.tail, l.head
}
