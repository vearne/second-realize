package hashmap

import (
	"container/list"
	"fmt"
	"strings"
)

type LinkedList struct {
	// 存储 KeyValueItem
	innerList *list.List
}

func NewLinkedList() *LinkedList {
	l := LinkedList{}
	l.innerList = list.New()
	return &l
}

func (l *LinkedList) String() string {
	p := l.innerList.Front()
	tmp := make([]string, 0, l.innerList.Len())
	for p != nil {
		tmp = append(tmp, fmt.Sprintf("%v", p.Value))
		p = p.Next()
	}
	return strings.Join(tmp, ",")
}

// 实际添加了一条记录返回true
// 否则返回false
func (l *LinkedList) AddOrUpdate(key string, value interface{}) bool {
	item := KeyValueItem{key, value}
	if l.innerList.Len() <= 0 {
		l.innerList.PushBack(item)
		return true
	}

	p := l.innerList.Front()
	for p != nil && p.Value.(KeyValueItem).Key < key {
		p = p.Next()
	}
	if p == nil { // 链表中没有数据，或者key需要插入到最后
		l.innerList.PushBack(item)
		return true
	} else if p.Value.(KeyValueItem).Key == key {
		p.Value = item
		return false
	} else {
		l.innerList.InsertBefore(item, p)
		return true
	}
}

// 实际删除了一个数据返回true
// 没有找到对应的key，返回false
func (l *LinkedList) Delete(key string) bool {
	p := l.innerList.Front()
	for p != nil && p.Value.(KeyValueItem).Key != key {
		p = p.Next()
	}
	if p != nil {
		l.innerList.Remove(p)
		return true
	}
	return false
}

func (l *LinkedList) Get(key string) (interface{}, bool) {
	p := l.innerList.Front()
	for p != nil && p.Value.(KeyValueItem).Key != key {
		p = p.Next()
	}
	if p != nil {
		return p.Value.(KeyValueItem).Value, true
	}
	return nil, false
}
