package utils

import "log"

// AllEqual returns true if all items passed in to the function are of equal value.
func AllEqual[T comparable](in ...T) bool {
	if len(in) <= 1 {
		return true
	}

	for i := 1; i < len(in); i++ {
		prev := in[i-1]
		curr := in[i]

		log.Println(i, prev, curr)
		if prev != curr {
			return false
		}
	}

	return true
}
