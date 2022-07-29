package mathutil

import "golang.org/x/exp/constraints"

// Number is the type that supports any form of integer or float
type Number interface {
	float32 | float64 | constraints.Integer
}
