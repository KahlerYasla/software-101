package datastuctures

import "errors"

// Stack structure with a slice to hold the elements.
type Stack struct {
	elements []int
}

// Push adds an element to the top of the Stack.
func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element of the Stack.
// Returns an error if the stack is empty.
func (s *Stack) Pop() (int, error) {
	if len(s.elements) == 0 {
		return 0, errors.New("stack is empty")
	}

	// Get the top element (last item in the slice).
	topElement := s.elements[len(s.elements)-1]

	// Remove the top element by slicing off the last element.
	s.elements = s.elements[:len(s.elements)-1]

	return topElement, nil
}

// Peek returns the top element of the Stack without removing it.
// Returns error if the Stack is empty
func (s *Stack) Peek() (int, error) {
	if len(s.elements) == 0 {
		return 0, errors.New("stack is empty")
	}

	return s.elements[len(s.elements)-1], nil
}

// IsEmpty returns true if the Stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	return len(s.elements)
}
