package hashmap

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

func init() {
	rand.Seed(0)
}

type MapInterface interface {
	Delete(key string)
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

func benchmark(b *testing.B, m MapInterface, set, get, delete int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for k := 0; k < set*1000; k++ {
			wg.Add(1)
			go func() {
				key := strconv.Itoa(rand.Int())
				value := key
				m.Set(key, value)
				wg.Done()
			}()
		}
		for k := 0; k < get*1000; k++ {
			wg.Add(1)
			go func() {
				key := strconv.Itoa(rand.Int())
				m.Get(key)
				wg.Done()
			}()
		}
		for k := 0; k < delete*1000; k++ {
			wg.Add(1)
			go func() {
				key := strconv.Itoa(rand.Int())
				m.Delete(key)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkStdSetMore(b *testing.B) {
	m := NewStdMap(2)
	benchmark(b, m, 9, 1, 1)
}

func BenchmarkStdGetMore(b *testing.B) {
	m := NewStdMap(2)
	benchmark(b, m, 1, 9, 1)
}

func BenchmarkStdEqual(b *testing.B) {
	m := NewStdMap(2)
	benchmark(b, m, 5, 5, 5)
}

func BenchmarkMyHashTableSetMore(b *testing.B) {
	h := NewHashMap(2, 1)
	benchmark(b, h, 9, 1, 1)
}

func BenchmarkMyHashTableGetMore(b *testing.B) {
	h := NewHashMap(2, 1)
	benchmark(b, h, 1, 9, 1)
}

func BenchmarkMyHashTableEqual(b *testing.B) {
	h := NewHashMap(2, 1)
	benchmark(b, h, 5, 5, 5)
}
