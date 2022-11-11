package val_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/val"
)

func TestDefault(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 0, val.Default[int]())
	assert.Equal(t, nil, val.Default[*int]())
}
