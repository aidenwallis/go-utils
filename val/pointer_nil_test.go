package val_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/val"
)

func TestPointerNil(t *testing.T) {
	t.Parallel()
	assert.Equal(t, nil, val.PointerNil(""))
	assert.Equal(t, "some string", val.PointerValue(val.PointerNil("some string")))
}
