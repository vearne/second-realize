package main

import (
	"fmt"
	"github.com/vearne/second-realize/rwlock"
	"math/rand"
	"sync"
	"time"
)

var count int
var wg sync.WaitGroup
var rw *rwlock.RWLocker

func main() {
	// init
	rw = rwlock.NewRWLocker(rwlock.WithMaxReadLocked(3))
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(0)
	start := time.Now()
	wg.Add(80)

	go func() {
		for i := 0; i < 60; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			go read(i)
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			//time.Sleep(time.Millisecond *time.Duration(rand.Intn(100)))
			go write(i)
		}
	}()

	wg.Wait()
	fmt.Println("cost:", time.Since(start))
}

func read(n int) {
	rw.RLock()
	fmt.Printf("读goroutine %d 正在读取...\n", n)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
	v := count
	fmt.Printf("读goroutine %d 读取结束，值为：%d\n", n, v)
	wg.Done()

	rw.RUnLock()
}

func write(n int) {
	rw.WLock()
	fmt.Printf("写goroutine %d 正在写入...\n", n)
	time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(200)))

	v := rand.Intn(1000)

	count = v

	fmt.Printf("写goroutine %d 写入结束，新值为：%d\n", n, v)
	wg.Done()

	rw.WUnLock()
}
