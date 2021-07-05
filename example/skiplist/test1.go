package main

import (
	"fmt"
	"github.com/vearne/second-realize/skiplist"
	"math/rand"
	"time"
)

func main() {
	//nodeList := make([]*skiplist.Node, 10)
	//fmt.Println(nodeList[1])

	list := skiplist.NewSkipList(0.5, 5)
	rand.Seed(0)
	size := 20
	for i := 0; i < size; i++ {
		key := rand.Intn(100) + 1
		fmt.Println("key:", key)
		list.Add(key, key)
		//time.Sleep(1 * time.Second)
	}
	list.Print()
	var ok bool
	keys := []int{70, 20, 100, 36, 99, 1, 83}
	for _, key := range keys {
		_, ok = list.Find(key)
		fmt.Printf("find key:%v, result:%v\n", key, ok)
	}

	keys = []int{45, 100, 27}
	for _, key := range keys {
		ok := list.Delete(key)
		fmt.Println("########################################")
		fmt.Printf("delete key:%v, result:%v\n", key, ok)
		list.Print()
		time.Sleep(1 * time.Second)
		_, ok = list.Find(key)
		fmt.Printf("find key:%v, result:%v\n", key, ok)
	}
	fmt.Println("size:", list.Size())
}
