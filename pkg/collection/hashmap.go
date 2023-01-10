package collection

import (
	"fmt"
	"reflect"
	"strings"
)

type HashMap[K comparable, V any] struct {
	table map[K]V
}

type Entry[K comparable, V any] struct {
	key   K
	value V
}

func NewEntry[K comparable, V any](key K, value V) *Entry[K, V] {
	return &Entry[K, V]{
		key:   key,
		value: value,
	}
}

func (e Entry[K, V]) Key() K {
	return e.key
}

func (e Entry[K, V]) Value() V {
	return e.value
}

func (e Entry[K, V]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	sb.WriteString(fmt.Sprintf("[%v, %v]", e.key, e.value))
	sb.WriteString("]")
	return sb.String()
}

func NewHashMap[K comparable, V any](entries ...*Entry[K, V]) *HashMap[K, V] {
	m := make(map[K]V)
	for _, e := range entries {
		m[e.key] = e.value
	}
	return &HashMap[K, V]{m}
}

func (h *HashMap[K, V]) Empty() bool {
	return h.Size() == 0
}

func (h *HashMap[K, V]) Size() int {
	return len(h.table)
}

func (h *HashMap[K, V]) Get(key K) (V, bool) {
	v, ok := h.table[key]
	return v, ok
}

func (h *HashMap[K, V]) Put(key K, value V) {
	h.table[key] = value
}

func (h *HashMap[K, V]) ContainsKey(key K) bool {
	_, ok := h.table[key]
	return ok
}

func (h *HashMap[K, V]) ContainsValue(value V) bool {
	for _, v := range h.table {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

func (h *HashMap[K, V]) Delete(key K) bool {
	l := h.Size()
	delete(h.table, key)
	return l != h.Size()
}

func (h *HashMap[K, V]) Keys() Set[K] {
	set := NewHashSet[K]()
	for k := range h.table {
		set.Push(k)
	}
	return set
}

func (h *HashMap[K, V]) Values() Collection[V] {
	lst := NewSlice[V]()
	for _, v := range h.table {
		lst.PushBack(v)
	}
	return lst
}

func (h *HashMap[K, V]) EntryList() Collection[*Entry[K, V]] {
	lst := NewSlice[*Entry[K, V]]()
	for k, v := range h.table {
		lst.PushBack(NewEntry(k, v))
	}
	return lst
}

func (h *HashMap[K, V]) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	lst := h.EntryList()
	for it := lst.Iterator(); it.HasNext(); {
		var s string
		i, item := it.NextWithIndex()
		if i >= h.Size()-1 {
			s = fmt.Sprintf("%v", item)
		} else {
			s = fmt.Sprintf("%v, ", item)
		}
		sb.WriteString(s)
	}
	sb.WriteString("}")
	return sb.String()
}
