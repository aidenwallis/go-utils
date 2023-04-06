package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
)

func TestMoveInSlice(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in          []int
		out         []int
		oldPosition int
		newPosition int
	}{
		"normal operation": {
			in:          []int{0, 1, 2, 3, 4},
			out:         []int{0, 3, 1, 2, 4},
			oldPosition: 3,
			newPosition: 1,
		},

		"out of bounds in old position": {
			in:          []int{0, 1, 2, 3, 4},
			out:         []int{0, 1, 2, 3, 4},
			oldPosition: 5,
			newPosition: 0,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := utils.MoveInSlice(testCase.in, testCase.oldPosition, testCase.newPosition)

			assert.Equal(t, len(testCase.out), len(result))
			for i, v := range result {
				assert.Equal(t, testCase.out[i], v)
			}
		})
	}
}
