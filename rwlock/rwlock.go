package rwlock

import (
	"fmt"
	"sync"
)

const (
	UnLocked = iota
	ReadLocked
	WriteLocked
)

type RWLocker struct {
	sync.Mutex
	readLockCounter int
	// 0：unlocked 1: readLocked 2: writeLocked
	status int
	canAddReadLock *sync.Cond
	canAddWriteLock *sync.Cond
}


func NewRWLocker() *RWLocker{
	l := RWLocker{}
	l.readLockCounter = 0
	l.status = 0
	l.canAddReadLock = sync.NewCond(&l)
	l.canAddWriteLock = sync.NewCond(&l)
	return &l
}

func (locker *RWLocker) WLock(){
	locker.Lock()
	for  !(locker.status == UnLocked){
		locker.canAddWriteLock.Wait()
	}
	locker.status = WriteLocked
	locker.Unlock()
}

func (locker *RWLocker) WUnLock(){
	locker.Lock()
	locker.status = UnLocked
	// Operating critical section data
	locker.canAddWriteLock.Broadcast()
	fmt.Println("--1--")
	locker.canAddReadLock.Broadcast()
	fmt.Println("--2--")
	locker.Unlock()
}

func (locker *RWLocker) RLock(){
	locker.Lock()
	for  !(locker.status == UnLocked|| locker.status == ReadLocked){
		locker.canAddReadLock.Wait()
	}
	locker.status = ReadLocked
	locker.readLockCounter++
	locker.Unlock()
}

func (locker *RWLocker) RUnLock(){
	locker.Lock()
	// Operating critical section data
	locker.readLockCounter--
	if locker.readLockCounter <= 0{
		locker.status = UnLocked
	}

	locker.canAddWriteLock.Broadcast()
	locker.Unlock()
}