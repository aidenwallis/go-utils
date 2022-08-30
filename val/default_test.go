package val_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/val"
	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 0, val.Default[int]())
	assert.Nil(t, val.Default[*int]())
}
