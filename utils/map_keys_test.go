package utils_test

import (
	"sort"
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
)

func TestMapKeys(t *testing.T) {
	t.Parallel()

	out := []int{1, 2, 3}
	in := map[int]struct{}{
		1: {},
		2: {},
		3: {},
	}

	v := utils.MapKeys(in)
	sort.SliceStable(v, func(i, j int) bool {
		return v[i] < v[j]
	})

	assert.Equal(t, 3, len(v))
	for i := range out {
		assert.Equal(t, out[i], v[i])
	}
}
