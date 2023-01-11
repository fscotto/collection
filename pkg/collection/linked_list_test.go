package collection

import (
	"testing"
)

func TestLinkedList_Empty(t *testing.T) {
	useCases := []struct {
		description string
		list        *LinkedList[int]
		want        bool
	}{
		{description: "no items in list", list: NewLinkedList[int](), want: true},
		{description: "one item in list", list: NewLinkedList(1), want: false},
		{description: "more items in list", list: NewLinkedList(1, 2, 3), want: false},
	}

	for _, tt := range useCases {
		result := tt.list.Empty()
		if result != tt.want {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestLinkedList_Size(t *testing.T) {
	useCases := []struct {
		description string
		list        *LinkedList[int]
		want        int
	}{
		{description: "no items in list", list: NewLinkedList[int](), want: 0},
		{description: "one item in list", list: NewLinkedList(1), want: 1},
		{description: "more items in list", list: NewLinkedList(1, 2, 3), want: 3},
	}

	for _, tt := range useCases {
		result := tt.list.Size()
		if result != tt.want {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestLinkedList_GetAt(t *testing.T) {
	useCases := []struct {
		description string
		list        *LinkedList[int]
		pos         int
		err         error
		want        int
	}{
		{description: "get first item", list: NewLinkedList(1, 2, 3), pos: 0, want: 1, err: nil},
		{description: "get middle item", list: NewLinkedList(1, 2, 3), pos: 1, want: 2, err: nil},
		{description: "get last item", list: NewLinkedList(1, 2, 3), pos: 2, want: 3, err: nil},
		{description: "get last item", list: NewLinkedList(1, 2, 3), pos: 2, want: 3, err: nil},
		{description: "no items in list return zero value of the type", list: NewLinkedList[int](), pos: 0, want: 0, err: ErrEmptyCollection},
		{description: "get item with negative position", list: NewLinkedList(1, 2, 3), pos: -1, want: 0, err: ErrPositionNegative},
		{description: "get item with index out of bound", list: NewLinkedList(1, 2, 3), pos: 4, want: 0, err: ErrIndexOutOfBound{4, 3}},
	}

	for _, tt := range useCases {
		result, err := tt.list.GetAt(tt.pos)
		if result != tt.want || err != tt.err {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestLinkedList_Back(t *testing.T) {
	useCases := []struct {
		description string
		list        *LinkedList[int]
		err         error
		want        int
	}{
		{description: "no items in list return zero value of the type", list: NewLinkedList[int](), want: 0, err: ErrEmptyCollection},
		{description: "singleton list return first element", list: NewLinkedList(1), want: 1},
		{description: "get last item", list: NewLinkedList(1, 2, 3), want: 3},
	}

	for _, tt := range useCases {
		result, err := tt.list.Back()
		if result != tt.want || err != tt.err {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestLinkedList_Front(t *testing.T) {
	useCases := []struct {
		description string
		list        *LinkedList[int]
		err         error
		want        int
	}{
		{description: "no items in list return zero value of the type", list: NewLinkedList[int](), want: 0, err: ErrEmptyCollection},
		{description: "singleton list return first element", list: NewLinkedList(1), want: 1},
		{description: "get first item", list: NewLinkedList(1, 2, 3), want: 1},
	}

	for _, tt := range useCases {
		result, err := tt.list.Front()
		if result != tt.want || err != tt.err {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestLinkedList_Push(t *testing.T) {
	useCases := []struct {
		description string
		original    *LinkedList[int]
		modified    *LinkedList[int]
		item        int
	}{
		{description: "push item in empty slice",
			original: NewLinkedList[int](),
			modified: NewLinkedList(1),
			item:     1},
		{description: "push item in slice with item in last position",
			original: NewLinkedList(1, 2),
			modified: NewLinkedList(1, 2, 3),
			item:     3},
	}

	for _, tt := range useCases {
		tt.original.Push(tt.item)
		if !compareLists(tt.original, tt.modified) {
			t.Errorf("test: %s want %v got %v", tt.description, tt.modified, tt.original)
		}
	}
}

func TestLinkedList_PushAt(t *testing.T) {
	useCases := []struct {
		description string
		original    *LinkedList[int]
		modified    *LinkedList[int]
		pos         int
		item        int
	}{
		{description: "add item in first position in empty list",
			original: NewLinkedList[int](),
			modified: NewLinkedList(1),
			pos:      0,
			item:     1},
		{description: "add item in first position in full list",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(0, 1, 2, 3),
			pos:      0,
			item:     0},
		{description: "add item in middle position in full list",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(1, 2, 0, 3),
			pos:      2,
			item:     0},
		{description: "add item in last position in full list",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(1, 2, 3, 0),
			pos:      3,
			item:     0},
	}

	for _, tt := range useCases {
		tt.original.PushAt(tt.item, tt.pos)
		if !compareLists(tt.original, tt.modified) {
			t.Errorf("test: %s want %v got %v", tt.description, tt.modified, tt.original)
		}
	}
}

func TestLinkedList_PushBack(t *testing.T) {
	useCases := []struct {
		description string
		original    *LinkedList[int]
		modified    *LinkedList[int]
		pos         int
		item        int
	}{
		{description: "add item in first position in empty list",
			original: NewLinkedList[int](),
			modified: NewLinkedList(1),
			item:     1},
		{description: "add item in singleton list",
			original: NewLinkedList(1),
			modified: NewLinkedList(1, 0),
			item:     0},
		{description: "add item in full list",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(1, 2, 3, 0),
			item:     0},
	}

	for _, tt := range useCases {
		tt.original.PushBack(tt.item)
		if !compareLists(tt.original, tt.modified) {
			t.Errorf("test: %s want %v got %v", tt.description, tt.modified, tt.original)
		}
	}
}

func TestLinkedList_PushFront(t *testing.T) {
	useCases := []struct {
		description string
		original    *LinkedList[int]
		modified    *LinkedList[int]
		pos         int
		item        int
	}{
		{description: "add item in first position in empty list",
			original: NewLinkedList[int](),
			modified: NewLinkedList(1),
			item:     1},
		{description: "add item in singleton list",
			original: NewLinkedList(1),
			modified: NewLinkedList(0, 1),
			item:     0},
		{description: "add item in full list",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(0, 1, 2, 3),
			item:     0},
	}

	for _, tt := range useCases {
		tt.original.PushFront(tt.item)
		if !compareLists(tt.original, tt.modified) {
			t.Errorf("test: %s want %v got %v", tt.description, tt.modified, tt.original)
		}
	}
}

func TestLinkedList_Set(t *testing.T) {
	useCases := []struct {
		description string
		original    *LinkedList[int]
		modified    *LinkedList[int]
		pos         int
		item        int
		err         error
	}{
		{description: "change item in first position in singleton list",
			original: NewLinkedList[int](0),
			modified: NewLinkedList(1),
			pos:      0,
			item:     1},
		{description: "change item in first position in full list",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(0, 2, 3),
			pos:      0,
			item:     0},
		{description: "change item in middle position in full list",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(1, 0, 3),
			pos:      1,
			item:     0},
		{description: "add item in last position in full list",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(1, 2, 0),
			pos:      2,
			item:     0},
		{description: "no items in list return ErrEmptyCollection",
			original: NewLinkedList[int](),
			modified: NewLinkedList[int](),
			pos:      0,
			item:     0,
			err:      ErrEmptyCollection},
		{description: "get item with negative position",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(1, 2, 3),
			pos:      -1,
			item:     0,
			err:      ErrPositionNegative},
		{description: "get item with index out of bound",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(1, 2, 3),
			pos:      4,
			item:     0,
			err:      ErrIndexOutOfBound{4, 3}},
	}

	for _, tt := range useCases {
		err := tt.original.Set(tt.item, tt.pos)
		if !compareLists(tt.original, tt.modified) || tt.err != err {
			t.Errorf("test: %s want %v got %v", tt.description, tt.modified, tt.original)
		}
	}
}

func TestLinkedList_Delete(t *testing.T) {
	useCases := []struct {
		description string
		original    *LinkedList[int]
		modified    *LinkedList[int]
		item        int
		err         error
	}{
		{description: "delete item in empty list",
			original: NewLinkedList[int](),
			modified: NewLinkedList[int](),
			item:     0,
			err:      ErrItemNotFound{0}},
		{description: "delete item in list with item in first position",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(2, 3),
			item:     1},
		{description: "delete item in list with item in middle position",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(1, 3),
			item:     2},
		{description: "delete item in list with item in last position",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(1, 2),
			item:     3},
	}

	for _, tt := range useCases {
		err := tt.original.Delete(tt.item)
		if !compareLists(tt.original, tt.modified) || err != tt.err {
			t.Errorf("test: %s want {%v, %v} got {%v, %v}", tt.description, tt.modified, tt.err, tt.original, err)
		}
	}
}

func TestLinkedList_DeleteAt(t *testing.T) {
	useCases := []struct {
		description string
		original    *LinkedList[int]
		modified    *LinkedList[int]
		pos         int
		err         error
	}{
		{description: "delete item with negative position",
			original: NewLinkedList[int](),
			modified: NewLinkedList[int](),
			pos:      -1,
			err:      ErrPositionNegative},
		{description: "delete item in empty list",
			original: NewLinkedList[int](),
			modified: NewLinkedList[int](),
			pos:      0,
			err:      ErrEmptyCollection},
		{description: "delete item in position not found",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(1, 2, 3),
			pos:      5,
			err:      ErrNodeNotFound},
		{description: "delete item in first position",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(2, 3),
			pos:      0,
			err:      nil},
		{description: "delete item in middle position",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(1, 3),
			pos:      1,
			err:      nil},
		{description: "delete item in last position",
			original: NewLinkedList(1, 2, 3),
			modified: NewLinkedList(1, 2),
			pos:      2,
			err:      nil},
	}

	for _, tt := range useCases {
		err := tt.original.DeleteAt(tt.pos)
		if !compareLists(tt.original, tt.modified) || err != tt.err {
			t.Errorf("test: %s want {%v, %v} got {%v, %v}", tt.description, tt.modified, tt.err, tt.original, err)
		}
	}
}

func TestLinkedList_Contains(t *testing.T) {
	useCases := []struct {
		description string
		list        *LinkedList[int]
		item        int
		want        bool
	}{
		{description: "empty linked list contain item", list: NewLinkedList[int](), item: 0, want: false},
		{description: "linked list with items search item no found", list: NewLinkedList[int](1, 2, 3), item: 0, want: false},
		{description: "linked list with items search item found", list: NewLinkedList[int](1, 2, 3), item: 3, want: true},
	}

	for _, tt := range useCases {
		result := tt.list.Contains(tt.item)
		if result != tt.want {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestLinkedList_Index(t *testing.T) {
	useCases := []struct {
		description string
		list        *LinkedList[int]
		item        int
		pos         int
		err         error
	}{
		{description: "empty linked list contain item", list: NewLinkedList[int](), item: 0, pos: 0, err: ErrItemNotFound{0}},
		{description: "linked list with items search item no found", list: NewLinkedList[int](1, 2, 3), item: 0, pos: 0, err: ErrItemNotFound{0}},
		{description: "linked list with items search item found", list: NewLinkedList[int](1, 2, 3), item: 3, pos: 2, err: nil},
	}

	for _, tt := range useCases {
		pos, err := tt.list.Index(tt.item)
		if pos != tt.pos || err != tt.err {
			t.Errorf("test: %s want {%v, %v} got {%v, %v}", tt.description, tt.pos, tt.err, pos, err)
		}
	}
}

func compareLists(lst1, lst2 *LinkedList[int]) bool {
	for it := lst1.Iterator(); it.HasNext(); {
		i, item1 := it.NextWithIndex()
		item2, _ := lst2.GetAt(i)
		if item1 != item2 {
			return false
		}
	}
	return true
}
