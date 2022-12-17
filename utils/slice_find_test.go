package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
)

func TestSliceFind(t *testing.T) {
	t.Parallel()

	item, ok := utils.SliceFind([]string{"a", "b", "c"}, func(s string) bool {
		return s == "b"
	})
	assert.Equal(t, "b", item)
	assert.True(t, ok)

	item, ok = utils.SliceFind([]string{"a", "b", "c"}, func(s string) bool {
		return s == "d"
	})
	assert.Equal(t, "", item)
	assert.False(t, ok)
}
