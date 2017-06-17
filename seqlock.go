//
// 顺序锁实现
//
package golocker

import (
	"sync"
)

type SeqLocker struct {
	mutex sync.Mutex
	cond *sync.Cond
	wlock bool
}

func NewSeqLocker() *SeqLocker {
	locker := &SeqLocker{}
	locker.cond = sync.NewCond(&locker.mutex)
	return locker
}

func (locker *SeqLocker) Lock() {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()
	for {
		if locker.wlock == false {
			locker.wlock = true
			break
		}
		locker.cond.Wait()
	}
}

func (locker *SeqLocker) Unlock() {
	locker.wlock = false
	locker.cond.Signal()
}

func (locker *SeqLocker) RLock() {}

func (locker *SeqLocker) RUnlock() {}

func (locker *SeqLocker) TryLock() bool {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()
	if locker.wlock == false {
		locker.wlock = true
		return true
	}
	return false
}

func (locker *SeqLocker) TryRLock() bool {
	return true
}