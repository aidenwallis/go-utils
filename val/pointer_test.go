package val_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/val"
	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {
	t.Parallel()

	value := "string"
	assert.EqualValues(t, &value, val.Pointer(value))
}
