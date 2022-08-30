package val

// Default returns the default value of any type (if pointer, it is nil).
func Default[T any]() T {
	var r T
	return r
}
