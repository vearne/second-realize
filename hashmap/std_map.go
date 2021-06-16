package hashmap

import "sync"

type StdMap struct {
	sync.RWMutex
	innerMap map[string]interface{}
}

func (m *StdMap) Get(key string) (interface{}, bool) {
	m.RLock()
	defer m.RUnlock()
	value, ok := m.innerMap[key]
	return value, ok
}

func (m *StdMap) Delete(key string) {
	m.Lock()
	defer m.Unlock()
	delete(m.innerMap, key)
}

func (m *StdMap) Set(key string, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.innerMap[key] = value
}

func NewStdMap(size int) *StdMap {
	m := StdMap{}
	m.innerMap = make(map[string]interface{}, size)
	return &m
}
