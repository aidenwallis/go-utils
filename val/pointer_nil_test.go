package val_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/val"
	"github.com/stretchr/testify/assert"
)

func TestPointerNil(t *testing.T) {
	t.Parallel()
	assert.Nil(t, val.PointerNil(""))
	assert.EqualValues(t, "some string", val.PointerValue(val.PointerNil("some string")))
}
