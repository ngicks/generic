package maps

func Equal[K, V comparable](l, r map[K]V) bool {
	if l == nil || r == nil {
		return l == nil && r == nil
	}

	if len(l) != len(r) {
		return false
	}

	for k := range l {
		if l[k] != r[k] {
			return false
		}
	}
	return true
}
