package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/utils"
	"github.com/stretchr/testify/assert"
)

func TestReverseSlice(t *testing.T) {
	t.Parallel()
	v := []int{1, 2, 3, 4, 5}
	assert.EqualValues(t, []int{1, 2, 3, 4, 5}, v)
	utils.ReverseSlice(v)
	assert.EqualValues(t, []int{5, 4, 3, 2, 1}, v)
}
