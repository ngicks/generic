package slice

// LIFO stack
type Stack[T any] []T

func (s *Stack[T]) Len() int {
	return len(*s)
}

// Push adds an element to tail of underlying slice.
func (s *Stack[T]) Push(v T) {
	pushBack((*[]T)(s), v)
}

// Pop removes an element from tail of underlying slice, and then returns removed value.
// If slice is empty, returns zero of T and false.
func (s *Stack[T]) Pop() (v T, popped bool) {
	return popBack((*[]T)(s))
}

func (s *Stack[T]) Clone() Stack[T] {
	return Clone(*s)
}

func (s *Stack[T]) Append(elements ...T) {
	*s = append(*s, elements...)
}
