package golocker

import (
	"testing"
	"time"
	"sync"
)

var keeper *Keeper

func init() {
	keeper = NewKeeper()
}

func TestKeeper(t *testing.T) {
	var w sync.WaitGroup
	w.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			keeper.Wait()
			t.Logf("%d 唤醒\n", i)
			w.Done()
		}(i)
	}
	t.Log("Notify")
	keeper.Notify()
	time.Sleep(time.Second * 1)
	t.Log("Notify")
	keeper.Notify()
	time.Sleep(time.Second * 1)
	t.Log("NotifyAll")
	keeper.NotifyAll()
	w.Wait()
}

func BenchmarkKeeper(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		keeper.Wait()
		keeper.Notify()
	}
}