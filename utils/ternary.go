package utils

// Ternary allows you to write ternary expressions in go, if exp is true, then ifCond is returned, else, elseCond is returned.
func Ternary[T any](exp bool, ifCond T, elseCond T) T {
	if exp {
		return ifCond
	}
	return elseCond
}
