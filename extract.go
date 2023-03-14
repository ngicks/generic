package helper

func ExtractError[T any](v T, err error) error {
	return err
}

func ExtractError3[T, U any](v1 T, v2 U, err error) error {
	return err
}
