package hashmap

import (
	"sync/atomic"
)

const (
	Running    = 1
	NotRunning = 0
)

type KeyValueItem struct {
	Key   string
	Value interface{}
}

type HashMap struct {
	// 装载因子
	LoadFactor        float64
	Size              int64
	HashTableCapacity uint64
	// []*LinkedList
	HashTable atomic.Value
	// rehash相关
	// 1:表示正在执行
	// 0:没有执行
	rehashDoing uint32
	// tableIdx + keyIdentify
	tableIdx uint64
	// 已经被迁移的某个key
	keyIdentify string
	// []*LinkedList
	NewHashTable atomic.Value
}

func NewHashMap(hashTableCapacity uint64, loadFactor float64) *HashMap {
	h := HashMap{}
	h.Size = 0
	h.rehashDoing = NotRunning
	h.LoadFactor = loadFactor
	h.HashTableCapacity = hashTableCapacity
	hashTableValue := make([]*LinkedList, hashTableCapacity)
	h.tableIdx = 0
	h.keyIdentify = ""
	var i uint64
	for i = 0; i < hashTableCapacity; i++ {
		hashTableValue[i] = NewLinkedList()
	}
	h.HashTable.Store(hashTableValue)
	return &h
}

func (h *HashMap) Set(key string, value interface{}) {
	if h.Size > int64(float64(h.HashTableCapacity)*h.LoadFactor) && atomic.
		CompareAndSwapUint32(&h.rehashDoing, NotRunning, Running) {
		go func() {
			h.HashTableCapacity = h.HashTableCapacity * 2
			newHashTableValue := make([]*LinkedList, h.HashTableCapacity)
			var i uint64
			for i = 0; i < h.HashTableCapacity; i++ {
				newHashTableValue[i] = NewLinkedList()
			}
			h.NewHashTable.Store(newHashTableValue)
			oldHashTableValue := h.HashTable.Load().([]*LinkedList)
			for h.tableIdx = 0; h.tableIdx < uint64(len(oldHashTableValue)); h.tableIdx++ {
				// Iterate through list and print its contents.
				for e := oldHashTableValue[h.tableIdx].innerList.Front(); e != nil; e = e.Next() {
					item := e.Value.(KeyValueItem)
					// rehash并不影响size
					setItem(newHashTableValue, item.Key, item.Value)
					h.keyIdentify = item.Key
					//fmt.Println("---", h.keyIdentify)
					//time.Sleep(time.Second)
				}
			}
			h.HashTable.Store(newHashTableValue)
			// 修改
			atomic.CompareAndSwapUint32(&h.rehashDoing, Running, NotRunning)
		}()
	}

	var ok bool
	if h.rehashDoing == Running {
		result := h.CompareKey(key)
		// key -- boundary -->
		if result >= 0 { // 在新的hashtable中
			ok = setItem(h.NewHashTable.Load().([]*LinkedList), key, value)
		} else { // boundary -- key -->
			ok = setItem(h.HashTable.Load().([]*LinkedList), key, value)
		}
	} else {
		ok = setItem(h.HashTable.Load().([]*LinkedList), key, value)
	}
	if ok {
		atomic.AddInt64(&h.Size, 1)
	}
}

func (h *HashMap) Delete(key string) {
	var ok bool
	if h.rehashDoing == Running {
		result := h.CompareKey(key)
		// key -- boundary -->
		if result >= 0 { // 在新的hashtable中
			ok = delItem(h.NewHashTable.Load().([]*LinkedList), key)
		} else { // boundary -- key -->
			ok = delItem(h.HashTable.Load().([]*LinkedList), key)
		}
	} else {
		ok = delItem(h.HashTable.Load().([]*LinkedList), key)
	}
	if ok {
		atomic.AddInt64(&h.Size, -1)
	}
}

func (h *HashMap) Get(key string) (interface{}, bool) {
	if h.rehashDoing == Running {
		result := h.CompareKey(key)
		// key -- boundary -->
		if result >= 0 { // 在新的hashtable中
			return getItem(h.NewHashTable.Load().([]*LinkedList), key)
		} else { // boundary -- key -->
			return getItem(h.HashTable.Load().([]*LinkedList), key)
		}
	}
	return getItem(h.HashTable.Load().([]*LinkedList), key)
}

func delItem(hashTable []*LinkedList, key string) bool {
	idx := HashCode(key) % uint64(len(hashTable))
	list := hashTable[idx]
	return list.Delete(key)
}

func getItem(hashTable []*LinkedList, key string) (interface{}, bool) {
	idx := HashCode(key) % uint64(len(hashTable))
	list := hashTable[idx]
	return list.Get(key)
}

func setItem(hashTable []*LinkedList, key string, value interface{}) bool {
	idx := HashCode(key) % uint64(len(hashTable))
	list := hashTable[idx]
	return list.AddOrUpdate(key, value)
}

func (h *HashMap) CompareKey(key string) int {
	tableIdx2 := HashCode(key)
	if h.tableIdx < tableIdx2 {
		return -1
	} else if h.tableIdx > tableIdx2 {
		return 1
	} else if h.keyIdentify < key {
		return -1
	} else if h.keyIdentify > key {
		return 1
	} else {
		return 0
	}
}
