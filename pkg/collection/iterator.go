package collection

type Iterator[E any] interface {
	HasNext() bool
	Next() E
	NextWithIndex() (int, E)
}

type Iterable[E any] interface {
	Iterator() Iterator[E]
}
