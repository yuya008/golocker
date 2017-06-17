//
// 计数锁实现
//
package golocker

import (
	"sync/atomic"
)

type CountLocker struct {
	Count uint64
	recursiveLocker *RecursiveLocker
}

func NewCountLocker() *CountLocker {
	return &CountLocker{
		recursiveLocker: NewRecursiveLocker(),
	}
}

func (locker *CountLocker) Acquire(gid int64) {
	locker.recursiveLocker.Acquire(gid)
	atomic.AddUint64(&locker.Count, 1)
}

func (locker *CountLocker) Release() {
	if atomic.AddUint64(&locker.Count, ^uint64(0)) == 0 {
		locker.recursiveLocker.Release()
	}
}

func (locker *CountLocker) TryAcquire(gid int64) bool {
	if locker.recursiveLocker.TryAcquire(gid) {
		atomic.AddUint64(&locker.Count, 1)
		return true
	}
	return false
}

