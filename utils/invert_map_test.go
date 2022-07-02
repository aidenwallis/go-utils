package utils_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/utils"
	"github.com/stretchr/testify/assert"
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

	assert.EqualValues(t, output, utils.InvertMap(input))
}
