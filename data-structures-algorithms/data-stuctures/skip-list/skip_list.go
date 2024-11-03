package skiplist

import (
	"fmt"
	"math/rand"
)

// Node represents an element in the SkipList.
type Node struct {
	Val     int
	Forward []*Node // Forward links (next nodes) at different levels.
}

// SkipList represents the skip list structure.
type SkipList struct {
	Head        *Node
	Levels      int
	MaxLevel    int
	Probability float64
}

// NewSkipList creates a new skip list with the given maximum level and probability.
func NewSkipList(maxLevel int, probability float64) *SkipList {
	head := &Node{Forward: make([]*Node, maxLevel+1)}
	return &SkipList{
		Head:        head,
		Levels:      0,
		MaxLevel:    maxLevel,
		Probability: probability, // Between 0 and 1
	}
}

// RandomLevel generates a random level for a new node.
func (sl *SkipList) RandomLevel() int {
	level := 0
	for rand.Float64() < sl.Probability && level < sl.MaxLevel {
		level++
	}
	return level
}

// Insert adds a new value to the SkipList.
func (sl *SkipList) Insert(val int) {
	update := make([]*Node, sl.MaxLevel+1)
	current := sl.Head

	// Find position to insert the new node.
	for i := sl.Levels; i >= 0; i-- {
		for current.Forward != nil && current.Forward[i].Val < val {
			current = current.Forward[i]
		}
		update[i] = current
	}

	// Move to the next level.
	current = current.Forward[0]

	// Insert only if value is not already present.
	if current == nil || current.Val != val {
		newLevel := sl.RandomLevel()
		if newLevel > sl.Levels {
			for i := sl.Levels + 1; i <= newLevel; i++ {
				update[i] = sl.Head
			}
			sl.Levels = newLevel
		}

		newNode := &Node{Val: val, Forward: make([]*Node, newLevel+1)}
		for i := 0; i <= newLevel; i++ {
			newNode.Forward[i] = update[i].Forward[i]
			update[i].Forward[i] = newNode
		}
	}
}

// Search finds a value in the skip list.
func (sl *SkipList) Search(val int) bool {
	current := sl.Head
	for i := sl.Levels; i >= 0; i-- {
		for current.Forward[i] != nil && current.Forward[i].Val < val {
			current = current.Forward[i]
		}
	}

	current = current.Forward[0]
	return current != nil && current.Val == val
}

// Delete removes a value from the skip list.
func (sl *SkipList) Delete(val int) {
	update := make([]*Node, sl.MaxLevel+1)
	current := sl.Head

	for i := sl.Levels; i >= 0; i-- {
		for current.Forward[i] != nil && current.Forward[i].Val < val {
			current = current.Forward[i]
		}
		update[i] = current
	}

	current = current.Forward[0]
	if current != nil && current.Val == val {
		for i := 0; i <= sl.Levels; i++ {
			if update[i].Forward[i] != current {
				break
			}
			update[i].Forward[i] = current.Forward[i]
		}

		// Adjust the level of the skip list
		for sl.Levels > 0 && sl.Head.Forward[sl.Levels] == nil {
			sl.Levels--
		}
	}
}

// Display prints the skip list
func (sl *SkipList) Display() {
	for i := sl.Levels; i >= 0; i-- {
		fmt.Printf("Level %d: ", i)
		current := sl.Head.Forward[i]
		for current != nil {
			fmt.Printf("%d ", current.Val)
			current = current.Forward[i]
		}
		fmt.Println()
	}
}
