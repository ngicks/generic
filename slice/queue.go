package slice

// FIFO queue.
type Queue[T any] []T

// Len returns length of underlying slice.
func (q *Queue[T]) Len() int {
	return len(*q)
}

// Push adds an element to tail of underlying slice.
func (q *Queue[T]) Push(v T) {
	pushBack((*[]T)(q), v)
}

// Pop removes an element from head of underlying slice, and then returns removed value.
// If slice is empty, returns zero of T and false.
func (q *Queue[T]) Pop() (v T, popped bool) {
	return popFront((*[]T)(q))
}

// Clone copies inner slice.
func (q *Queue[T]) Clone() Queue[T] {
	return Clone(*q)
}

func (q *Queue[T]) Prepend(elements ...T) {
	*q = Prepend(*q, elements...)
}
