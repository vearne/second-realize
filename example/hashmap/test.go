package main

import (
	"fmt"
	"github.com/vearne/second-realize/hashmap"
)

func main() {
	list := hashmap.NewLinkedList()
	list.AddOrUpdate("1", "1")
	list.AddOrUpdate("2", "2")
	fmt.Println("list1", list.String())
	list.AddOrUpdate("4", "4")
	list.AddOrUpdate("5", "5")
	fmt.Println("list2", list.String())
	list.AddOrUpdate("3", "3")
	fmt.Println("list3", list.String())
	list.Delete("4")
	fmt.Println("list4", list.String())
}
