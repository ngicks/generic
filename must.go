package generic

// Must unwraps error possibility from a set of return value from a function.
// Must panics if err is non nil error.
// The main use case is initialization of package-top level variables.
//
//	func initializeSomeValue() (int, error) {
//		return 0, errors.New("foo")
//	}
//
//	var SomeValue = Must(initializeSomeValue()) // this line panics, meaning incorrect implementation.
//
// This example omits use of the init() function.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// Must3 is same as Must but takes 3 args.
//
// It panics if err is non nil.
func Must3[T any, U any](val1 T, val2 U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return val1, val2
}
