//
// 递归锁实现
//
package golock

import (
	"errors"
	"sync/atomic"
)

type RecursiveLocker struct {
	gid uint64
	signalChan chan bool
}

func NewRecursiveLocker() *RecursiveLocker {
	return &RecursiveLocker{
		signalChan: make(chan bool),
	}
}

func (locker *RecursiveLocker) Acquire(gid uint64) error {
	if gid == 0 {
		return errors.New("goroutine id == 0")
	}
	for {
		if atomic.CompareAndSwapUint64(&locker.gid, 0, gid) ||
				atomic.CompareAndSwapUint64(&locker.gid, gid, gid) {
			break
		}
		<- locker.signalChan
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

func (locker *RecursiveLocker) Release(gid uint64) error {
	if gid == 0 {
		return errors.New("goroutine id == 0")
	}
	if !atomic.CompareAndSwapUint64(&locker.gid, gid, 0) {
		return errors.New("goroutine id invalid")
	}
	locker.signalChan <- true
	return nil
}
