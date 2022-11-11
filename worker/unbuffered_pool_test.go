package worker_test

import (
	"sync"
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/worker"
)

func TestUnbufferedPool(t *testing.T) {
	t.Parallel()

	p := worker.NewUnbufferedPool()

	var (
		m sync.Mutex
		v int
	)

	for i := 0; i < 10; i++ {
		p.Do(func() {
			m.Lock()
			v++
			m.Unlock()
		})
	}

	p.Wait()

	assert.Equal(t, 10, v)
}
