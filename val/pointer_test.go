package val_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/val"
)

func TestPointer(t *testing.T) {
	t.Parallel()
	value := "string"
	assert.Equal(t, value, val.PointerValue(val.Pointer(value)))
}
