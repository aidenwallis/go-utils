package worker

import "sync"

// UnbufferedPool is an interface which implements BufferedPoolImpl, allowing you to mock out the pool for tests.
type UnbufferedPool interface {
	// Do enqueues a task and blocks the goroutine until it's enqueued.
	Do(func())

	// Wait blocks the goroutine until all tasks are complete.
	Wait()
}

// compile time assertion
var _ UnbufferedPool = (*UnbufferedPoolImpl)(nil)

// UnbufferedPoolImpl is an implementation that is compatible with UnbufferedPool.
//
// The main purpose of this tool is to let you ensure all goroutines have been closed before exiting your app, you can
// pass an instance of this around everywhere instead.
type UnbufferedPoolImpl struct {
	// WaitGroup tracks how many running/queued tasks there are, we expose Wait() so you can wait until all tasks are complete.
	WaitGroup sync.WaitGroup
}

// NewUnbufferedPool creates a new instance of BufferedPoolImpl
func NewUnbufferedPool() *UnbufferedPoolImpl {
	return &UnbufferedPoolImpl{}
}

// Do increments the wait group and invokes the goroutine, then decrements it.
func (w *UnbufferedPoolImpl) Do(cb func()) {
	w.WaitGroup.Add(1)
	go func(cb func()) {
		cb()
		w.WaitGroup.Done()
	}(cb)
}

// Wait blocks the goroutine until all tasks are complete.
func (w *UnbufferedPoolImpl) Wait() {
	w.WaitGroup.Wait()
}
