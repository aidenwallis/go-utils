package factory_test

import (
	"testing"

	"github.com/aidenwallis/go-utils/factory"
	"github.com/aidenwallis/go-utils/internal/assert"
)

func TestChainableModifier(t *testing.T) {
	t.Parallel()

	type testValue struct {
		Foo bool
		Bar bool
	}

	chain := factory.ChainableModifier(func() *testValue {
		return &testValue{
			Foo: true,
			Bar: false,
		}
	})

	out := chain(func(v *testValue) {
		v.Foo = false
	})
	assert.False(t, out.Foo)
	assert.False(t, out.Bar)
}
