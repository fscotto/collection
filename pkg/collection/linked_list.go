package collection

import (
	"fmt"
	"reflect"
	"strings"
)

type node[E any] struct {
	value E
	prev  *node[E]
	next  *node[E]
}

type listIterator[E any] struct {
	currentNode  *node[E]
	currentIndex int
}

func (it *listIterator[E]) HasNext() bool {
	return it.currentNode != nil
}

func (it *listIterator[E]) Next() E {
	curr := it.currentNode
	it.currentNode = curr.next
	it.currentIndex++
	return curr.value
}

func (it *listIterator[E]) NextWithIndex() (int, E) {
	curr := it.currentNode
	index := it.currentIndex
	it.currentNode = curr.next
	it.currentIndex++
	return index, curr.value
}

type emptyListIterator[E any] struct{}

func (it *emptyListIterator[E]) HasNext() bool {
	return false
}

func (it *emptyListIterator[E]) Next() E {
	return *new(E)
}

func (it *emptyListIterator[E]) NextWithIndex() (int, E) {
	return 0, *new(E)
}

type LinkedList[E any] struct {
	head *node[E]
	size int
}

func NewLinkedList[E any](items ...E) *LinkedList[E] {
	lst := &LinkedList[E]{}
	for _, item := range items {
		lst.PushBack(item)
	}
	return lst
}

func (l *LinkedList[E]) Iterator() Iterator[E] {
	if l.Empty() {
		return &emptyListIterator[E]{}
	}
	return &listIterator[E]{l.head, 0}
}

func (l *LinkedList[E]) Empty() bool {
	return l.head == nil
}

func (l *LinkedList[E]) Size() int {
	return l.size
}

func (l *LinkedList[E]) Back() (E, error) {
	return l.GetAt(l.size - 1)
}

func (l *LinkedList[E]) Front() (E, error) {
	return l.GetAt(0)
}

func (l *LinkedList[E]) GetAt(pos int) (E, error) {
	node, err := l.findNode(pos)
	if err != nil {
		return *new(E), err
	}
	return node.value, nil
}

// findNode returns node at given position from linked list
func (l *LinkedList[E]) findNode(pos int) (*node[E], error) {
	if l.Empty() {
		return nil, ErrEmptyCollection
	}

	ptr := l.head
	if pos < 0 {
		return nil, ErrPositionNegative
	}

	if pos > (l.size - 1) {
		return nil, ErrIndexOutOfBound{pos, l.size}
	}

	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}

	return ptr, nil
}

func (l *LinkedList[E]) Contains(item E) bool {
	for it := l.Iterator(); it.HasNext(); {
		x := it.Next()
		if reflect.DeepEqual(item, x) {
			return true
		}
	}
	return false
}

func (l *LinkedList[E]) Index(item E) (int, error) {
	for it := l.Iterator(); it.HasNext(); {
		i, x := it.NextWithIndex()
		if reflect.DeepEqual(item, x) {
			return i, nil
		}
	}
	return 0, ErrItemNotFound{item}
}

func (l *LinkedList[E]) Push(item E) {
	l.PushBack(item)
}

func (l *LinkedList[E]) PushFront(item E) {
	l.PushAt(item, 0)
}

func (l *LinkedList[E]) PushBack(item E) {
	l.PushAt(item, l.size)
}

// PushAt inserts new node at given position
func (l *LinkedList[E]) PushAt(item E, pos int) {
	// create a new node
	newNode := node[E]{value: item}

	// validate the position
	if pos < 0 {
		return
	}

	if pos == 0 {
		oldNode := l.head
		newNode.next = oldNode
		l.head = &newNode
		if oldNode != nil {
			oldNode.prev = l.head
		}
		l.size++
		return
	}
	if pos > l.size {
		return
	}
	n, _ := l.findNode(pos)
	newNode.next = n
	var prevNode *node[E]
	if n != nil {
		prevNode = n.prev
	} else {
		prevNode, _ = l.findNode(pos - 1)
	}
	prevNode.next = &newNode
	newNode.prev = prevNode
	l.size++
}

func (l *LinkedList[E]) Delete(item E) error {
	pos, err := l.Index(item)
	if err != nil {
		return err
	}
	return l.DeleteAt(pos)
}

// DeleteAt deletes node at given position from linked list
func (l *LinkedList[E]) DeleteAt(pos int) error {
	// validate the position
	if pos < 0 {
		return ErrPositionNegative
	}

	if l.size == 0 {
		return ErrEmptyCollection
	}

	if pos == 0 {
		// For first position not exists prev node
		myNode, err := l.findNode(pos)
		if err != nil {
			return err
		}
		l.head = myNode.next
		l.head.prev = nil
	} else {
		prevNode, _ := l.findNode(pos - 1)
		if prevNode == nil {
			return ErrNodeNotFound
		}
		myNode, err := l.findNode(pos)
		if err != nil {
			return err
		}
		prevNode.next = myNode.next
		if myNode.next != nil {
			myNode.next.prev = prevNode
		}
	}

	l.size--
	return nil
}

func (l *LinkedList[E]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < l.Size(); i++ {
		var s string
		item, _ := l.GetAt(i)
		if i >= l.Size()-1 {
			s = fmt.Sprintf("%v", item)
		} else {
			s = fmt.Sprintf("%v, ", item)
		}
		sb.WriteString(s)
	}
	sb.WriteString("]")
	return sb.String()
}
