package queue

import "errors"

// Queue structure with a slice to hold the elements.
type Queue struct {
	elements []int
}

// Enqueue adds an element to the back of the Queue.
func (q *Queue) Enqueue(val int) {
	q.elements = append(q.elements, val)
}

// Dequeue removes and returns the front element of the Queue.
// Returns an error if the Queue is empty.
func (q *Queue) Dequeue() (int, error) {
	if len(q.elements) == 0 {
		return 0, errors.New("queue is empty")
	}

	// Get the front element
	frontElement := q.elements[0]

	// Remove the front element by slicing off the first element.
	q.elements = q.elements[1:]

	return frontElement, nil
}

// Peek returns the front element of the queue without removing it.
// Returns an error if the Queue is empty.
func (q *Queue) Peek() (int, error) {
	if len(q.elements) == 0 {
		return 0, errors.New("queue is empty")
	}

	return q.elements[0], nil
}

// IsEmpty returns true if the Queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

// Size returns the number of elements in the Queue
func (q *Queue) Size() int {
	return len(q.elements)
}
