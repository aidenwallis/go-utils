package mathutil_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/mathutil"
)

func TestMax(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 6, mathutil.Max(1, 2, 3, 4, 5, 6))
	assert.Equal(t, 4, mathutil.Max(4))
	assert.Equal(t, 7.0, mathutil.Max(4.0, 3.234, 1.0, 1.01, 5.0, 7.0, 2.0, 5.0, 5.0))
}
