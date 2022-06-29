package localid

import (
	"container/list"
	"sync"
)

// Generator implements the local id generator
type Generator interface {
	// ID returns a new ID
	ID() int
	// ReturnIDs adds new IDs back to the queue.
	ReturnIDs(...int)
}

// GeneratorImpl is the struct that implements Generator
type GeneratorImpl struct {
	mu     sync.Mutex
	lastID int
	q      *list.List
}

// NewGenerator creates a new Generator instance
func NewGenerator() Generator {
	return &GeneratorImpl{
		lastID: -1, // start at 0
		q:      list.New(),
	}
}

func (g *GeneratorImpl) ID() int {
	g.mu.Lock()
	defer g.mu.Unlock()
	if id := g.q.Front(); id != nil {
		g.q.Remove(id)
		// if an ID is available in queue, use that
		return id.Value.(int)
	}

	// else, generate new DI
	g.lastID++
	return g.lastID
}

// ReturnIDs adds new IDs back to the queue.
func (g *GeneratorImpl) ReturnIDs(items ...int) {
	g.mu.Lock()
	for _, id := range items {
		g.q.PushBack(id)
	}
	g.mu.Unlock()
}
