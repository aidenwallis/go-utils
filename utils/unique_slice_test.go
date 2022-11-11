package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
)

func TestUniqueSlice(t *testing.T) {
	t.Parallel()

	t.Run("ints", func(t *testing.T) {
		t.Parallel()

		input := []int{1, 2, 3, 4, 5, 1, 2, 2, 6, 6}
		output := []int{1, 2, 3, 4, 5, 6}
		v := utils.UniqueSlice(input)
		for i, vv := range v {
			assert.Equal(t, output[i], vv)
		}
	})

	t.Run("strings", func(t *testing.T) {
		t.Parallel()

		input := []string{"a", "b", "A", "b", "C", "d", "E"}
		output := []string{"a", "b", "A", "C", "d", "E"}
		v := utils.UniqueSlice(input)
		for i, vv := range v {
			assert.Equal(t, output[i], vv)
		}
	})
}
