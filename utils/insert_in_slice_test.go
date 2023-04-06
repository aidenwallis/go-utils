package utils_test

import (
	"log"
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
)

func TestInsertInSlice(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in         []int
		out        []int
		valueToAdd int
		index      int
	}{
		"normal operation": {
			in:         []int{0, 1, 2, 3, 4},
			out:        []int{0, 1, 2000, 2, 3, 4},
			valueToAdd: 2000,
			index:      2,
		},

		"out of bounds": {
			in:         []int{0, 1, 2, 3},
			out:        []int{0, 1, 2, 3, 2000},
			valueToAdd: 2000,
			index:      2000,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := utils.InsertInSlice(testCase.in, testCase.valueToAdd, testCase.index)
			log.Println(result)

			assert.Equal(t, len(testCase.out), len(result))
			for i, v := range result {
				assert.Equal(t, testCase.out[i], v)
			}
		})
	}
}
