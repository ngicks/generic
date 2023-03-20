package maps

func Clone[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}

	out := make(map[K]V, len(m))
	for k, v := range m {
		out[k] = v
	}

	return out
}

func CloneSelected[K comparable, V any](m map[K]V, selector func(k K, v V) bool) map[K]V {
	if m == nil {
		return nil
	}

	out := make(map[K]V)
	for k, v := range m {
		if selector(k, v) {
			out[k] = v
		}
	}

	return out
}

func CloneExcluded[K comparable, V any](m map[K]V, excluder func(k K, v V) bool) map[K]V {
	if m == nil {
		return nil
	}

	out := make(map[K]V)
	for k, v := range m {
		if !excluder(k, v) {
			out[k] = v
		}
	}

	return out
}
