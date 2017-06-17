package golocker

import (
	"testing"
	"sync"
	"fmt"
)

var countLocker *CountLocker

func init() {
	countLocker = NewCountLocker()
}

func TestCountLocker(t *testing.T) {
	var all int64
	var wg sync.WaitGroup
	wg.Add(1000000)
	for i := 1; i <= 1000000; i++ {
		go func(j int64) {
			countLocker.Acquire(j)
			countLocker.Acquire(j)
			countLocker.Acquire(j)
			all += j
			countLocker.Release()
			countLocker.Release()
			countLocker.Release()
			wg.Done()
		}(int64(i))
	}
	wg.Wait()
	t.Log(all)
}

func BenchmarkCountLocker(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		countLocker.Acquire(int64(i+1))
		countLocker.Acquire(int64(i+1))
		countLocker.Release()
		countLocker.Release()
	}
}
