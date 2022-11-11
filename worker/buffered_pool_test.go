package worker_test

import (
	"sync"
	"testing"

	"github.com/aidenwallis/go-utils/internal/assert"
	"github.com/aidenwallis/go-utils/worker"
)

func TestBufferedPool(t *testing.T) {
	t.Parallel()

	t.Run("generic use", func(t *testing.T) {
		t.Parallel()

		outputs := []int{}
		outputsMutex := sync.Mutex{}

		pool := worker.NewBufferedPool(10, 1).Interface()

		for i := 0; i < 20; i++ {
			pool.Do(func() {
				outputsMutex.Lock()
				outputs = append(outputs, 123)
				outputsMutex.Unlock()
			})
		}

		pool.WaitAndClose()

		assert.Equal(t, len(outputs), 20)
	})

	t.Run("try to write to filled queue", func(t *testing.T) {
		// this is kinda hacky, but essentially we'll block the workers using this channel until we've validated this test works.
		ch := make(chan struct{}, 1)
		pool := worker.NewBufferedPool(10, 1)
		defer pool.Close()

		// 10 workers + 1 in the queue
		for i := 0; i < 11; i++ {
			pool.Do(func() {
				// this will pause the goroutine until we validate TryDo returns false
				<-ch
			})
		}

		// all goroutines blocked + queue is full, should be false now
		assert.False(t, pool.TryDo(func() {}))
		close(ch) // closing the channel will exit all blocked workers

		// wait until the prev test is empty, then we should get true next
		pool.Wait()

		assert.True(t, pool.TryDo(func() {}))
	})
}
