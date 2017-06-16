//
// 一种线程等待工具，提供了更清晰的接口
//
package golocker

import (
	"sync"
)

type Keeper struct {
	mutex sync.Mutex
	cond *sync.Cond
}

func NewKeeper() *Keeper {
	keeper := &Keeper{}
	keeper.cond = sync.NewCond(&keeper.mutex)
	return keeper
}

func (keeper *Keeper) Wait() {
	keeper.mutex.Lock()
	defer keeper.mutex.Unlock()
	keeper.cond.Wait()
}

func (keeper *Keeper) Notify() {
	keeper.cond.Signal()
}

func (keeper *Keeper) NotifyAll() {
	keeper.cond.Broadcast()
}
