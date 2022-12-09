package collection

import "fmt"

var (
	// ErrEmptyStack is error when you have an empty stack structure
	ErrEmptyStack = fmt.Errorf("this stack is empty")
)

// Stack is an structure with method for handle underlying slice as a stack.
type Stack[E any] struct {
	items []E
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
// if stack is empty well this method return ErrEmptyStack error.
func (s *Stack[E]) Top() (E, error) {
	if s.Empty() {
		return *new(E), ErrEmptyStack
	}
	return s.items[s.Size()-1], nil
}

// Pop get and remove first item in the stack structure, but
// if stack is empty well this method return ErrEmptyStack error.
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
