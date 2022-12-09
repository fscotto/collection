package collection_test

import (
	"testing"

	"asd.me/pkg/collection"
)

type pair struct {
	value int
	err   error
}

var staticTests = []struct {
	description string
	input       *collection.Stack[int]
	empty       bool
	size        int
	top         pair
}{
	{"empty stack", newStack([]int{}), true, 0, pair{0, collection.ErrEmptyStack}},
	{"stack with items", newStack([]int{1, 2, 3}), false, 3, pair{3, nil}},
}

var popTests = []struct {
	description string
	input       *collection.Stack[int]
	item        pair
}{
	{"empty stack pop item", newStack([]int{}), pair{0, collection.ErrEmptyStack}},
	{"stack with items pop item", newStack([]int{1, 2, 3}), pair{3, nil}},
}

var pushTests = []struct {
	description string
	items       []int
}{
	{"zero item pushed", []int{}},
	{"one item pushed", []int{1}},
	{"five item pushed", []int{1, 2, 3, 4, 5}},
}

func TestStackStatus(t *testing.T) {
	for i, tt := range staticTests {
		t.Logf("Test %v: %s\n", i, tt.description)

		stack := tt.input
		if tt.empty != stack.Empty() {
			t.Errorf("stack empty %v, want %v\n", stack.Empty(), tt.empty)
		}

		if tt.size != stack.Size() {
			t.Errorf("stack size %v, want %v\n", stack.Size(), tt.size)
		}

		value, err := stack.Top()
		if tt.top.err != nil && tt.top.err != err {
			t.Errorf("stack top is (%v, %v), want (%v, %v)\n", value, err, tt.top.value, tt.top.err)
		}
	}
}

func TestPop(t *testing.T) {
	for i, tt := range popTests {
		t.Logf("Test %v: %s\n", i, tt.description)

		result, err := tt.input.Pop()
		if result != tt.item.value {
			t.Errorf("found result: %v, want %v\n", result, tt.item.value)
		}

		if err != tt.item.err {
			t.Errorf("found error: %v, want %v\n", err, tt.item.err)
		}
	}
}

func TestPush(t *testing.T) {
	for i, tt := range pushTests {
		t.Logf("Test %v: %s\n", i, tt.description)

		stack := &collection.Stack[int]{}
		for _, item := range tt.items {
			stack.Push(item)
		}

		if stack.Size() != len(tt.items) {
			t.Errorf("stack size %v, want %v\n", stack.Size(), len(tt.items))
		}
	}
}

func newStack(slice []int) *collection.Stack[int] {
	stack := &collection.Stack[int]{}
	for _, x := range slice {
		stack.Push(x)
	}
	return stack
}
