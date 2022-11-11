package bitmask_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/bitmask"
	"github.com/aidenwallis/go-utils/internal/assert"
)

func TestBits(t *testing.T) {
	t.Parallel()

	const (
		a = 1 << 0
		b = 1 << 1
		c = 1 << 2
	)

	assert.Equal(t, a|b, bitmask.Add(a, b))
	assert.Equal(t, a, bitmask.Remove(a|b, b))
	assert.True(t, bitmask.Has(a|b, b))
	assert.False(t, bitmask.Has(a|b, c))
}
