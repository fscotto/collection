package collection

type Iterator[E any] interface {
	HasNext() bool
	Next() E
}
