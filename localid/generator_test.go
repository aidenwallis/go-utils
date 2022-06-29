package localid_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/localid"
	"github.com/stretchr/testify/assert"
)

func TestLocalID(t *testing.T) {
	t.Parallel()

	g := localid.NewGenerator()

	assert.Equal(t, 0, g.ID())
	assert.Equal(t, 1, g.ID())
	assert.Equal(t, 2, g.ID())
	assert.Equal(t, 3, g.ID())
	assert.Equal(t, 4, g.ID())

	g.ReturnIDs(5, 6, 7, 8, 9)

	assert.Equal(t, 5, g.ID())
	assert.Equal(t, 6, g.ID())
	assert.Equal(t, 7, g.ID())
	assert.Equal(t, 8, g.ID())
	assert.Equal(t, 9, g.ID())
}
