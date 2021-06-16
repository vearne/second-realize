package main

import (
	"fmt"
	"github.com/vearne/second-realize/hashmap"
	"strconv"
	"sync/atomic"
)

func main() {
	h := hashmap.NewHashMap(4, 2)
	for i := 0; i < 10000; i++ {
		key := strconv.Itoa(i)
		value := key
		h.Set(key, value)
	}
	fmt.Println(h.Get("1"))
	h.Delete("3")
	fmt.Println(h.Get("3"))
	fmt.Println(h.Get("995"))
	fmt.Println(h.Size)
	fmt.Println("setMaxduration1", "HashTableCapacity1:", len(h.HashTable.Load().([]*hashmap.LinkedList)))
	fmt.Println("setMaxduration1", "HashTableCapacity2:", h.HashTableCapacity)
	fmt.Println("setMaxduration1", "size:", atomic.LoadInt64(&h.Size))
}
