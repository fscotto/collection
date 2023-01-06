package collection

import (
	"fmt"
	"strings"
)

var (
	ErrEmptyQueue = fmt.Errorf("empty queue")
)

type Queue[E any] struct {
	items []E
}

func NewQueue[E any](items ...E) *Queue[E] {
	if items == nil {
		items = make([]E, 0)
	}
	return &Queue[E]{items}
}

func (q *Queue[E]) Empty() bool {
	return q.Size() == 0
}

func (q *Queue[E]) Size() int {
	return len(q.items)
}

func (q *Queue[E]) Enqueue(item E) {
	q.items = append(q.items, item)
}

func (q *Queue[E]) Dequeue() (E, error) {
	if q.Empty() {
		return *new(E), ErrEmptyQueue
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *Queue[E]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < q.Size(); i++ {
		var s string
		item := q.items[i]
		if i >= q.Size()-1 {
			s = fmt.Sprintf("%v", item)
		} else {
			s = fmt.Sprintf("%v, ", item)
		}
		sb.WriteString(s)
	}
	sb.WriteString("]")
	return sb.String()
}
