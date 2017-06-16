//
// 自旋锁实现
//
package golocker

import (
	"sync/atomic"
)

type SpinLocker struct {
	v int32
}

func NewSpinLocker() *SpinLocker {
	return &SpinLocker{}
}

func (lock *SpinLocker) Lock() {
	for {
		if atomic.CompareAndSwapInt32(&lock.v, 0, 1) {
			return
		}
	}
}

func (lock *SpinLocker) TryLock() bool {
	return atomic.CompareAndSwapInt32(&lock.v, 0, 1)
}

func (lock *SpinLocker) Unlock() {
	atomic.CompareAndSwapInt32(&lock.v, 1, 0)
}