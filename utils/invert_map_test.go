package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/utils"
)

func TestInvertMap(t *testing.T) {
	t.Parallel()

	input := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	output := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	for k, v := range input {
		assert.Equal(t, v, utils.InvertMap(output)[k])
	}
}
