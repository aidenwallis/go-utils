package val

// PointerNil behaves similarly to Pointer, except if the value of v is a zeroed value (eg. empty string), it returns nil instead.
func PointerNil[T comparable](v T) *T {
	var zeroValue T
	if zeroValue == v {
		// zeroed values with PointerNil should therefore return nil
		return nil
	}
	return Pointer(v)
}
