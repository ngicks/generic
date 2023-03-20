package set

type Set[T comparable] struct {
	inner map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{
		inner: map[T]struct{}{},
	}
}

func (s *Set[T]) Len() int {
	return len(s.inner)
}

func (s *Set[T]) Add(v T) {
	s.inner[v] = struct{}{}
}

func (s *Set[T]) Clear() {
	s.inner = make(map[T]struct{})
}

func (s *Set[T]) Delete(v T) (deleted bool) {
	_, deleted = s.inner[v]
	if deleted {
		delete(s.inner, v)
	}
	return
}

// Order is undefined.
func (s *Set[T]) ForEach(f func(v T, idx int)) {
	var idx int
	for v := range s.inner {
		f(v, idx)
		idx++
	}
}

func (s *Set[T]) Has(v T) (has bool) {
	_, has = s.inner[v]
	return
}

func (s *Set[T]) Values() []T {
	sl := make([]T, 0)
	s.ForEach(func(v T, _ int) {
		sl = append(sl, v)
	})
	return sl
}
