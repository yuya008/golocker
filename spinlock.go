//
// 自旋锁实现
//
package golocker

import (
	"sync/atomic"
)

const (
	DefaultMaxSpinTimes = 100000
)

type SpinLocker struct {
	v int32
	MaxSpinTimes uint32
	keeper *Keeper
}

func NewSpinLocker() *SpinLocker {
	locker := &SpinLocker{
		keeper: NewKeeper(),
	}
	return locker
}

func (lock *SpinLocker) Lock() {
	if lock.MaxSpinTimes == 0 {
		lock.MaxSpinTimes = DefaultMaxSpinTimes
	}
	var curSpinTimes uint32
	for {
		curSpinTimes = 0
		for ; curSpinTimes < lock.MaxSpinTimes; curSpinTimes++ {
			if atomic.CompareAndSwapInt32(&lock.v, 0, 1) {
				return
			}
		}
		lock.keeper.Wait()
	}
	return
}

func (lock *SpinLocker) TryLock() bool {
	return atomic.CompareAndSwapInt32(&lock.v, 0, 1)
}

func (lock *SpinLocker) Unlock() {
	atomic.CompareAndSwapInt32(&lock.v, 1, 0)
}