package utils

// InvertMap inverts the map keys and values
func InvertMap[K comparable, V comparable](m map[K]V) map[V]K {
	resp := make(map[V]K)
	for k, v := range m {
		resp[v] = k
	}
	return resp
}
