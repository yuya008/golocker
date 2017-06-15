//
// 自旋锁实现
//
package golock

import (
	"sync/atomic"
)

type SpinLock struct {
	v int32
}

func (lock *SpinLock) Lock() {
	for {
		if atomic.CompareAndSwapInt32(&lock.v, 0, 1) {
			break
		}
	}
}

func (lock *SpinLock) TryLock() bool {
	return atomic.CompareAndSwapInt32(&lock.v, 0, 1)
}

func (lock *SpinLock) Unlock() {
	atomic.CompareAndSwapInt32(&lock.v, 1, 0)
}