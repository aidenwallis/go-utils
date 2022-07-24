package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/utils"
	"github.com/stretchr/testify/assert"
)

func TestFirstOf(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "a", utils.FirstOf([]string{"a", "b", "c"}))
	assert.Equal(t, "", utils.FirstOf([]string{}))
	assert.Equal(t, "", utils.FirstOf[string](nil))
}
