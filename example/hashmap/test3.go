package main

import (
	"fmt"
	"github.com/vearne/second-realize/hashmap"
	"strconv"
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
	h := hashmap.NewHashMap(1, 1)
	for i := 0; i < 100000; i++ {
		if i%10000 == 0 {
			fmt.Println("setMaxduration1-i", i)
		}
		key := strconv.Itoa(i)
		start := time.Now()
		h.Set(key, key)
		maxCost = max(maxCost, time.Since(start))
	}
	fmt.Println("setMaxduration1", maxCost)
	fmt.Println("setMaxduration1", len(h.HashTable.Load().([]*hashmap.LinkedList)))
}

func setMaxduration2() {
	var maxCost time.Duration
	h := make(map[string]string, 1)
	for i := 0; i < 100000; i++ {
		if i%10000 == 0 {
			fmt.Println("setMaxduration2-i", i)
		}
		key := strconv.Itoa(i)
		start := time.Now()
		h[key] = key
		maxCost = max(maxCost, time.Since(start))
	}
	fmt.Println("setMaxduration2", maxCost)
}

func main() {
	setMaxduration1()
	setMaxduration2()
}
