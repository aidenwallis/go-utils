package mathutil

// Clamp returns a value restricted between lo and hi.
func Clamp[T Number](v, lo, hi T) T {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
