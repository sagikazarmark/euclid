package multilock

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLock(t *testing.T) {
	var mu Lock
	var wg sync.WaitGroup

	wg.Add(2)

	var counter []int

	go func() {
		defer wg.Done()

		for i := 2; i < 10; i += 2 {
			mu.Lock("counter")
			counter = append(counter, i)
			mu.Unlock("counter")
			time.Sleep(50 * time.Microsecond)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 1; i < 10; i += 2 {
			mu.Lock("counter")
			counter = append(counter, i)
			mu.Unlock("counter")
			time.Sleep(50 * time.Microsecond)
		}
	}()

	wg.Wait()

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, counter)
}
