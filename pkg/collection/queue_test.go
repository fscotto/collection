package collection

import (
	"reflect"
	"testing"
)

func TestQueue_Empty(t *testing.T) {
	useCases := []struct {
		description string
		queue       *Queue[int]
		want        bool
	}{
		{description: "empty queue", queue: NewQueue[int](), want: true},
		{description: "non empty queue", queue: NewQueue(1), want: false},
	}

	for _, tt := range useCases {
		result := tt.queue.Empty()
		if result != tt.want {
			t.Errorf("test: %s, want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestQueue_Size(t *testing.T) {
	useCases := []struct {
		description string
		queue       *Queue[int]
		want        int
	}{
		{description: "empty queue", queue: NewQueue[int](), want: 0},
		{description: "non empty queue", queue: NewQueue(1), want: 1},
	}

	for _, tt := range useCases {
		result := tt.queue.Size()
		if result != tt.want {
			t.Errorf("test: %s, want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestQueue_Enqueue(t *testing.T) {
	useCases := []struct {
		description string
		original    *Queue[int]
		modified    *Queue[int]
		item        int
	}{
		{description: "enqueue one item in empty queue", original: NewQueue[int](), modified: NewQueue(1), item: 1},
		{description: "enqueue one item in full queue", original: NewQueue(1, 2, 3), modified: NewQueue(1, 2, 3, 4), item: 4},
	}

	for _, tt := range useCases {
		tt.original.Enqueue(tt.item)
		if !reflect.DeepEqual(tt.original.items, tt.modified.items) {
			t.Errorf("test: %s, want %v got %v", tt.description, tt.modified, tt.original)
		}
	}
}

func TestQueue_Dequeue(t *testing.T) {
	useCases := []struct {
		description string
		original    *Queue[int]
		modified    *Queue[int]
		item        int
		err         error
	}{
		{description: "dequeue one item in empty queue",
			original: NewQueue[int](),
			modified: NewQueue[int](),
			item:     0,
			err:      ErrEmptyQueue},
		{description: "dequeue one item in single item queue",
			original: NewQueue(1),
			modified: NewQueue[int](),
			item:     1,
			err:      nil},
		{description: "dequeue one item in full queue",
			original: NewQueue(1, 2, 3),
			modified: NewQueue(2, 3),
			item:     1,
			err:      nil},
	}

	for _, tt := range useCases {
		_, err := tt.original.Dequeue()
		if !reflect.DeepEqual(tt.original.items, tt.modified.items) || err != tt.err {
			t.Errorf("test: %s, want %v got %v", tt.description, tt.modified, tt.original)
		}
	}
}
