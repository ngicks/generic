package slice

func pushBack[T any](sl *[]T, v T) {
	*sl = append(*sl, v)
}

func pushFront[T any](sl *[]T, v T) {
	var zero T
	appended := append(*sl, zero)
	copy(appended[1:], appended)
	appended[0] = v
	*sl = appended
}

func popBack[T any](sl *[]T) (v T, popped bool) {
	var zero T

	if len(*sl) == 0 {
		return zero, false
	}

	v = (*sl)[len(*sl)-1]

	// avoiding memory leak.
	(*sl)[len(*sl)-1] = zero

	*sl = (*sl)[:len(*sl)-1]

	return v, true
}

func popFront[T any](sl *[]T) (v T, popped bool) {
	var zero T

	if len(*sl) == 0 {
		return zero, false
	}

	v = (*sl)[0]

	// avoiding memory leak
	(*sl)[0] = zero

	*sl = (*sl)[1:]

	return v, true
}
