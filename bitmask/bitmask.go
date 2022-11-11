package bitmask

import "golang.org/x/exp/constraints"

// Bit is a generic type that exports the supported values that can be used in the bitmask.
// Any form of integer is currently supported.
type Bit interface {
	constraints.Integer
}

// Add adds a given bit to the sum.
func Add[T Bit](sum, bit T) T {
	sum |= bit
	return sum
}

// Remove removes bit from the sum.
func Remove[T Bit](sum, bit T) T {
	sum &= ^bit
	return sum
}

// Has checks whether a given bit exists in sum.
func Has[T Bit](sum, bit T) bool {
	return (sum & bit) == bit
}
