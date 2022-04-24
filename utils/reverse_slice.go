package utils

// ReverseSlice reverses all slice elements in place. Note: this will mutate your given slice.
func ReverseSlice[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
