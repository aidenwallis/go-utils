package utils

// FirstOf returns the first item in a slice
func FirstOf[T any](v []T) T {
	if len(v) == 0 {
		var zeroValue T
		return zeroValue
	}
	return v[0]
}
