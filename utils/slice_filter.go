package utils

// SliceFilter calls testFunc on every item given in input, and only returns those that return true.
func SliceFilter[T any](input []T, testFunc func(T) bool) []T {
	resp := make([]T, 0, len(input))
	for _, item := range input {
		if testFunc(item) {
			resp = append(resp, item)
		}
	}
	return resp
}
