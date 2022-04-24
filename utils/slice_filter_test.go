package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/utils"
	"github.com/stretchr/testify/assert"
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

	// test that it only returns the even items
	assert.Equal(t, output, utils.SliceFilter(input, func(i int) bool {
		return i%2 == 0
	}))
}
