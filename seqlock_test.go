package golocker

import (
	"testing"
	"sync"
)

var seqLocker *SeqLocker

func init() {
	seqLocker = NewSeqLocker()
}

func TestSeqLocker(t *testing.T) {
	var b uint64
	var wg sync.WaitGroup
	wg.Add(1000000)
	for i := 1; i <= 1000000; i++ {
		go func(i uint64) {
			seqLocker.Lock()
			b += i
			seqLocker.Unlock()
			wg.Done()
		}(uint64(i))
	}
	wg.Wait()
	t.Log(b)
}

func BenchmarkSeqLocker(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		seqLocker.Lock()
		seqLocker.Unlock()
	}
}
