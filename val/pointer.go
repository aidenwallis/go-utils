package val

// AsPointer returns a pointer of a given value, useful if you need to generate a pointer to a string.
// For example, val.AsPointer("string")
func Pointer[T any](v T) *T {
	return &v
}
