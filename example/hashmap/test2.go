package main

import (
	"fmt"
	"github.com/vearne/second-realize/hashmap"
	"strconv"
)

func main() {
	h := hashmap.NewHashMap(4, 1)
	for i := 0; i < 1000; i++ {
		key := strconv.Itoa(i)
		value := key
		h.Set(key, value)
	}
	fmt.Println(h.Get("1"))
	h.Delete("3")
	fmt.Println(h.Get("3"))
	fmt.Println(h.Get("995"))
	fmt.Println(h.Size)
}
