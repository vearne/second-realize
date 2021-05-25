package rwlock

import (
	"sync"
)

const (
	UnLocked = iota
	ReadLocked
	WriteLocked
)

type RWLocker struct {
	Options
	sync.Mutex
	// 已经加上读锁的协程数量
	readLockCounter int
	// 0：unlocked 1: readLocked 2: writeLocked
	status int
	/*
		引入 wantReadLock wantWriteLock
		的目的是为了防止读饥饿或写饥饿问题
	*/
	// 有协程想要想要加读锁
	wantReadLock int
	// 有协程想要想要加写锁
	wantWriteLock int
	// 等待加读锁的条件发生
	canAddReadLock *sync.Cond
	// 等待加写锁的条件发生
	canAddWriteLock *sync.Cond

	//// 最大同时可以加的读锁的数量
	//// 相当于可以控制并发的读协程的数量
	//maxReadLocked int
	//// 为防止出现写协程长时间等待的问题
	//// 超过此值，读协程就需要挂起自己
	//// 减小此值，可以保证写入更及时
	//maxWaitWriteGoroutine int
}

func NewRWLocker(opts ...Option) *RWLocker {
	l := RWLocker{}
	l.readLockCounter = 0
	l.status = 0
	l.wantReadLock = 0
	l.wantWriteLock = 0
	l.canAddReadLock = sync.NewCond(&l)
	l.canAddWriteLock = sync.NewCond(&l)
	l.maxReadLocked = 3
	l.maxWaitWriteGoroutine = 3

	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated
		opt(&l.Options)
	}
	return &l
}

func (rw *RWLocker) WLock() {
	rw.Lock()
	rw.wantWriteLock++
	for !(rw.status == UnLocked) {
		//log.Println("status:", rw.status, "readLockCounter:", rw.readLockCounter,
		//	"wantReadLock:", rw.wantReadLock,
		//	"wantWriteLock:", rw.wantWriteLock)
		rw.canAddWriteLock.Wait()
	}
	rw.status = WriteLocked
	rw.Unlock()
}

func (rw *RWLocker) WUnLock() {
	rw.Lock()
	// Operating critical section data
	rw.status = UnLocked

	if rw.wantReadLock > 0 {
		rw.canAddReadLock.Broadcast()
	} else if rw.wantWriteLock > 0 {
		rw.canAddWriteLock.Broadcast()
	}

	rw.wantWriteLock--
	rw.Unlock()
}

func (rw *RWLocker) RLock() {
	rw.Lock()
	rw.wantReadLock++
	for !(rw.status == UnLocked || (rw.status == ReadLocked &&
		(rw.wantWriteLock < rw.maxWaitWriteGoroutine || rw.readLockCounter < rw.maxReadLocked))) {
		rw.canAddReadLock.Wait()
	}
	rw.status = ReadLocked
	rw.readLockCounter++
	rw.Unlock()
}

func (rw *RWLocker) RUnLock() {
	rw.Lock()
	// Operating critical section data
	rw.readLockCounter--
	if rw.readLockCounter <= 0 {
		rw.status = UnLocked
	}
	if rw.wantWriteLock > 0 {
		rw.canAddWriteLock.Broadcast()
	} else if rw.wantReadLock > 0 {
		rw.canAddReadLock.Broadcast()
	}
	rw.wantReadLock--
	rw.Unlock()
}
