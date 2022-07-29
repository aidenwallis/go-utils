package mathutil

// Max finds the largest value of all given values
func Max[T Number](a T, rest ...T) T {
	result := a

	for _, v := range rest {
		if v > result {
			result = v
		}
	}

	return result
}
