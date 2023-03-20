package generic

// Assert does type assertion on v.
// v is asserted against T and *T.
// Assert returns true ok when v is either of T or *T.
// If v is nil of *T, then asserted is zero value of T.
func Assert[T any](v any) (asserted T, ok bool) {
	asserted, ok = v.(T)
	if !ok {
		p, ok := v.(*T)
		if !ok {
			return asserted, false
		}
		if p != nil {
			asserted = *p
		}
	}

	return asserted, true
}
