package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
)

func TestRemoveFromSlice(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in            []int
		out           []int
		indexToRemove int
	}{
		"normal operation": {
			in:            []int{0, 1, 2, 3},
			out:           []int{0, 1, 3},
			indexToRemove: 2,
		},

		"does not panic when out of bounds": {
			in:            []int{0, 1, 2},
			out:           []int{0, 1, 2},
			indexToRemove: 4,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := utils.RemoveFromSlice(testCase.in, testCase.indexToRemove)
			assert.Equal(t, len(testCase.out), len(result))
			for i, v := range result {
				assert.Equal(t, testCase.out[i], v)
			}
		})
	}
}
