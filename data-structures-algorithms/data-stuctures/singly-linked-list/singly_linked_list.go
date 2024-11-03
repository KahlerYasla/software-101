package singlylinkedlist

import (
	"errors"
	"fmt"
)

// Node represents an element in the linked list.
type Node struct {
	Val  int
	Next *Node
}

// SinglyLinkedList represents the linked list with a head and a tail.
type SinglyLinkedList struct {
	Head *Node
	Tail *Node
}

// AddToFront adds a new node with the given value to the front of the list.
func (l *SinglyLinkedList) AddToFront(val int) {
	newNode := &Node{Val: val}
	if l.Head == nil { // List is empty
		l.Head = newNode
		l.Tail = newNode
	} else { // Not empty
		newNode.Next = l.Head
		l.Tail = newNode
	}
}

// AddToEnd adds a new node with the given value to the end of the list.
//
//	Parameters:
//	> val int: "the element to add at the end of the list"
func (l *SinglyLinkedList) AddToEnd(val int) {
	newNode := &Node{Val: val}
	if l.Tail == nil { // List is empty
		l.Head = newNode
		l.Tail = newNode
	} else { // Not empty
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

// RemoveFromFront removes the node from the front of the list.
func (l *SinglyLinkedList) RemoveFromFront() (int, error) {
	if l.Head == nil {
		return 0, errors.New("list is empty")
	}

	val := l.Head.Val
	l.Head = l.Head.Next
	if l.Head == nil {
		l.Tail = nil
	}

	return val, nil
}

// RemoveFromEnd removes the node from the end of the list.
//
// Returns:
//
//	val int: last node after remove
//	err error: if any error occured
func (l *SinglyLinkedList) RemoveFromEnd() (int, error) {
	if l.Head == nil {
		return 0, errors.New("list is empty")
	}

	if l.Head == l.Tail { // If there is a single node left
		val := l.Head.Val
		l.Head = nil
		l.Tail = nil
		return val, nil
	}

	// Traverse to the second-last node
	current := l.Head
	for current.Next != l.Tail {
		current = current.Next
	}
	val := l.Tail.Val
	l.Tail = current
	l.Tail.Next = nil
	return val, nil
}

// Display prints all the elements in the list
func (l *SinglyLinkedList) Display() {
	current := l.Head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}
