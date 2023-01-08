package collection

import (
	"fmt"
	"strings"
)

// Stack is an structure with method for handle underlying slice as a stack.
type Stack[E any] struct {
	items []E
}

// NewStack is a constructor function for stack
func NewStack[E any](items ...E) *Stack[E] {
	stack := &Stack[E]{}
	for _, item := range items {
		stack.Push(item)
	}
	return stack
}

// Size returns the number of items in the stack.
func (s *Stack[E]) Size() int {
	return len(s.items)
}

// Empty returns true if the stack is empty, otherwise false.
func (s *Stack[E]) Empty() bool {
	return s.Size() == 0
}

// Top get first item in the stack structure, but
// if stack is empty well this method return ErrEmptyCollection error.
func (s *Stack[E]) Top() (E, error) {
	if s.Empty() {
		return *new(E), ErrEmptyCollection
	}
	return s.items[s.Size()-1], nil
}

// Pop get and remove first item in the stack structure, but
// if stack is empty well this method return ErrEmptyCollection error.
func (s *Stack[E]) Pop() (E, error) {
	item, err := s.Top()
	if err != nil {
		return *new(E), err
	}
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Push adds an item to the top of the stack.
func (s *Stack[E]) Push(item E) error {
	s.items = append(s.items, item)
	return nil
}

func (s *Stack[E]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < s.Size(); i++ {
		var str string
		item := s.items[i]
		if i >= s.Size()-1 {
			str = fmt.Sprintf("%v", item)
		} else {
			str = fmt.Sprintf("%v, ", item)
		}
		sb.WriteString(str)
	}
	sb.WriteString("]")
	return sb.String()
}
