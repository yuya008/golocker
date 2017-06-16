package golocker

import (
	"testing"
	"sync"
)

var recursiveLocker *RecursiveLocker

func init() {
	recursiveLocker = NewRecursiveLocker()
}

func TestRecursiveLocker(t *testing.T) {
	var all int64
	var wg sync.WaitGroup
	wg.Add(10000000)
	for i := 1; i <= 10000000; i++ {
		go func(j int64) {
			recursiveLocker.Acquire(j)
			all += j
			recursiveLocker.Release()
			wg.Done()
		}(int64(i))
	}
	wg.Wait()
	t.Log(all)
}

func TestReentrant(t *testing.T) {
	recursiveLocker.Acquire(1)
	recursiveLocker.Acquire(1)
	recursiveLocker.Acquire(1)
	ok := recursiveLocker.TryAcquire(2)
	if !ok {
		t.Log("加锁失败")
	}
	recursiveLocker.Release()
	ok = recursiveLocker.TryAcquire(100)
	if !ok {
		t.Error("竟然加锁不成功")
	}
	recursiveLocker.Release()
}

func BenchmarkRecursiveLocker_1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		recursiveLocker.Acquire(int64(i+1))
		recursiveLocker.Release()
	}
}
