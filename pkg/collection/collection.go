package collection

type Iterator[E any] interface {
	HasNext() bool
	Next() E
	NextWithIndex() (int, E)
}

type Iterable[E any] interface {
	Iterator() Iterator[E]
}

type Collection[E any] interface {
	Iterable[E]
	Empty() bool
	Size() int
	Push(item E)
	Contains(item E) bool
	Delete(item E) error
}

type List[E any] interface {
	Collection[E]
	Back() (E, error)
	Front() (E, error)
	PushBack(item E)
	PushFront(item E)
	Index(item E) (int, error)
	GetAt(pos int) (E, error)
	PushAt(item E, pos int)
	DeleteAt(pos int) error
}

type Set[E comparable] interface {
	Collection[E]
}
