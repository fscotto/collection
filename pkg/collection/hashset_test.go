package collection

import "testing"

func TestHashSet_Empty(t *testing.T) {
	useCases := []struct {
		description string
		set         *HashSet[int]
		want        bool
	}{
		{description: "empty set", set: NewHashSet[int](), want: true},
		{description: "set with items", set: NewHashSet[int](1, 2, 3), want: false},
	}

	for _, tt := range useCases {
		result := tt.set.Empty()
		if result != tt.want {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestHashSet_Size(t *testing.T) {
	useCases := []struct {
		description string
		set         *HashSet[int]
		want        int
	}{
		{description: "empty set", set: NewHashSet[int](), want: 0},
		{description: "set with items", set: NewHashSet[int](1, 2, 3), want: 3},
	}

	for _, tt := range useCases {
		result := tt.set.Size()
		if result != tt.want {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestHashSet_Push(t *testing.T) {
	useCases := []struct {
		description string
		original    *HashSet[int]
		modified    *HashSet[int]
		item        int
	}{
		{description: "empty set push new item",
			original: NewHashSet[int](),
			modified: NewHashSet[int](1),
			item:     1},
		{description: "set with items push new item",
			original: NewHashSet[int](1, 2),
			modified: NewHashSet[int](1, 2, 3),
			item:     3},
	}

	for _, tt := range useCases {
		tt.original.Push(tt.item)
		if tt.original.Size() != tt.modified.Size() {
			t.Errorf("test: %s want %v got %v", tt.description, tt.modified, tt.original)
		}
	}
}

func TestHashSet_Contains(t *testing.T) {
	useCases := []struct {
		description string
		set         *HashSet[int]
		item        int
		want        bool
	}{
		{description: "empty set", set: NewHashSet[int](), item: 0, want: false},
		{description: "set with items seek item not in set", set: NewHashSet[int](1, 2, 3), item: 0, want: false},
		{description: "set with items seek item in set", set: NewHashSet[int](1, 2, 3), item: 1, want: true},
	}

	for _, tt := range useCases {
		result := tt.set.Contains(tt.item)
		if result != tt.want {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestHashSet_Delete(t *testing.T) {
	useCases := []struct {
		description string
		original    *HashSet[int]
		modified    *HashSet[int]
		item        int
		err         error
	}{
		{description: "singleton set delete item",
			original: NewHashSet[int](1),
			modified: NewHashSet[int](),
			item:     1},
		{description: "set with items delete item",
			original: NewHashSet[int](1, 2, 3),
			modified: NewHashSet[int](1, 2),
			item:     3},
		{description: "set with items delete item not in set",
			original: NewHashSet[int](1, 2, 3),
			modified: NewHashSet[int](1, 2, 3),
			item:     9},
	}

	for _, tt := range useCases {
		err := tt.original.Delete(tt.item)
		if tt.original.Size() != tt.modified.Size() || err != tt.err {
			t.Errorf("test: %s want %v got %v", tt.description, tt.modified, tt.original)
		}
	}
}
