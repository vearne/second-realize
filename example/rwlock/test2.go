package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var count int
var wg sync.WaitGroup
var rw sync.RWMutex

func main() {
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(0)
	start := time.Now()
	wg.Add(40)

	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Millisecond * 20)
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
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	rw.RLock()

	fmt.Printf("读goroutine %d 正在读取...\n", n)
	time.Sleep(5 * time.Second)
	v := count

	fmt.Printf("读goroutine %d 读取结束，值为：%d\n", n, v)
	wg.Done()
	//fmt.Println("----")
	rw.RUnlock()

}

func write(n int) {
	rw.Lock()
	fmt.Printf("写goroutine %d 正在写入...\n", n)

	time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(200)))

	v := rand.Intn(1000)
	count = v

	fmt.Printf("写goroutine %d 写入结束，新值为：%d\n", n, v)
	wg.Done()
	rw.Unlock()
}
