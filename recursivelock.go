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
	signalChannel chan bool
}

func NewRecursiveLocker() *RecursiveLocker {
	return &RecursiveLocker{
		signalChannel: make(chan bool),
	}
}

func (locker *RecursiveLocker) Acquire(gid uint64) error {
	if gid == 0 {
		return errors.New("goroutine id == 0")
	}
	for {
		if atomic.CompareAndSwapUint64(&locker.gid, 0, gid) {
			break
		} else if atomic.CompareAndSwapUint64(&locker.gid, gid, gid) {
			break
		}
		<- locker.signalChannel
	}
	return nil
}

func (locker *RecursiveLocker) TryAcquire() (uint64, bool) {
	return 0, false
}

func (locker *RecursiveLocker) Release(gid uint64) error {
	if gid == 0 {
		return errors.New("goroutine id == 0")
	}
	if !atomic.CompareAndSwapUint64(&locker.gid, gid, 0) {
		return errors.New("goroutine id invalid")
	}
	locker.signalChannel <- true
	return nil
}
