package collection

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

type sliceIterator[E any] struct {
	index int
	slice []E
}

func (it *sliceIterator[E]) HasNext() bool {
	return it.index < len(it.slice)
}

func (it *sliceIterator[E]) Next() E {
	item := it.slice[it.index]
	it.index += 1
	return item
}

func (it *sliceIterator[E]) NextWithIndex() (int, E) {
	index := it.index
	item := it.slice[index]
	it.index += 1
	return index, item
}

type Slice[E any] struct {
	inner []E
}

func NewSlice[E any](items ...E) *Slice[E] {
	return &Slice[E]{items}
}

func (s *Slice[E]) Iterator() Iterator[E] {
	if s.Empty() {
		return &emptyListIterator[E]{}
	}
	return &sliceIterator[E]{0, s.inner}
}

func (s *Slice[E]) Empty() bool {
	return s.Size() == 0
}

func (s *Slice[E]) Size() int {
	return len(s.inner)
}

func (s *Slice[E]) Back() (E, error) {
	return s.GetAt(s.Size() - 1)
}

func (s *Slice[E]) Front() (E, error) {
	return s.GetAt(0)
}

func (s *Slice[E]) GetAt(pos int) (E, error) {
	if s.Empty() {
		return *new(E), ErrEmptyCollection
	}
	if pos < 0 {
		return *new(E), ErrPositionNegative
	}
	if pos < 0 || pos >= s.Size() {
		return *new(E), ErrIndexOutOfBound{pos, s.Size()}
	}
	return s.inner[pos], nil
}

func (s *Slice[E]) Push(item E) {
	s.PushBack(item)
}

func (s *Slice[E]) PushBack(item E) {
	pos := int(math.Max(0.0, float64(s.Size())))
	s.PushAt(item, pos)
}

func (s *Slice[E]) PushFront(item E) {
	s.PushAt(item, 0)
}

func (s *Slice[E]) PushAt(item E, pos int) {
	s.inner = insert(s.inner, item, pos)
}

func insert[E any](slice []E, item E, pos int) []E {
	// found https://stackoverflow.com/questions/46128016/insert-a-value-in-a-slice-at-a-given-index
	n := len(slice)
	if pos < 0 {
		pos = (pos%n + n) % n
	}
	switch {
	case pos == n: // nil or empty slice or after last element
		return append(slice, item)

	case pos < n: // pos < len(slice)
		slice = append(slice[:pos+1], slice[pos:]...)
		slice[pos] = item
		return slice

	case pos < cap(slice): // pos > len(slice)
		slice = slice[:pos+1]
		var zero E
		for i := n; i < pos; i++ {
			slice[i] = zero
		}
		slice[pos] = item
		return slice

	default:
		b := make([]E, pos+1) // malloc
		if n > 0 {
			copy(b, slice)
		}
		b[pos] = item
		return b
	}
}

func (s *Slice[E]) Delete(item E) error {
	pos, err := s.Index(item)
	if err != nil {
		return err
	}
	return s.DeleteAt(pos)
}

func (s *Slice[E]) DeleteAt(pos int) error {
	if s.Empty() {
		return ErrEmptyCollection
	}
	s.inner = append(s.inner[:pos], s.inner[pos+1:]...)
	return nil
}

func (s *Slice[E]) Contains(item E) bool {
	for _, elem := range s.inner {
		if reflect.DeepEqual(elem, item) {
			return true
		}
	}
	return false
}

func (s *Slice[E]) Index(item E) (int, error) {
	for i, x := range s.inner {
		if reflect.DeepEqual(x, item) {
			return i, nil
		}
	}
	return 0, ErrItemNotFound{item}
}

func (s *Slice[E]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < s.Size(); i++ {
		var str string
		item, _ := s.GetAt(i)
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
