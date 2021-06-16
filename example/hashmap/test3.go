package main

import (
	"fmt"
	"github.com/vearne/second-realize/hashmap"
	"strconv"
	"sync/atomic"
	"time"
)

func max(d1, d2 time.Duration) time.Duration {
	if d1 < d2 {
		return d2
	}
	return d1
}

func setMaxduration1() {
	var maxCost time.Duration
	h := hashmap.NewHashMap(2, 5)
	for i := 0; i < 10000000; i++ {
		if i%100000 == 0 {
			fmt.Println("setMaxduration1-i", i)
		}
		key := strconv.Itoa(i)
		start := time.Now()
		h.Set(key, key)
		maxCost = max(maxCost, time.Since(start))
	}
	fmt.Println("setMaxduration1", maxCost)
	fmt.Println("setMaxduration1", "HashTableCapacity1:", len(h.HashTable.Load().([]*hashmap.LinkedList)))
	fmt.Println("setMaxduration1", "HashTableCapacity2:", h.HashTableCapacity)
	fmt.Println("setMaxduration1", "size:", atomic.LoadInt64(&h.Size))
}

func setMaxduration2() {
	var maxCost time.Duration
	h := hashmap.NewStdMap(2)
	for i := 0; i < 10000000; i++ {
		if i%100000 == 0 {
			fmt.Println("setMaxduration2-i", i)
		}
		key := strconv.Itoa(i)
		start := time.Now()
		h.Set(key, key)
		maxCost = max(maxCost, time.Since(start))
	}
	fmt.Println("setMaxduration2", maxCost)
}

func main() {
	setMaxduration1()
	setMaxduration2()
}
