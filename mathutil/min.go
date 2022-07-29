package mathutil

// Min finds the smallest value of all given values
func Min[T Number](a T, rest ...T) T {
	result := a

	for _, v := range rest {
		if v < result {
			result = v
		}
	}

	return result
}
