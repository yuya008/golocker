package golocker

import (
	"testing"
	"sync"
)

var spinLocker *SpinLocker

func init() {
	spinLocker = NewSpinLocker()
}

func TestSpinLocker(t *testing.T) {
	var all int64
	var wg sync.WaitGroup
	wg.Add(100000)
	for i := 1; i <= 100000; i++ {
		go func() {
			spinLocker.Lock()
			all++
			spinLocker.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(all)
}

func BenchmarkSpinLocker(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		spinLocker.Lock()
		spinLocker.Unlock()
	}
}