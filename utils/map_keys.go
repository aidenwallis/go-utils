package utils

// MapKeys returns all keys in a given map, **the order will not be stable, you should sort these values if needed.**
func MapKeys[K comparable, V any](v map[K]V) []K {
	resp := make([]K, 0, len(v))
	for k := range v {
		resp = append(resp, k)
	}
	return resp
}
