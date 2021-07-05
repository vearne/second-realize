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
	rw = rwlock.NewRWLocker()
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(0)
	start := time.Now()
	wg.Add(10)

	for i := 0; i < 5; i++ {
		go read(i)
	}

	for i := 0; i < 5; i++ {
		//time.Sleep(time.Millisecond *time.Duration(rand.Intn(100)))
		go write(i)
	}

	wg.Wait()
	fmt.Println("cost:", time.Since(start))
}

func read(n int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	rw.RLock()
	fmt.Printf("读goroutine %d 正在读取...\n", n)
	time.Sleep(5 * time.Second)
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
