package utils

// IndexInSlice returns whether an index exists within a given slice
func IndexInSlice[T any](v []T, index int) bool {
	if index < 0 {
		// slices cannot be less than 0 in size
		return false
	}
	return len(v) > index
}
