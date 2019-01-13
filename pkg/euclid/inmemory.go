package euclid

import (
	"sync"

	"github.com/sagikazarmark/euclid/pkg/multilock"
)

// InMemorySequence generates an ID based on an internal state.
// Note: this generator does not persist it's state, use it only for testing.
type InMemorySequence struct {
	cur map[string]uint
	ml  multilock.Lock
	mu  sync.Mutex
}

// Generate generates the next ID of a sequence.
func (s *InMemorySequence) Next(name string) (uint, error) {
	s.ml.Lock(name)
	defer s.ml.Unlock(name)

	s.mu.Lock()
	if s.cur == nil {
		s.cur = make(map[string]uint)
	}
	s.mu.Unlock()

	s.cur[name]++

	return s.cur[name], nil
}
