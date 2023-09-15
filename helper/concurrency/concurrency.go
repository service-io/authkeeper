// Package concurrency
// @author tabuyos
// @since 2023/8/3
// @description concurrency
package concurrency

import (
	cmap "github.com/orcaman/concurrent-map/v2"
)

// var wg sync.WaitGroup // WaitGroup 来保证子 goroutine 完成任务之前，主协程不会退出。
// wg.Wait()

var locking = cmap.New[*mutexLock]()

type MutexLock interface {
	Lock()
	Unlock()
}

func NewMutexLockRoutine() MutexLock {
	lock := &mutexLock{
		semaphore: make(semaphore, 1),
	}
	return lock
}

func NewMutexLockContent(key string) MutexLock {
	if len(key) == 0 {
		panic("key 不能为空")
	}
	lock, ok := locking.Get(key)
	// 存在会返回 true
	if ok {
		lock.times += 1
		return lock
	}
	lock = &mutexLock{
		semaphore: make(semaphore, 1),
		key:       key,
		times:     1,
	}
	locking.Set(key, lock)
	return lock
}

type mutexLock struct {
	semaphore semaphore
	key       string
	times     int
}

type Empty interface{}

type semaphore chan Empty

// P acquire n resources
func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

// V release n resources
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

// Lock acquire lock
func (ml *mutexLock) Lock() {
	ml.semaphore.P(1)
}

// Unlock release lock
func (ml *mutexLock) Unlock() {
	key := ml.key
	if len(key) == 0 {
		ml.semaphore.V(1)
		return
	}

	if ml.times > 0 {
		ml.times -= 1
		ml.semaphore.V(1)
		return
	}

	locking.Remove(key)

	ml.semaphore.V(1)
}
