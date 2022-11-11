package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
)

func TestSliceMap(t *testing.T) {
	t.Parallel()

	input := []struct {
		a int
	}{
		{a: 1},
		{a: 2},
		{a: 3},
	}

	output := []struct {
		a int
		b int
	}{
		{a: 1, b: 2},
		{a: 2, b: 3},
		{a: 3, b: 4},
	}

	iterator := func(v struct{ a int }) struct {
		a int
		b int
	} {
		return struct {
			a int
			b int
		}{a: v.a, b: v.a + 1}
	}

	for i, v := range utils.SliceMap(input, iterator) {
		assert.Equal(t, output[i].a, v.a)
		assert.Equal(t, output[i].b, v.b)
	}
}
