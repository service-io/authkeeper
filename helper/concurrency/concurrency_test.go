// Package concurrency
// @author tabuyos
// @since 2023/8/3
// @description concurrency
package concurrency

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	lock := NewMutexLockRoutine()
	lock.Lock()
	fmt.Println(123)
	lock.Unlock()
}

func TestGoRoutineWithoutLock(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	// 无序打印
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestGoRoutineWithLock(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	lock := NewMutexLockRoutine()
	// 有序打印
	for i := 0; i < 10; i++ {
		lock.Lock()
		i := i
		go func() {
			fmt.Println(i)
			lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestOneContentWithLock(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	name := "a"
	lock := NewMutexLockContent(name)
	// 有序打印
	for i := 0; i < 10; i++ {
		lock.Lock()
		i := i
		go func() {
			fmt.Println(i)
			lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestContentWithoutLock(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(50)
	for i := 0; i < 5; i++ {
		// 无序打印
		for j := 0; j < 10; j++ {
			m := i
			n := j
			go func() {
				fmt.Printf("%v -> %v\n", m, n)
				wg.Done()
			}()
		}
		// 保证每个组都打印完
		time.Sleep(10 * time.Millisecond)
	}
	wg.Wait()
}

func TestMultiContentWithLock(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(50)
	for i := 0; i < 5; i++ {
		lock := NewMutexLockContent(strconv.Itoa(i))
		// 有序打印
		for j := 0; j < 10; j++ {
			lock.Lock()
			m := i
			n := j
			go func() {
				fmt.Printf("%v -> %v\n", m, n)
				lock.Unlock()
				wg.Done()
			}()
		}
		// 保证每个组都打印完
		time.Sleep(10 * time.Millisecond)
	}
	wg.Wait()
}

func TestMultiContentWithLockUseRoutine(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(500)
	for i := 0; i < 50; i++ {
		m := i
		go func() {
			lock := NewMutexLockContent(strconv.Itoa(m))
			// 有序打印
			for j := 0; j < 10; j++ {
				lock.Lock()
				n := j
				go func() {
					fmt.Printf("%v -> %v\n", m, n)
					lock.Unlock()
					wg.Done()
				}()
			}
		}()
	}
	wg.Wait()
	// fmt.Println(len(locking))
}
