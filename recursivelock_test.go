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
	var all uint64
	var wg sync.WaitGroup
	wg.Add(10000000)
	for i := 1; i <= 10000000; i++ {
		go func(j uint64) {
			if err := recursiveLocker.Acquire(j); err != nil {
				t.Error(err)
			}
			all += j
			recursiveLocker.Release()
			wg.Done()
		}(uint64(i))
	}
	wg.Wait()
	t.Log(all)
}

func TestReentrant(t *testing.T) {
	if err := recursiveLocker.Acquire(1); err != nil {
		t.Error(err)
	}
	if err := recursiveLocker.Acquire(1); err != nil {
		t.Error(err)
	}
	err, ok := recursiveLocker.TryAcquire(2)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Log("加锁失败")
	}
	recursiveLocker.Release()
	err, ok = recursiveLocker.TryAcquire(100)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("竟然加锁不成功")
	}
	recursiveLocker.Release()
}

func BenchmarkRecursiveLocker_1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := recursiveLocker.Acquire(uint64(i+1)); err != nil {
			b.Error(err)
		}
		recursiveLocker.Release()
	}
}
