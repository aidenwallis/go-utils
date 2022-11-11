package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
)

func TestTernary(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "a", utils.Ternary(1 == 1, "a", "b")) // nolint:staticcheck // intended
	assert.Equal(t, "b", utils.Ternary(1 == 2, "a", "b"))
}
