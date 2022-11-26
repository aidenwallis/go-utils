package utils_test

import (
	"sort"
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
)

func TestMapValues(t *testing.T) {
	t.Parallel()

	out := []int{1, 2, 3}
	in := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	v := utils.MapValues(in)
	sort.SliceStable(v, func(i, j int) bool {
		return v[i] < v[j]
	})

	assert.Equal(t, 3, len(v))
	for i := range out {
		assert.Equal(t, out[i], v[i])
	}
}
