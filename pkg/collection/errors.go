package collection

import "fmt"

var (
	ErrPositionNegative = fmt.Errorf("position can not be negative")
	ErrEmptyCollection  = fmt.Errorf("this collection is empty")
	ErrNodeNotFound     = fmt.Errorf("node not found")
)

type ErrIndexOutOfBound struct {
	index int
	size  int
}

func (e ErrIndexOutOfBound) Error() string {
	return fmt.Sprintf("index %d out of bound from range %d and %d", e.index, 0, e.size-1)
}

type ErrItemNotFound struct {
	item any
}

func (e ErrItemNotFound) Error() string {
	return fmt.Sprintf("item %v not found", e.item)
}
