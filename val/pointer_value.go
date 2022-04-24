package val

// PointerValue returns the shallow deferenced pointer of a given value, or the zeroed value if nil.
func PointerValue[T any](v *T) T {
	if v == nil {
		var zeroValue T
		return zeroValue
	}
	return *v
}
