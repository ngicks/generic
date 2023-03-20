package slice

// Doubly ended queue.
// This is intentionally a simple slice wrapper,
// so that it can be converted into a slice with no cost.
//
// The complexity of pushing front is typically O(N) where N is length of slice.
// If this cost needs to be avoided, use deque implementations backed by list or ring buffer.
// (for example, github.com/gammazero/deque.)
type Deque[T any] []T

// Len returns length of underlying slice.
func (d *Deque[T]) Len() int {
	return len(*d)
}

// Push is an alias for PushBack.
func (d *Deque[T]) Push(v T) {
	d.PushBack(v)
}

// Pop is an alias for PopBack.
func (d *Deque[T]) Pop() (v T, popped bool) {
	return d.PopBack()
}

// PushBack adds an element to tail of underlying slice.
func (d *Deque[T]) PushBack(v T) {
	pushBack((*[]T)(d), v)
}

// PopBack removes an element from tail of underlying slice, and then returns removed value.
// If slice is empty, returns zero of T and false.
func (d *Deque[T]) PopBack() (v T, popped bool) {
	return popBack((*[]T)(d))
}

// PushFront adds an element to head of underlying slice.
func (d *Deque[T]) PushFront(v T) {
	pushFront((*[]T)(d), v)
}

// PopFront removes an element from head of underlying slice, and then returns removed value.
// If slice is empty, returns zero of T and false.
func (d *Deque[T]) PopFront() (v T, popped bool) {
	return popFront((*[]T)(d))
}

func (d *Deque[T]) Get(index uint) (v T, ok bool) {
	return Get(*d, index)
}

// Clone copies inner slice.
func (d *Deque[T]) Clone() Deque[T] {
	return Clone(*d)
}

func (d *Deque[T]) Insert(index uint, ele T) {
	*d = Insert(*d, index, ele)
}

func (d *Deque[T]) Remove(index uint) T {
	removed := (*d)[index]
	*d = Remove(*d, index)
	return removed
}

func (d *Deque[T]) Append(elements ...T) {
	*d = append(*d, elements...)
}

func (d *Deque[T]) Prepend(elements ...T) {
	*d = Prepend(*d, elements...)
}
