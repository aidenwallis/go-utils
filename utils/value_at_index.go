package utils

import "github.com/aidenwallis/go-utils/val"

// ValueAtIndex returns the value at a given index in a slice
func ValueAtIndex[T any](v []T, index int) (T, bool) {
	if IndexInSlice(v, index) {
		return v[index], true
	}
	return val.Default[T](), false
}
