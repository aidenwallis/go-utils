package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
)

func TestSliceFilter(t *testing.T) {
	t.Parallel()

	input := []int{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
	}

	output := []int{
		0,
		2,
		4,
		6,
	}

	v := utils.SliceFilter(input, func(i int) bool {
		return i%2 == 0
	})

	for i, vv := range v {
		assert.Equal(t, output[i], vv)
	}
}
