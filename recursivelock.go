//
// 递归锁实现
//
package golocker

import (
	"sync"
)

type RecursiveLocker struct {
	gid int64
	mutex sync.Mutex
	cond *sync.Cond
}

func NewRecursiveLocker() *RecursiveLocker {
	locker := &RecursiveLocker{
		gid: -1,
	}
	locker.cond = sync.NewCond(&locker.mutex)
	return locker
}

func (locker *RecursiveLocker) Acquire(gid int64) {
	if gid < 0 {
		panic("goroutine id < 0")
	}
	locker.mutex.Lock()
	defer locker.mutex.Unlock()
	for {
		if locker.gid < 0 {
			locker.gid = gid
			return
		} else if locker.gid == gid {
			return
		}
		locker.cond.Wait()
	}
}

func (locker *RecursiveLocker) TryAcquire(gid int64) (ret bool) {
	if gid < 0 {
		panic("goroutine id < 0")
	}
	locker.mutex.Lock()
	defer locker.mutex.Unlock()
	if locker.gid < 0 {
		locker.gid = gid
		ret = true
	} else if locker.gid == gid {
		ret = true
	} else {
		ret = false
	}
	return
}

func (locker *RecursiveLocker) Release() {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()
	locker.gid = -1
	locker.cond.Signal()
}
