package helper

// Escape escapes v to a pointer of v.
//
// It is useful when setting built-in type T (e.g. string, int) to struct fields of *T.
func Escape[T any](v T) *T {
	return &v
}
