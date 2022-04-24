package utils

// SliceMap takes a slice, calls the func given for every item, then returns a new slice of the
// transformed value.
func SliceMap[T any, R any](input []T, mapFunc func(T) R) []R {
	v := make([]R, len(input))
	for i, item := range input {
		v[i] = mapFunc(item)
	}
	return v
}
