package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
	"github.com/aidenwallis/go-utils/val"
)

func TestValueAtIndex(t *testing.T) {
	t.Parallel()

	// handles normal slices
	{
		v := []int{0, 1, 2}

		{
			r, ok := utils.ValueAtIndex(v, 0)
			assert.Equal(t, 0, r)
			assert.True(t, ok)
		}

		{
			r, ok := utils.ValueAtIndex(v, 1)
			assert.Equal(t, 1, r)
			assert.True(t, ok)
		}

		{
			r, ok := utils.ValueAtIndex(v, 2)
			assert.Equal(t, 2, r)
			assert.True(t, ok)
		}

		{
			r, ok := utils.ValueAtIndex(v, 3)
			assert.Equal(t, 0, r)
			assert.False(t, ok)
		}
	}

	// handles pointer slices
	{
		v := []*int{val.Pointer(0)}

		{
			r, ok := utils.ValueAtIndex(v, 0)
			assert.Equal(t, 0, *r)
			assert.True(t, ok)
		}

		{
			r, ok := utils.ValueAtIndex(v, 1)
			assert.Equal(t, nil, r)
			assert.False(t, ok)
		}
	}
}
