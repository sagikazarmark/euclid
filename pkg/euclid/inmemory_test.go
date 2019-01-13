package euclid

import (
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// UintSlice attaches the methods of Interface to []int, sorting in increasing order.
type UintSlice []uint

func (p UintSlice) Len() int           { return len(p) }
func (p UintSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p UintSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func TestInMemorySequence_Next(t *testing.T) {
	var sequence InMemorySequence

	var wg sync.WaitGroup

	wg.Add(2)

	var oddCounter, evenCounter []uint

	go func() {
		defer wg.Done()

		for i := 2; i < 10; i += 2 {
			evenCounter = append(evenCounter, sequence.Next("counter"))
			time.Sleep(500 * time.Microsecond)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 1; i < 10; i += 2 {
			oddCounter = append(oddCounter, sequence.Next("counter"))
			time.Sleep(500 * time.Microsecond)
		}
	}()

	wg.Wait()

	assert.True(t, sort.IsSorted(UintSlice(oddCounter)))
	assert.True(t, sort.IsSorted(UintSlice(evenCounter)))
}
