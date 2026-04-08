package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
)

func TestAllEqual(t *testing.T) {
	assert.True(t, utils.AllEqual(1, 1, 1))
	assert.False(t, utils.AllEqual(1, 2, 3))
}
