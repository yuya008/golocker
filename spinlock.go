//
// 自旋锁实现
//
package golock

import (
	"sync/atomic"
)

type SpinLocker struct {
	v int32
}

func (lock *SpinLocker) Lock() {
	for {
		if atomic.CompareAndSwapInt32(&lock.v, 0, 1) {
			break
		}
	}
}

func (lock *SpinLocker) TryLock() bool {
	return atomic.CompareAndSwapInt32(&lock.v, 0, 1)
}

func (lock *SpinLocker) Unlock() {
	atomic.CompareAndSwapInt32(&lock.v, 1, 0)
}