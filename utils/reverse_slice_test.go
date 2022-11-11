package utils_test

import (
	"encoding/json"
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
)

func TestReverseSlice(t *testing.T) {
	t.Parallel()

	mustJSON := func(v interface{}) string {
		bs, err := json.Marshal(v)
		if err != nil {
			t.Error("failed to marshal json")
		}
		return string(bs)
	}

	v := []int{1, 2, 3, 4, 5}
	assert.Equal(t, mustJSON([]int{1, 2, 3, 4, 5}), mustJSON(v))
	utils.ReverseSlice(v)
	assert.Equal(t, mustJSON([]int{5, 4, 3, 2, 1}), mustJSON(v))
}
