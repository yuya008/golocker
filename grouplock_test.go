package golocker

import (
	"testing"
	"sync"
)

var groupLocker *GroupLocker

func init() {
	groupLocker = NewGroupLocker()
}

func TestGroupLocker1(t *testing.T) {
	var b uint64
	var wg sync.WaitGroup
	wg.Add(100000)
	for i := 1; i <= 100000; i++ {
		go func(i uint64) {
			groupLocker.Acquire(NewGroup(i))
			b += i
			groupLocker.Release()
			wg.Done()
		}(uint64(i))
	}
	wg.Wait()
	t.Log(b)
}

func TestGroupLocker2(t *testing.T) {
	group1 := NewGroup(100)
	group1.AddSlave(90)
	group1.AddSlave(80)
	group1.AddSlave(70)

	group2 := NewGroup(90)
	group3 := NewGroup(80)
	group4 := NewGroup(70)
	group5 := NewGroup(60)

	groupLocker.Acquire(group1)
	groupLocker.Acquire(group2)
	groupLocker.Acquire(group3)
	groupLocker.Acquire(group4)
	if !groupLocker.TryAcquire(group5) {
		t.Log("加锁失败")
	} else {
		t.Error("加锁成功")
	}
	groupLocker.Release()
	if !groupLocker.TryAcquire(group5) {
		t.Error("加锁不成功")
	}
	if groupLocker.TryAcquire(group2) {
		t.Error("加锁成功")
	}
	groupLocker.Release()
}

func BenchmarkGroupLocker1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		groupLocker.Acquire(NewGroup(uint64(i)))
		groupLocker.Release()
	}
}

func BenchmarkGroupLocker2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !groupLocker.TryAcquire(NewGroup(uint64(i))) {
			b.Error("加锁失败")
		}
		groupLocker.Release()
	}
}