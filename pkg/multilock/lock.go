package multilock

import (
	"sync"
)

// Lock can acquire locks for a specific holder.
type Lock struct {
	locks map[interface{}]*sync.Mutex
	mu    sync.Mutex
}

// Lock acquires a lock for a specific holder.
func (l *Lock) Lock(holder interface{}) {
	mu := l.getLocker(holder)

	mu.Lock()
}

// Unlock releases a lock of a holder.
func (l *Lock) Unlock(holder interface{}) {
	mu := l.getLocker(holder)

	mu.Unlock()
}

func (l *Lock) getLocker(holder interface{}) *sync.Mutex {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.locks == nil {
		l.locks = make(map[interface{}]*sync.Mutex)
	}

	mu, ok := l.locks[holder]
	if !ok {
		mu = &sync.Mutex{}
		l.locks[holder] = mu
	}

	return mu
}
