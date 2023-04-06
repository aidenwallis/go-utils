package utils

// InsertInSlice inserts a value into a slice at a given position
//
// If the position gives is out of bounds then it resorts to inserting to the end of the slice
func InsertInSlice[T any](in []T, valueToAdd T, index int) []T {
	if index > len(in) {
		index = len(in)
	}
	return append(in[:index], append([]T{valueToAdd}, in[index:]...)...)
}
