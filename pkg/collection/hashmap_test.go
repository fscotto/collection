package collection

import "testing"

func TestHashMap_Empty(t *testing.T) {
	useCases := []struct {
		description string
		table       *HashMap[int, string]
		want        bool
	}{
		{description: "empty hash map", table: NewHashMap[int, string](), want: true},
		{description: "hash map with one entry", table: NewHashMap[int, string](NewEntry(1, "hello there")), want: false},
	}

	for _, tt := range useCases {
		result := tt.table.Empty()
		if result != tt.want {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestHashMap_Size(t *testing.T) {
	useCases := []struct {
		description string
		table       *HashMap[int, string]
		want        int
	}{
		{description: "empty hash map", table: NewHashMap[int, string](), want: 0},
		{description: "hash map with one entry", table: NewHashMap[int, string](NewEntry(1, "hello there")), want: 1},
	}

	for _, tt := range useCases {
		result := tt.table.Size()
		if result != tt.want {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestHashMap_Get(t *testing.T) {
	useCases := []struct {
		description string
		table       *HashMap[int, string]
		key         int
		want        string
	}{
		{description: "Ask value by not found key", table: NewHashMap[int, string](), key: 1, want: ""},
		{description: "Ask value by found key", table: NewHashMap[int, string](NewEntry(1, "hello there")), key: 1, want: "hello there"},
	}

	for _, tt := range useCases {
		result, _ := tt.table.Get(tt.key)
		if result != tt.want {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestHashMap_Put(t *testing.T) {
	useCases := []struct {
		description string
		table       *HashMap[int, string]
		entry       *Entry[int, string]
		want        int
	}{
		{description: "empty map put new pair",
			table: NewHashMap[int, string](),
			entry: NewEntry(1, "hello there"),
			want:  1},
		{description: "put new pair in map with elements",
			table: NewHashMap[int, string](
				NewEntry(1, "hello there"),
				NewEntry(2, "welcome general Kenobi"),
			),
			entry: NewEntry(3, "what's you name?"),
			want:  3},
	}

	for _, tt := range useCases {
		tt.table.Put(tt.entry.Key(), tt.entry.Value())
		if tt.table.Size() != tt.want {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, tt.table.Size())
		}
	}
}

func TestHashMap_ContainsKey(t *testing.T) {
	useCases := []struct {
		description string
		table       *HashMap[int, string]
		key         int
		want        bool
	}{
		{description: "Seek key not present in map",
			table: NewHashMap[int, string](),
			key:   1,
			want:  false},
		{description: "Seek key present in table",
			table: NewHashMap[int, string](
				NewEntry(1, "hello there"),
			),
			key:  1,
			want: true},
	}

	for _, tt := range useCases {
		result := tt.table.ContainsKey(tt.key)
		if result != tt.want {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestHashMap_ContainsValue(t *testing.T) {
	useCases := []struct {
		description string
		table       *HashMap[int, string]
		value       string
		want        bool
	}{
		{description: "Seek key not present in map",
			table: NewHashMap[int, string](),
			value: "hello there",
			want:  false},
		{description: "Seek key present in map",
			table: NewHashMap[int, string](
				NewEntry(1, "hello there"),
				NewEntry(2, "welcome general Kenobi"),
			),
			value: "hello there",
			want:  true},
	}

	for _, tt := range useCases {
		result := tt.table.ContainsValue(tt.value)
		if result != tt.want {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, result)
		}
	}
}

func TestHashMap_Delete(t *testing.T) {
	useCases := []struct {
		description string
		table       *HashMap[int, string]
		key         int
		want        int
		result      bool
	}{
		{description: "in empty map delete item is no-op action",
			table:  NewHashMap[int, string](),
			key:    1,
			want:   0,
			result: false},
		{description: "in empty map delete item is no-op action",
			table: NewHashMap[int, string](
				NewEntry(1, "hello there"),
			),
			key:    1,
			want:   0,
			result: true},
		{description: "with three entry delete entry with key 1",
			table: NewHashMap[int, string](
				NewEntry(1, "hello there"),
				NewEntry(2, "welcome general Kenobi"),
				NewEntry(3, "what's you name?"),
			),
			key:    1,
			want:   2,
			result: true},
	}

	for _, tt := range useCases {
		ok := tt.table.Delete(tt.key)
		if tt.table.Size() != tt.want || ok != tt.result {
			t.Errorf("test: %s want %v got %v", tt.description, tt.want, tt.table.Size())
		}
	}
}
