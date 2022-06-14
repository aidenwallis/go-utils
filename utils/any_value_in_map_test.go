package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/utils"
	"github.com/stretchr/testify/assert"
)

func TestAnyValueInMap(t *testing.T) {
	t.Parallel()

	m := map[string]struct{}{
		"1": {},
		"2": {},
		"3": {},
	}

	assert.True(t, utils.AnyValueInMap(m, "foobar", "1"))
	assert.False(t, utils.AnyValueInMap(m, "foobar"))
}
