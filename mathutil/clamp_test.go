package mathutil_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/mathutil"
)

func TestClamp(t *testing.T) {
	assert.Equal(t, 20, mathutil.Clamp(1, 20, 100))
	assert.Equal(t, 100, mathutil.Clamp(101, 20, 100))
	assert.Equal(t, 35, mathutil.Clamp(35, 20, 100))
}
