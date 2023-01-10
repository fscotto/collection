package collection

import (
	"fmt"
	"strings"
)

type HashSet[E comparable] struct {
	inner map[E]bool
}

func NewHashSet[E comparable](items ...E) *HashSet[E] {
	inner := make(map[E]bool)
	for _, item := range items {
		inner[item] = true
	}
	return &HashSet[E]{inner}
}

func (h *HashSet[E]) Iterator() Iterator[E] {
	// FIXME this is only hack for iterate over map
	items := NewSlice[E]()
	for k := range h.inner {
		items.Push(k)
	}
	return items.Iterator()
}

func (h *HashSet[E]) Empty() bool {
	return h.Size() == 0
}

func (h *HashSet[E]) Size() int {
	return len(h.inner)
}

func (h *HashSet[E]) Push(item E) {
	h.inner[item] = true
}

func (h *HashSet[E]) Contains(item E) bool {
	return h.inner[item]
}

func (h *HashSet[E]) Delete(item E) error {
	delete(h.inner, item)
	return nil
}

func (h *HashSet[E]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for it := h.Iterator(); it.HasNext(); {
		var s string
		i, item := it.NextWithIndex()
		if i >= h.Size()-1 {
			s = fmt.Sprintf("%v", item)
		} else {
			s = fmt.Sprintf("%v, ", item)
		}
		sb.WriteString(s)
	}
	sb.WriteString("]")
	return sb.String()
}
