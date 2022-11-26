package utils

// MapValues returns all values in a given map, **the order is not stable, you should sort it if you need a stable order.**
func MapValues[K comparable, V any](in map[K]V) []V {
	out := make([]V, 0, len(in))
	for _, v := range in {
		out = append(out, v)
	}
	return out
}
