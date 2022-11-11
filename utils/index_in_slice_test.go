package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
)

func TestIndexInSlice(t *testing.T) {
	t.Parallel()
	v := []int{0, 1, 2, 3}
	assert.False(t, utils.IndexInSlice(v, -1))
	assert.True(t, utils.IndexInSlice(v, 0))
	assert.True(t, utils.IndexInSlice(v, 1))
	assert.True(t, utils.IndexInSlice(v, 2))
	assert.True(t, utils.IndexInSlice(v, 3))
	assert.False(t, utils.IndexInSlice(v, 4))
}
