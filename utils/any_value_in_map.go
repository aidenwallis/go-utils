package utils

// AnyValueInMap iterates through values and returns true if any has a key in the map.
func AnyValueInMap[T comparable, V any](m map[T]V, values ...T) bool {
	for _, v := range values {
		if _, ok := m[v]; ok {
			return true
		}
	}
	return false
}
