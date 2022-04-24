package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/utils"
	"github.com/stretchr/testify/assert"
)

func TestUniqueSlice(t *testing.T) {
	t.Parallel()

	t.Run("ints", func(t *testing.T) {
		t.Parallel()

		input := []int{1, 2, 3, 4, 5, 1, 2, 2, 6, 6}
		output := []int{1, 2, 3, 4, 5, 6}
		assert.Equal(t, output, utils.UniqueSlice(input))
	})

	t.Run("strings", func(t *testing.T) {
		t.Parallel()

		input := []string{"a", "b", "A", "b", "C", "d", "E"}
		output := []string{"a", "b", "A", "C", "d", "E"}
		assert.Equal(t, output, utils.UniqueSlice(input))
	})
}
