package worker

import "sync"

// BufferedPool is an interface which implements BufferedPoolImpl, allowing you to mock out the pool for tests.
type BufferedPool interface {
	// Do enqueues a task and blocks the goroutine until it's enqueued.
	Do(func())

	// TryDo will attempt to enqueue the task, if the buffer is full, it will return false
	TryDo(func()) bool

	// Wait blocks the goroutine until all tasks are complete.
	Wait()

	// WaitAndClose closes the worker pool once all tasks are complete, you must not try to enqueue more tasks after calling this.
	WaitAndClose()
}

// compile time assertion
var _ BufferedPool = (*BufferedPoolImpl)(nil)

// PoolImpl is an implementation that is compatible with Pool
type BufferedPoolImpl struct {
	closeOnce sync.Once

	// Tasks is a buffer of all enqueued tasks to be ran by the workers
	Tasks chan func()

	// WaitGroup tracks how many running/queued tasks there are, we expose Wait() so you can wait until all tasks are complete.
	WaitGroup sync.WaitGroup
}

func NewBufferedPool(workers, maxQueue int) *BufferedPoolImpl {
	w := &BufferedPoolImpl{
		Tasks: make(chan func(), maxQueue),
	}
	go w.init(workers)
	return w
}

// Interface is a convenience method which gives you the same instance as an interface.
func (w *BufferedPoolImpl) Interface() BufferedPool {
	return w
}

// Do enqueues a task and blocks the goroutine until it's enqueued.
func (w *BufferedPoolImpl) Do(cb func()) {
	w.WaitGroup.Add(1)
	w.Tasks <- cb
}

func (w *BufferedPoolImpl) TryDo(cb func()) bool {
	// to prevent a race condition, we optimistically add to the wait group first
	w.WaitGroup.Add(1)

	select {
	case w.Tasks <- cb:
		return true

	default:
		// close it out early
		w.WaitGroup.Done()
		return false
	}
}

// init creates the pool of goroutines
func (w *BufferedPoolImpl) init(workers int) {
	for i := 0; i < workers; i++ {
		go func() {
			for task := range w.Tasks {
				task()

				// task is complete, decrement wait group
				w.WaitGroup.Done()
			}
		}()
	}
}

// Wait blocks the goroutine until all tasks are complete.
func (w *BufferedPoolImpl) Wait() {
	w.WaitGroup.Wait()
}

// Close closes the worker pool, you must not try to enqueue more tasks after calling this.
func (w *BufferedPoolImpl) Close() {
	w.closeOnce.Do(func() {
		close(w.Tasks)
	})
}

// WaitAndClose closes the worker pool once all tasks are complete, you must not try to enqueue more tasks after calling this.
func (w *BufferedPoolImpl) WaitAndClose() {
	w.Wait()
	w.Close()
}
