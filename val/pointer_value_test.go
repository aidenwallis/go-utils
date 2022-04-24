package val_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/val"
	"github.com/stretchr/testify/assert"
)

func TestPointerValue(t *testing.T) {
	t.Parallel()

	value := "some string"
	assert.Equal(t, value, val.PointerValue(val.Pointer("some string")))

	var nilString *string
	assert.Equal(t, "", val.PointerValue(nilString))
}
