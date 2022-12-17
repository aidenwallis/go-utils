package utils

// SliceFilter calls testFunc on every item given in input, and returns the first item that returns true
func SliceFind[T any](input []T, testFunc func(T) bool) (item T, ok bool) {
	for _, item := range input {
		if testFunc(item) {
			return item, true
		}
	}

	// else return zero value
	var zero T
	return zero, false
}
