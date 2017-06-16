//
// 递归锁实现
//
package golocker

import (
	"errors"
	"sync/atomic"
)

type RecursiveLocker struct {
	gid uint64
	keeper *Keeper
}

func NewRecursiveLocker() *RecursiveLocker {
	return &RecursiveLocker{
		keeper: NewKeeper(),
	}
}

func (locker *RecursiveLocker) Acquire(gid uint64) error {
	if gid == 0 {
		return errors.New("goroutine id == 0")
	}
	for {
		if atomic.CompareAndSwapUint64(&locker.gid, 0, gid) ||
				atomic.CompareAndSwapUint64(&locker.gid, gid, gid) {
			return nil
		}
		locker.keeper.Wait()
	}
	return nil
}

func (locker *RecursiveLocker) TryAcquire(gid uint64) (error, bool) {
	if gid == 0 {
		return errors.New("goroutine id == 0"), false
	}
	if atomic.CompareAndSwapUint64(&locker.gid, 0, gid) ||
			atomic.CompareAndSwapUint64(&locker.gid, gid, gid) {
		return nil, true
	}
	return nil, false
}

func (locker *RecursiveLocker) Release() {
	atomic.StoreUint64(&locker.gid, 0)
	locker.keeper.Notify()
}
