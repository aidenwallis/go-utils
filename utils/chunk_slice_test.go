package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
)

func TestChunkSlice(t *testing.T) {
	t.Parallel()

	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	out := utils.ChunkSlice(in, 3)

	for i, chunk := range out {
		assert.Equal(t, 3, len(chunk))
		for j, v := range chunk {
			assert.Equal(t, expected[i][j], v)
		}
	}
}
