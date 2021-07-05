package skiplist

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

type Node struct {
	key   int
	value interface{}
	// 在某一level，在链表中，指向下一个node
	forwards []*Node
}

func NewNode(maxLevel int, key int, value interface{}) *Node {
	var node Node
	node.forwards = make([]*Node, maxLevel)
	node.key = key
	node.value = value
	return &node
}

type SkipList struct {
	head *Node
	tail *Node
	// 从0开始
	topLevel    int
	maxLevel    int
	count       int
	probability float64
}

func NewSkipList(p float64, maxLevel int) *SkipList {
	var list SkipList
	list.probability = p
	list.topLevel = 0
	list.maxLevel = maxLevel
	list.count = 0
	// 尝试加入岗哨位
	list.head = NewNode(maxLevel, math.MinInt64, nil)
	list.tail = NewNode(maxLevel, math.MaxInt64, nil)
	for i := 0; i < maxLevel; i++ {
		list.head.forwards[i] = list.tail
	}
	return &list
}

func (list *SkipList) Print() {
	top := list.topLevel
	fmt.Printf("-------level-%03d------\n", top)
	for ; top >= 0; top-- {
		fmt.Printf("level-%2d:", top)
		// level-{top}
		node := list.head
		for node != nil {
			if node.key == math.MinInt64 {
				fmt.Print("[#MIN]")
			} else if node.key == math.MaxInt64 {
				fmt.Print("[#MAX]")
			} else {
				fmt.Printf("[%04d]", node.key)
			}
			//fmt.Printf(buildMargin(node.marginStep[top]))
			node = node.forwards[top]
		}
		fmt.Printf("\n")
	}
}

func (list *SkipList) Add(key int, value interface{}) {
	prevs := make([]*Node, list.maxLevel)
	// 1. 定位它的位置
	top := list.topLevel
	node := list.head
	targetKey := key
	for top >= 0 {
		if targetKey == node.forwards[top].key {
			// 直接找到了
			node = node.forwards[top]
			node.value = value
			return
		} else if targetKey < node.forwards[top].key {
			prevs[top] = node
			top--
		} else {
			node = node.forwards[top]
		}
	}

	// 2. 开始执行插入操作
	newNode := NewNode(100, key, value)
	newNode.forwards[0] = prevs[0].forwards[0]
	prevs[0].forwards[0] = newNode
	level := 1
	for level < list.maxLevel && isPromote(list.probability) {
		if prevs[level] == nil {
			prevs[level] = list.head
		}
		newNode.forwards[level] = prevs[level].forwards[level]
		prevs[level].forwards[level] = newNode
		list.topLevel = max(list.topLevel, level)
		level++
	}
	list.count++
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (list *SkipList) Delete(key int) bool {
	prevs := make([]*Node, 100)

	// 1. 定位它的位置
	top := list.topLevel
	node := list.head
	targetKey := key
	occurMaxLevel := -1
	for top >= 0 {
		if targetKey == node.forwards[top].key {
			// 直接找到了
			prevs[top] = node
			occurMaxLevel = max(occurMaxLevel, top)
			top--
		} else if targetKey < node.forwards[top].key {
			prevs[top] = node
			top--
		} else {
			node = node.forwards[top]
		}
	}

	if occurMaxLevel == -1 { // 没有找到
		return false
	}
	// 逐层从多个链表中删除
	for i := occurMaxLevel; i >= 0; i-- {
		previous := prevs[i]
		node := previous.forwards[i]
		previous.forwards[i] = node.forwards[i]
	}

	log.Println("list.topLevel", list.topLevel, "occurMaxLevel", occurMaxLevel)
	if list.topLevel == occurMaxLevel {
		for i := occurMaxLevel; i >= 0; i-- {
			if list.head.forwards[i] == list.tail {
				list.topLevel = max(i-1, 0)
			} else {
				break
			}
		}
	}

	list.count--
	return true
}
func (list *SkipList) Size() int {
	return list.count
}
func (list *SkipList) GetTopLevel() int {
	return list.topLevel
}

func (list *SkipList) Find(key int) (*Node, bool) {
	top := list.topLevel
	node := list.head
	targetKey := key
	for top >= 0 {
		if targetKey == node.forwards[top].key {
			// 直接找到了
			node = node.forwards[top]
			return node, true
		} else if targetKey < node.forwards[top].key {
			//prevs[top] = node
			top--
		} else {
			node = node.forwards[top]
		}
	}
	return nil, false
}

func isPromote(p float64) bool {
	return rand.Intn(100) <= int(p*float64(100))
}
