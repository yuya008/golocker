package golock

import (
	"testing"
	"time"
)

var recursiveLocker *RecursiveLocker

func init() {
	recursiveLocker = NewRecursiveLocker()
}

func TestRecursiveLocker(t *testing.T) {
	var all uint64
	for i := 1; i <= 1000000; i++ {
		go func(j uint64) {
			if err := recursiveLocker.Acquire(j); err != nil {
				t.Error(err)
			}
			all += j
			if err := recursiveLocker.Release(j); err != nil {
				t.Error(err)
			}

		}(uint64(i))
	}
	time.Sleep(time.Second * 3)
	t.Log(all)
}

//func TestNormalLocker(t *testing.T) {
//	var all int
//	for i := 1; i <= 50000; i++ {
//		all += i
//	}
//	t.Log(all)
//}

//func TestNormalLocker(t *testing.T) {
//	var wg sync.WaitGroup
//	var lock sync.Mutex
//	wg.Add(100000)
//	var all int
//	for i := 1; i <= 100000; i++ {
//		go func(i int) {
//			lock.Lock()
//			all += i
//			lock.Unlock()
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//	t.Log(all)
//}
