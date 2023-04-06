package utils

// RemoveFromSlice removes a value from slice at a given index
func RemoveFromSlice[T any](in []T, index int) []T {
	if index >= len(in) {
		return in
	}
	return append(in[:index], in[index+1:]...)
}
