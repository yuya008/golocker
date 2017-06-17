//
// 组锁实现
//
package golocker

import "sync"

type GroupLocker struct {
	g *Group
	mutex sync.Mutex
	cond *sync.Cond
}

func NewGroupLocker() *GroupLocker {
	locker := &GroupLocker{}
	locker.cond = sync.NewCond(&locker.mutex)
	return locker
}

func (locker *GroupLocker) Acquire(g *Group) {
	if g == nil {
		panic("*Group == nil")
	}
	locker.mutex.Lock()
	defer locker.mutex.Unlock()
	for {
		if locker.g == nil {
			locker.g = g
			return
		} else if locker.g == g {
			return
		} else {
			for _, gid := range locker.g.slaveGids {
				if gid == g.masterGid {
					return
				}
			}
		}
		locker.cond.Wait()
	}
}

func (locker *GroupLocker) Release() {
	locker.mutex.Lock()
	defer locker.mutex.Unlock()
	locker.g = nil
	locker.cond.Signal()
}

func (locker *GroupLocker) TryAcquire(g *Group) bool {
	if g == nil {
		panic("*Group == nil")
	}
	locker.mutex.Lock()
	defer locker.mutex.Unlock()
	if locker.g == nil {
		locker.g = g
		return true
	} else if locker.g == g {
		return true
	} else {
		for _, gid := range locker.g.slaveGids {
			if gid == g.masterGid {
				return true
			}
		}
	}
	return false
}

type Group struct {
	masterGid uint64
	slaveGids []uint64
}

func NewGroup(masterGid uint64) *Group {
	return &Group{
		masterGid: masterGid,
	}
}

func (g *Group) AddSlave(gid uint64) {
	g.slaveGids = append(g.slaveGids, gid)
}
