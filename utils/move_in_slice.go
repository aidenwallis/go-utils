package utils

// MoveInSlice moves a value within the slice to a specified new position
//
// If either the old or new position is out of bounds, the existing slice contents is returned.
func MoveInSlice[T any](in []T, oldPosition, newPosition int) []T {
	if oldPosition >= len(in) {
		return in
	}
	value := in[oldPosition]
	return InsertInSlice(RemoveFromSlice(in, oldPosition), value, newPosition)
}
