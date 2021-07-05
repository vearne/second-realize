package skiplist

import (
	"fmt"
	"math"
	"math/rand"
)

type Node struct {
	key   int
	value interface{}
	// 在某一level，在链表中，指向下一个node
	levelNexts []*Node
}

func NewNode(topLevel int, key int, value interface{}) *Node {
	var node Node
	node.levelNexts = make([]*Node, topLevel+1)
	node.key = key
	node.value = value
	return &node
}

type SkipList struct {
	head *Node
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
	list.maxLevel = 5
	list.count = 0
	// 尝试加入岗哨位
	list.head = NewNode(100, math.MinInt64, nil)
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
				fmt.Print("[MINI]")
			} else {
				fmt.Printf("[%04d]", node.key)
			}
			//fmt.Printf(buildMargin(node.marginStep[top]))
			node = node.levelNexts[top]
		}
		fmt.Printf("\n")
	}
}

func (list *SkipList) Add(key int, value interface{}) {
	prevs := make([]*Node, list.maxLevel+1)

	// 1. 定位它的位置
	top := list.topLevel
	node := list.head
	targetKey := key
	for top >= 0 {
		if node.levelNexts[top] != nil && targetKey == node.levelNexts[top].key {
			// 直接找到了
			node = node.levelNexts[top]
			node.value = value
			return
		} else if node.levelNexts[top] == nil || targetKey < node.levelNexts[top].key {
			prevs[top] = node
			top--
		} else {
			node = node.levelNexts[top]
		}
	}

	// 2. 开始执行插入操作
	newNode := NewNode(100, key, value)
	newNode.levelNexts[0] = prevs[0].levelNexts[0]
	prevs[0].levelNexts[0] = newNode
	level := 1
	for level < list.maxLevel && isPromote(list.probability) {
		if prevs[level] == nil {
			prevs[level] = list.head
		}
		newNode.levelNexts[level] = prevs[level].levelNexts[level]
		prevs[level].levelNexts[level] = newNode
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

func (list *SkipList) Delete(key int) bool {
	prevs := make([]*Node, 100)

	// 1. 定位它的位置
	top := list.topLevel
	node := list.head
	targetKey := key
	occurMaxLevel := -1
	for top >= 0 {
		if node.levelNexts[top] != nil && targetKey == node.levelNexts[top].key {
			// 直接找到了
			prevs[top] = node
			occurMaxLevel = max(occurMaxLevel, top)
			top--
		} else if node.levelNexts[top] == nil || targetKey < node.levelNexts[top].key {
			prevs[top] = node
			top--
		} else {
			node = node.levelNexts[top]
		}
	}

	if occurMaxLevel == -1 { // 没有找到
		return false
	}
	// 逐层从多个链表中删除
	for i := 0; i <= occurMaxLevel; i++ {
		previous := prevs[i]
		node := previous.levelNexts[i]
		previous.levelNexts[i] = node.levelNexts[i]
	}
	list.count--
	return true
}
func (list *SkipList) Size() int {
	return list.count
}

func (list *SkipList) Find(key int) (*Node, bool) {
	top := list.topLevel
	node := list.head
	targetKey := key
	for top >= 0 {
		if node.levelNexts[top] != nil && targetKey == node.levelNexts[top].key {
			// 直接找到了
			node = node.levelNexts[top]
			return node, true
		} else if node.levelNexts[top] == nil || targetKey < node.levelNexts[top].key {
			//prevs[top] = node
			top--
		} else {
			node = node.levelNexts[top]
		}
	}
	return nil, false
}

func isPromote(p float64) bool {
	return rand.Intn(100) <= int(p*float64(100))
}
