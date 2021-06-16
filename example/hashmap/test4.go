package main

import (
	"fmt"
	"github.com/vearne/second-realize/hashmap"
)

func main() {
	fmt.Printf("%#x\n", hashmap.HashCode("1"))
	fmt.Printf("%#x\n", hashmap.HashCode("2"))
	fmt.Printf("%#x\n", hashmap.HashCode("3"))
}
