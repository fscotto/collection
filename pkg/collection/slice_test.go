package collection

import (
	"reflect"
	"testing"
)

func TestSlice_Empty(t *testing.T) {
	useCases := []struct {
		description string
		slice       *Slice[int]
		want        bool
	}{
		{description: "empty slice", slice: NewSlice[int](), want: true},
		{description: "full slice", slice: NewSlice[int](1, 2, 3), want: false},
	}

	for _, tt := range useCases {
		result := tt.slice.Empty()
		if result != tt.want {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestSlice_Size(t *testing.T) {
	useCases := []struct {
		description string
		slice       *Slice[int]
		want        int
	}{
		{description: "empty slice", slice: NewSlice[int](), want: 0},
		{description: "full slice", slice: NewSlice[int](1, 2, 3), want: 3},
	}

	for _, tt := range useCases {
		result := tt.slice.Size()
		if result != tt.want {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestSlice_GetAt(t *testing.T) {
	useCases := []struct {
		description string
		slice       *Slice[int]
		pos         int
		item        int
		err         error
	}{
		{description: "get item on empty slice", slice: NewSlice[int](), pos: 0, item: 0, err: ErrEmptyCollection},
		{description: "get item on full slice with out of bound index", slice: NewSlice[int](1, 2, 3), pos: 5, item: 0, err: ErrIndexOutOfBound{5, 3}},
		{description: "get item on full slice with negative index", slice: NewSlice[int](1, 2, 3), pos: -1, item: 0, err: ErrPositionNegative},
		{description: "get item on full slice in first position", slice: NewSlice[int](1, 2, 3), pos: 0, item: 1, err: nil},
		{description: "get item on full slice in middle position", slice: NewSlice[int](1, 2, 3), pos: 1, item: 2, err: nil},
		{description: "get item on full slice in last position", slice: NewSlice[int](1, 2, 3), pos: 2, item: 3, err: nil},
	}

	for _, tt := range useCases {
		result, err := tt.slice.GetAt(tt.pos)
		if result != tt.item || err != tt.err {
			t.Errorf("test: %s want {%v, %v} got {%v, %v}", tt.description, tt.item, tt.err, result, err)
		}
	}
}

func TestSlice_Back(t *testing.T) {
	useCases := []struct {
		description string
		slice       *Slice[int]
		pos         int
		item        int
		err         error
	}{
		{description: "get item on empty slice", slice: NewSlice[int](), pos: 0, item: 0, err: ErrEmptyCollection},
		{description: "get item on full slice in last position", slice: NewSlice[int](1, 2, 3), pos: 2, item: 3, err: nil},
	}

	for _, tt := range useCases {
		result, err := tt.slice.Back()
		if result != tt.item || err != tt.err {
			t.Errorf("test: %s want {%v, %v} got {%v, %v}", tt.description, tt.item, tt.err, result, err)
		}
	}
}

func TestSlice_Front(t *testing.T) {
	useCases := []struct {
		description string
		slice       *Slice[int]
		pos         int
		item        int
		err         error
	}{
		{description: "get item on empty slice", slice: NewSlice[int](), pos: 0, item: 0, err: ErrEmptyCollection},
		{description: "get item on full slice in first position", slice: NewSlice[int](1, 2, 3), pos: 2, item: 1, err: nil},
	}

	for _, tt := range useCases {
		result, err := tt.slice.Front()
		if result != tt.item || err != tt.err {
			t.Errorf("test: %s want {%v, %v} got {%v, %v}", tt.description, tt.item, tt.err, result, err)
		}
	}
}

func TestSlice_PushAt(t *testing.T) {
	useCases := []struct {
		description string
		original    *Slice[int]
		modified    *Slice[int]
		item        int
		pos         int
	}{
		{description: "push item in empty slice",
			original: NewSlice[int](),
			modified: NewSlice(1),
			item:     1,
			pos:      0},
		{description: "push item in slice with item in first position",
			original: NewSlice(1, 2, 3),
			modified: NewSlice(0, 1, 2, 3),
			item:     0,
			pos:      0},
		{description: "push item in slice with item in middle position",
			original: NewSlice(1, 3),
			modified: NewSlice(1, 2, 3),
			item:     2,
			pos:      1},
		{description: "push item in slice with item in last position",
			original: NewSlice(1, 2),
			modified: NewSlice(1, 2, 3),
			item:     3,
			pos:      2},
	}

	for _, tt := range useCases {
		tt.original.PushAt(tt.item, tt.pos)
		if !reflect.DeepEqual(tt.original.inner, tt.modified.inner) {
			t.Errorf("test: %s want %v got %v", tt.description, tt.modified, tt.original)
		}
	}
}

func TestSlice_Push(t *testing.T) {
	useCases := []struct {
		description string
		original    *Slice[int]
		modified    *Slice[int]
		item        int
	}{
		{description: "push item in empty slice",
			original: NewSlice[int](),
			modified: NewSlice(1),
			item:     1},
		{description: "push item in slice with item in last position",
			original: NewSlice(1, 2),
			modified: NewSlice(1, 2, 3),
			item:     3},
	}

	for _, tt := range useCases {
		tt.original.Push(tt.item)
		if !reflect.DeepEqual(tt.original.inner, tt.modified.inner) {
			t.Errorf("test: %s want %v got %v", tt.description, tt.modified, tt.original)
		}
	}
}

func TestSlice_PushBack(t *testing.T) {
	useCases := []struct {
		description string
		original    *Slice[int]
		modified    *Slice[int]
		item        int
	}{
		{description: "push item in empty slice",
			original: NewSlice[int](),
			modified: NewSlice(1),
			item:     1},
		{description: "push item in slice with item in last position",
			original: NewSlice(1, 2),
			modified: NewSlice(1, 2, 3),
			item:     3},
	}

	for _, tt := range useCases {
		tt.original.PushBack(tt.item)
		if !reflect.DeepEqual(tt.original.inner, tt.modified.inner) {
			t.Errorf("test: %s want %v got %v", tt.description, tt.modified, tt.original)
		}
	}
}

func TestSlice_PushFront(t *testing.T) {
	useCases := []struct {
		description string
		original    *Slice[int]
		modified    *Slice[int]
		item        int
	}{
		{description: "push item in empty slice",
			original: NewSlice[int](),
			modified: NewSlice(1),
			item:     1},
		{description: "push item in slice with item in first position",
			original: NewSlice(2, 3),
			modified: NewSlice(1, 2, 3),
			item:     1},
	}

	for _, tt := range useCases {
		tt.original.PushFront(tt.item)
		if !reflect.DeepEqual(tt.original.inner, tt.modified.inner) {
			t.Errorf("test: %s want %v got %v", tt.description, tt.modified, tt.original)
		}
	}
}

func TestSlice_Delete(t *testing.T) {
	useCases := []struct {
		description string
		original    *Slice[int]
		modified    *Slice[int]
		item        int
		err         error
	}{
		{description: "delete item in empty slice",
			original: NewSlice[int](),
			modified: NewSlice[int](),
			item:     0,
			err:      ErrItemNotFound{0}},
		{description: "delete item in slice with item in first position",
			original: NewSlice(1, 2, 3),
			modified: NewSlice(2, 3),
			item:     1},
		{description: "delete item in slice with item in middle position",
			original: NewSlice(1, 2, 3),
			modified: NewSlice(1, 3),
			item:     2},
		{description: "delete item in slice with item in last position",
			original: NewSlice(1, 2, 3),
			modified: NewSlice(1, 2),
			item:     3},
	}

	for _, tt := range useCases {
		err := tt.original.Delete(tt.item)
		if !reflect.DeepEqual(tt.original.inner, tt.modified.inner) || err != tt.err {
			t.Errorf("test: %s want %v got %v", tt.description, tt.modified, tt.original)
		}
	}
}

func TestSlice_DeleteAt(t *testing.T) {
	useCases := []struct {
		description string
		original    *Slice[int]
		modified    *Slice[int]
		pos         int
		err         error
	}{
		{description: "delete item in empty slice",
			original: NewSlice[int](),
			modified: NewSlice[int](),
			pos:      0,
			err:      ErrEmptyCollection},
		{description: "delete item in slice with item in first position",
			original: NewSlice(1, 2, 3),
			modified: NewSlice(2, 3),
			pos:      0},
		{description: "delete item in slice with item in middle position",
			original: NewSlice(1, 2, 3),
			modified: NewSlice(1, 3),
			pos:      1},
		{description: "delete item in slice with item in last position",
			original: NewSlice(1, 2, 3),
			modified: NewSlice(1, 2),
			pos:      2},
	}

	for _, tt := range useCases {
		err := tt.original.DeleteAt(tt.pos)
		if !reflect.DeepEqual(tt.original.inner, tt.modified.inner) || err != tt.err {
			t.Errorf("test: %s want %v got %v", tt.description, tt.modified, tt.original)
		}
	}
}

func TestSlice_Contains(t *testing.T) {
	useCases := []struct {
		description string
		slice       *Slice[int]
		item        int
		want        bool
	}{
		{description: "empty slice contain item", slice: NewSlice[int](), item: 0, want: false},
		{description: "slice with items search item no found", slice: NewSlice[int](1, 2, 3), item: 0, want: false},
		{description: "slice with items search item found", slice: NewSlice[int](1, 2, 3), item: 3, want: true},
	}

	for _, tt := range useCases {
		result := tt.slice.Contains(tt.item)
		if result != tt.want {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestSlice_Index(t *testing.T) {
	useCases := []struct {
		description string
		slice       *Slice[int]
		item        int
		pos         int
		err         error
	}{
		{description: "empty slice contain item", slice: NewSlice[int](), item: 0, pos: 0, err: ErrItemNotFound{0}},
		{description: "slice with items search item no found", slice: NewSlice[int](1, 2, 3), item: 0, pos: 0, err: ErrItemNotFound{0}},
		{description: "slice with items search item found", slice: NewSlice[int](1, 2, 3), item: 3, pos: 2, err: nil},
	}

	for _, tt := range useCases {
		pos, err := tt.slice.Index(tt.item)
		if pos != tt.pos || err != tt.err {
			t.Errorf("test: %s want {%v, %v} got {%v, %v}", tt.description, tt.pos, tt.err, pos, err)
		}
	}
}
