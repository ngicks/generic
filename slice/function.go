package slice

import (
	"reflect"
)

// This file sticks its naming convention to it of Rust:
//   - https://doc.rust-lang.org/std/iter/trait.Iterator.html
//   - https://doc.rust-lang.org/std/vec/struct.Vec.html

func Clone[T any](sl []T) []T {
	cloned := make([]T, len(sl))
	copy(cloned, sl)
	return cloned
}

func Eq[T comparable](left, right []T) bool {
	if len(left) != len(right) {
		return false
	} else if (left == nil || right == nil) && (left != nil || right != nil) {
		return false
	}

	// Comparing types which is constrained by comparable
	// may not be actually comparable in GO 1.20 or later version.
	// It might cause a runtime panic.
	var zero T
	if !reflect.TypeOf(zero).Comparable() {
		return false
	}

	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func Find[T any](sl []T, predicate func(v T) bool) (v T, found bool) {
	p := Position(sl, predicate)
	if p < 0 {
		return
	}
	return sl[p], true
}

func FindLast[T any](sl []T, predicate func(v T) bool) (v T, found bool) {
	p := PositionLast(sl, predicate)
	if p < 0 {
		return
	}
	return sl[p], true
}

func Get[T any](sl []T, index uint) (v T, ok bool) {
	if index >= uint(len(sl)) {
		return
	}
	return sl[index], true
}

func Has[T comparable](sl []T, target T) bool {
	return Position(sl, func(v T) bool { return v == target }) >= 0
}

func Insert[T any](sl []T, index uint, ele T) []T {
	var zero T
	sl = append(sl, zero)
	copy(sl[index+1:], sl[index:])
	sl[index] = ele
	return sl
}

func Remove[T any](sl []T, i uint) []T {
	sl = sl[:int(i)+copy(sl[int(i):], sl[int(i)+1:])]
	return sl
}

func RemoveUnordered[T any](sl []T, index uint) []T {
	sl[index] = sl[len(sl)-1]
	sl = sl[:len(sl)-1]
	return sl
}

func Position[T any](sl []T, predicate func(v T) bool) int {
	if len(sl) == 0 || predicate == nil {
		return -1
	}
	for idx, v := range sl {
		if predicate(v) {
			return idx
		}
	}
	return -1
}

func PositionLast[T any](sl []T, predicate func(v T) bool) int {
	if len(sl) == 0 || predicate == nil {
		return -1
	}
	for i := len(sl) - 1; i >= 0; i-- {
		if predicate(sl[i]) {
			return i
		}
	}
	return -1
}

func Prepend[T any](sl []T, elements ...T) []T {
	sl = append(sl, elements...)
	copy(sl[len(elements):], sl)
	copy(sl, elements)
	return sl
}
