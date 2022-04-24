package utils

// UniqueSlice is a helper function that iterates through your slice and returns a copy with only the unique values.
func UniqueSlice[T comparable](input []T) []T {
	resp := make([]T, 0, len(input))
	valuesMap := make(map[T]struct{}, len(resp))
	for _, item := range input {
		if _, ok := valuesMap[item]; !ok {
			valuesMap[item] = struct{}{}
			resp = append(resp, item)
		}
	}
	return resp
}
